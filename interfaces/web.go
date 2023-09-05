package interfaces

import (
	"architecture/handlers/web"
	"architecture/message"

	"github.com/labstack/echo/v4"
)

type Rest struct {
	Service *message.Service
}

func (i Rest) Run() {
	e := echo.New()

	sms := web.NewSmsWeb(i.Service)

	e.POST("/web/sms/register", sms.Register)
	e.POST("/web/sms/funding", sms.Funding)

	e.Start(":1717")
}

func (i Rest) IsBackground() bool {
	return false
}

func (i Rest) Priority() int {
	return -255
}
