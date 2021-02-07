package endpoints

import "github.com/Smart-Pot/pkg/adapter/amqp"

func MakeVerificationMailConsumer() (amqp.Consumer, error) {
	return amqp.MakeConsumer("verif-mail-1", "VerificationMail")
}
