package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

const (
	defaultIssuer = "Property Service API"
)

var (
	defaultTokenTTL = time.Duration(time.Hour * 24 * 30)
	InvalidTokenErr = errors.New("Invalid Token")
)

type AuthUser struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type Auth interface {
	BearerAuth(t string) (*AuthUser, error)
	//
	CreateToken(userID string) (string, error)
	VerifyPassword(pwd, hashedPwd string) bool
}

type auth struct {
	enable bool
	s []byte
}

func NewAuth(secret string) Auth {
	return &auth{
		enable: true,
		s:      []byte(secret),
	}
}

func (a *auth) SetEnabled(e bool) {
	a.enable = e
}

func (a auth) VerifyPassword(pwd, hashedPwd string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPwd),
		[]byte(pwd),
	)

	return err == nil
}

func (a auth) BearerAuth(t string) (*AuthUser, error) {
	if strings.HasPrefix(t, "Bearer ") {
		token, err := jwt.ParseWithClaims(t[7:], &AuthUser{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, InvalidTokenErr
			}
			return a.s, nil
		})
		if err != nil {
			jwtError, ok := err.(*jwt.ValidationError)
			if ok {
				if (jwtError.Errors & jwt.ValidationErrorExpired) != jwt.ValidationErrorExpired {
					return nil, InvalidTokenErr
				}
			}
		}
		if claims, ok := token.Claims.(*AuthUser); ok {
			return claims, nil
		}
	}
	return nil, InvalidTokenErr
}

func (a auth) CreateToken(userID string) (string, error) {
	u := AuthUser{
		UserID:         userID,
	}
	u.ExpiresAt = time.Now().Add(defaultTokenTTL).Unix()
	u.IssuedAt = time.Now().Unix()
	u.Issuer = defaultIssuer
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	return token.SignedString(a.s)
}