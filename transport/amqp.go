package transport

import (
	"context"
	"encoding/json"
	"log"
	"notifservice/service"

	"github.com/Smart-Pot/pkg/adapter/amqp"
)

func MakeVerificationMailTask(c amqp.Consumer, s service.Service) func() {
	for {
		var r struct {
			Name  string `json:"name"`
			Hash  string `json:"hash"`
			Email string `json:"email"`
		}

		b := c.Consume()

		json.Unmarshal(b, &r)

		if err := s.SendVerificationMail(context.TODO(), r.Name, r.Email, r.Hash); err != nil {
			log.Fatal("asd", err)
		}
	}
}
