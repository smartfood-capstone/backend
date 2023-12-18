package foods

import "github.com/sirupsen/logrus"

type service struct {
	r IRepository
	l *logrus.Logger
}

type IService interface{}

func NewService(r IRepository, l *logrus.Logger) IService {
	return &service{
		r: r,
		l: l,
	}
}
