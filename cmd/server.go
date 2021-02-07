package cmd

import (
	"net/http"
	"notifservice/endpoints"
	"notifservice/service"
	"notifservice/transport"

	"os"
	"time"

	"github.com/Smart-Pot/pkg"
	"github.com/go-kit/kit/log"
)

func startServer() error {
	// Defining Logger
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	service := service.NewService(logger)

	s := http.Server{
		Addr:         pkg.Config.Server.Address, // configure the bind address
		Handler:      nil,                       // set the default handler
		ErrorLog:     nil,                       // set the logger for the server
		ReadTimeout:  5 * time.Second,           // max time to read request from the client
		WriteTimeout: 10 * time.Second,          // max time to write response to the client
		IdleTimeout:  120 * time.Second,         // max time for connections using TCP Keep-Alive
	}

	if err := startAMQP(service); err != nil {
		return err
	}
	return s.ListenAndServe()
}

func startAMQP(s service.Service) error {
	c, err := endpoints.MakeVerificationMailConsumer()
	if err != nil {
		return err
	}

	task := transport.MakeVerificationMailTask(c, s)
	go task()
	return nil
}
