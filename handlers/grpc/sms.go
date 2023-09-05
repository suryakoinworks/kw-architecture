package grpc

import (
	"architecture/message"
	context "context"
	"net/http"
)

type sms struct {
	service *message.Service
	UnimplementedSmsServer
}

func NewSmsGRpc(service *message.Service) *sms {
	return &sms{service: service}
}

func (s *sms) Register(c context.Context, msg *Message) (*Status, error) {
	return s.handle(c, message.TypeSmsRegister, msg)
}

func (s *sms) Funding(c context.Context, msg *Message) (*Status, error) {
	return s.handle(c, message.TypeSmsFunding, msg)
}

func (s *sms) handle(c context.Context, smsType string, msg *Message) (*Status, error) {
	if msg.GetTo() == "" {
		return &Status{
			Code:    http.StatusBadRequest,
			Message: "receiver is empty",
		}, nil
	}

	// TODO: Validate receiver (mobile number validation)

	params := map[string]interface{}{}
	for _, param := range msg.GetParams() {
		params[param.GetKey()] = param.GetValue()
	}

	model := message.NewMessage(smsType, "", msg.To, params)
	if !model.IsValidType() {
		return &Status{
			Code:    http.StatusNotAcceptable,
			Message: "message type is not valid",
		}, nil
	}

	if err := s.service.Send(model); err != nil {
		return &Status{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	return &Status{
		Code:    http.StatusOK,
		Message: "sms sent",
	}, nil
}
