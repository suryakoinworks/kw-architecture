package interfaces

import (
	"architecture/handlers/mobile"
	"architecture/message"

	"github.com/labstack/echo/v4"
)

type Mobile struct {
	Service *message.Service
}

func (i Mobile) Run() {
	e := echo.New()

	sms := mobile.NewSmsMobile(i.Service)

	e.POST("/mobile/sms/register", sms.Register)
	e.POST("/mobile/sms/funding", sms.Funding)

	e.Start(":2727")
}

func (i Mobile) IsBackground() bool {
	return true
}

func (i Mobile) Priority() int {
	return 0
}
