package service

type service struct {
	buildDate  string
	commitHash string
}


type Service interface {
	GetBuildDate() string
	GetCommitHash() string
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


func (s service) GetBuildDate() string {
	return s.buildDate

}

func (s service) GetCommitHash() string {
	return s.commitHash
}