package service

import "fmt"

type mailService struct {
}

func (m *mailService) sendMail(to, msg string) error {

	fmt.Println("SEND MESSAGE TO", to, " That", msg)

	return nil
}
