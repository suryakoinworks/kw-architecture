package message

import (
	"errors"
	"fmt"

	"golang.org/x/exp/slices"
)

func NewMessage(messageType string, from string, to string, params map[string]interface{}) *message {
	return &message{
		messageType: messageType,
		from:        from,
		to:          to,
		params:      params,
	}
}

func (m *message) IsValidType() bool {
	return slices.Contains([]string{
		TypeSmsRegister, TypeSmsFunding,
		TypeEmailRegister, TypeEmailFunding,
		TypePushNotificationRegister, TypePushNotificationFunding,
	}, m.messageType)
}

func NewService(senders ...Sender) *Service {
	return &Service{senders: senders}
}

func (h *Service) Send(message *message) error {
	fmt.Println(message.messageType)

	for _, s := range h.senders {
		if s.Support(message) {
			return s.Send(message)
		}
	}

	return errors.New(fmt.Sprintf("unhandled message[type: %s]", message.messageType))
}
