package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type mailService struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	SmtpHost string `json:"smtpHost"`
	SmtpPort string `json:"smtpPort"`
}

func (m *mailService) sendMail(to, msg string) error {

	fmt.Println("SEND MESSAGE TO", to, " That", msg)

	return nil
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
	return &s
}
