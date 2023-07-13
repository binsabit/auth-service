package sms

import (
	"log"

	"github.com/koorgoo/smsc"
)

type SMSSender struct {
	Client smsc.Client
}

func NewSmsSender(login, password string) *SMSSender {
	client, err := smsc.New(smsc.Config{Login: login, Password: password})
	if err != nil {
		log.Fatal(err)
	}

	return &SMSSender{
		Client: *client,
	}
}
