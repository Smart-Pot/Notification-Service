package transport

import (
	"context"
	"encoding/json"
	"notifservice/service"

	"github.com/Smart-Pot/pkg/adapter/amqp"
)

func MakeVerificationMailTask(c amqp.Consumer, s service.Service) func() {
	for {
		var r struct {
			Hash  string `json:"hash"`
			Email string `json:"email"`
		}

		b := c.Consume()

		json.Unmarshal(b, &r)

		s.SendVerificationMail(context.TODO(), r.Hash, r.Email)
	}
}
