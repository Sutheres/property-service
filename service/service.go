package service

import (
	"github.com/property/internal/datastore/database"
	"github.com/property/internal/models"
)

type service struct {
	buildDate  string
	commitHash string
	//
	db database.Database
}


type Service interface {
	GetBuildDate() string
	GetCommitHash() string
	//
	GetProperties() ([]models.Property, error)
}
type option func(s *service)

func NewService(buildDate, commitHash string, opts ...option) Service {
	service := &service{
		buildDate:  buildDate,
		commitHash: commitHash,
	}
	for _, opt := range opts {
		opt(service)
	}
	return service
}

func WithDatabase(db database.Database) option  {
	return func(s *service) {
		s.db = db
	}
}

func (s service) GetBuildDate() string {
	return s.buildDate

}

func (s service) GetCommitHash() string {
	return s.commitHash
}