package message

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

const emailTypePrefix = "email-"

type (
	EmailSender interface {
		Send(from string, to string, content string) error
		Support(messageType string) bool
	}

	email struct {
		senders []EmailSender
	}
)

func NewEmailHandler(senders ...EmailSender) *email {
	return &email{senders: senders}
}

func (s *email) Support(message *message) bool {
	return message.IsValidType() && strings.HasPrefix(message.messageType, emailTypePrefix)
}

func (s *email) Send(message *message) error {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var content bytes.Buffer

	tpl := template.Must(template.ParseFiles(fmt.Sprintf("%s/templates/%s.txt", wd, message.messageType)))
	if err := tpl.Execute(&content, message.params); err != nil {
		return err
	}

	for _, s := range s.senders {
		if s.Support(message.messageType) {
			return s.Send(message.from, message.to, content.String())
		}
	}

	return errors.New(fmt.Sprintf("unhandled email[type: %s]", strings.TrimPrefix(message.messageType, emailTypePrefix)))
}
