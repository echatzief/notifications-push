package mail

import (
	"fmt"

	"notifications-push/config"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(from string, to string, name string, message string) {
	m := mail.NewV3Mail()

	m.SetFrom(mail.NewEmail("from", from))

	m.SetTemplateID(config.TEMPLATE_ID)

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail("to", to),
	}
	p.AddTos(tos...)

	p.SetDynamicTemplateData("name", name)
	p.SetDynamicTemplateData("message", message)

	m.AddPersonalizations(p)

	request := sendgrid.GetRequest(config.SENDGRID_API, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = mail.GetRequestBody(m)
	request.Body = Body
	_, err := sendgrid.API(request)
	if err != nil {
		fmt.Println("Failed to sent mail")
	} else {
		fmt.Println("Succesfully sent mail to ", to)
	}
}
