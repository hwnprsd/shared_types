package shared_types

const WORK_TYPE_SEND_MAIL = "SEND_MAIL"

type SendMailMessage struct {
	MessagingBase
	EmailAddress   string
	Subject        string
	BodyTemplate   string
	TemplateValues []string
}

func NewSendMailMessage(email string, subject string, template string, values []string) *SendMailMessage {
	return &SendMailMessage{
		MessagingBase: MessagingBase{
			WorkType: WORK_TYPE_SEND_MAIL,
		},
		EmailAddress:   email,
		Subject:        subject,
		BodyTemplate:   template,
		TemplateValues: values,
	}
}
