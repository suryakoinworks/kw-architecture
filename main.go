package main

import (
	"architecture/email"
	"architecture/interfaces"
	"architecture/message"
	"architecture/sms"
)

func main() {
	messenger := message.NewService(
		message.NewEmailHandler(
			email.NewEmailRegister(),
			email.NewEmailFunding(),
		),
		message.NewSmsHandler(
			sms.NewSmsRegister(),
			sms.NewSmsFunding(),
		),
	)

	factory := interfaces.NewInterfaceFactory(
		interfaces.GRpc{Service: messenger},
		interfaces.Mobile{Service: messenger},
		interfaces.Rest{Service: messenger},
	)

	factory.Run()
}
