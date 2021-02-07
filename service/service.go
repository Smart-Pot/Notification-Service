package service

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
)

type service struct {
	l log.Logger
}

func NewService(l log.Logger) Service {
	return &service{
		l: l,
	}
}

type Service interface {
	SendVerificationMail(context.Context, string, string) error
}

func (s *service) SendVerificationMail(ctx context.Context, email, hash string) error {
	fmt.Println("SEND MESSAGE TO", email, "HASH", hash)
	return nil
}
