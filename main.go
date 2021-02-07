package main

import (
	"log"
	"notifservice/cmd"
	"os"

	"github.com/Smart-Pot/pkg"
	"github.com/Smart-Pot/pkg/adapter/amqp"
)

func main() {

	er(
		pkg.Config.ReadConfig(),
		amqp.Set(pkg.Config.AMQPAddress),
		cmd.Execute(),
	)

}

func er(es ...error) {
	for _, e := range es {
		if e != nil {
			log.Fatal(e)
			os.Exit(1)
		}
	}
}
