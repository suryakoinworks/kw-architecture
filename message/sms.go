package message

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"
)

const smsTypePrefix = "sms-"

type (
	SmsSender interface {
		Send(to string, content string) error
		Support(messageType string) bool
	}

	sms struct {
		senders []SmsSender
	}
)

func NewSmsHandler(senders ...SmsSender) *sms {
	return &sms{senders: senders}
}

func (s *sms) Support(message *message) bool {
	return message.IsValidType() && strings.HasPrefix(message.messageType, smsTypePrefix)
}

func (s *sms) Send(message *message) error {
	wd, _ := os.Getwd()

	var content bytes.Buffer

	if message.params == nil {
		message.params = map[string]interface{}{}
	}

	message.params["From"] = message.from

	tpl := template.Must(template.ParseFiles(fmt.Sprintf("%s/templates/%s.txt", wd, message.messageType)))
	if err := tpl.Execute(&content, message.params); err != nil {
		return err
	}

	for _, s := range s.senders {
		if s.Support(message.messageType) {
			return s.Send(message.to, content.String())
		}
	}

	return errors.New(fmt.Sprintf("unhandled sms[type: %s]", strings.TrimPrefix(message.messageType, smsTypePrefix)))
}
