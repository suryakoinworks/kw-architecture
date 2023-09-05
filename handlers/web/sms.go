package web

import (
	"architecture/message"
	"net/http"

	"github.com/labstack/echo/v4"
)

type sms struct {
	service *message.Service
	To      string                 `json:"to"`
	Params  map[string]interface{} `json:"params"`
}

func NewSmsWeb(service *message.Service) *sms {
	return &sms{service: service}
}

func (h *sms) Register(c echo.Context) error {
	return h.handle(c, message.TypeSmsRegister)
}

func (h *sms) Funding(c echo.Context) error {
	return h.handle(c, message.TypeSmsFunding)
}

func (h *sms) handle(c echo.Context, smsType string) error {
	c.Bind(h)
	if h.To == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "receiver is empty",
			"code":    http.StatusBadRequest,
		})
	}

	// TODO: Validate receiver (mobile number validation)

	msg := message.NewMessage(smsType, "KW-SMS-SVC", h.To, h.Params)
	if !msg.IsValidType() {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "message type is not valid",
			"code":    http.StatusNotAcceptable,
		})
	}

	if err := h.service.Send(msg); err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "sms sent",
		"code":    http.StatusOK,
	})
}
