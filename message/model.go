package message

const TypeSmsRegister = "sms-register"
const TypeSmsFunding = "sms-funding"

const TypeEmailRegister = "email-register"
const TypeEmailFunding = "email-funding"

const TypePushNotificationRegister = "push-notification-register"
const TypePushNotificationFunding = "push-notification-funding"

type (
	Sender interface {
		Send(message *message) error
		Support(message *message) bool
	}

	message struct {
		messageType string
		from        string
		to          string
		params      map[string]interface{}
	}

	Service struct {
		senders []Sender
	}
)
