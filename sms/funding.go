package sms

import (
	"architecture/message"
	"fmt"
)

type funding struct{}

func NewSmsFunding() *funding {
	return &funding{}
}

func (s *funding) Send(to string, content string) error {
	fmt.Printf("Send sms funding to %s with content '%s'\n", to, content)

	return nil
}

func (s *funding) Support(messageType string) bool {
	return message.TypeSmsFunding == messageType
}
