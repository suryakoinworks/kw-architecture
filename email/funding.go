package email

import (
	"architecture/message"
	"fmt"
)

type funding struct{}

func NewEmailFunding() *funding {
	return &funding{}
}

func (s *funding) Send(from string, to string, content string) error {
	fmt.Printf("Send email funding from %s to %s with content '%s'\n", from, to, content)

	return nil
}

func (s *funding) Support(messageType string) bool {
	return message.TypeEmailFunding == messageType
}
