package email

import (
	"architecture/message"
	"fmt"
)

type register struct{}

func NewEmailRegister() *register {
	return &register{}
}

func (s *register) Send(from string, to string, content string) error {
	fmt.Printf("Send sms register from %s to %s with content '%s'\n", from, to, content)

	return nil
}

func (s *register) Support(messageType string) bool {
	return message.TypeEmailRegister == messageType
}
