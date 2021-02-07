package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"path/filepath"
)

type mailService struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	SmtpHost string `json:"smtpHost"`
	SmtpPort string `json:"smtpPort"`
	auth     smtp.Auth
}

func (m *mailService) sendMail(to []string, msg string) error {
	return smtp.SendMail(m.address(), m.auth, m.Email, to, []byte(msg))
}
func (m *mailService) address() string {
	return fmt.Sprintf("%s:%s", m.SmtpHost, m.SmtpPort)
}

func newMailService() *mailService {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "config", "email.json")

	var s mailService

	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(b, &s); err != nil {
		panic(err)
	}
	fmt.Println("HELLO S", s)
	s.auth = smtp.PlainAuth("", s.Email, s.Password, s.SmtpHost)

	return &s
}
