package sms

import (
	"architecture/message"
	"fmt"
)

type register struct{}

func NewSmsRegister() *register {
	return &register{}
}

func (s *register) Send(to string, content string) error {
	fmt.Printf("Send sms register to %s with content '%s'\n", to, content)

	return nil
}

func (s *register) Support(messageType string) bool {
	return message.TypeSmsRegister == messageType
}
