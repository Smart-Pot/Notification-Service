package service

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
)

type service struct {
	l      log.Logger
	mailer *mailService
}

func NewService(l log.Logger) Service {
	return &service{
		l:      l,
		mailer: newMailService(),
	}
}

type Service interface {
	SendVerificationMail(context.Context, string, string, string) error
}

func (s *service) SendVerificationMail(ctx context.Context, name, email, hash string) error {
	url := fmt.Sprintf("http://localhost:3000/user/verify/%s", hash)
	msg, err := GetActivationMail(name, url)
	if err != nil {
		return err
	}
	return s.mailer.sendMail([]string{email}, msg)
}
