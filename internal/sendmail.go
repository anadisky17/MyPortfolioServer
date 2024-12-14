package internal

import (
	"fmt"

	gomail "gopkg.in/mail.v2"
)

type Msg struct {
	Name           string
	RecipientEmail string
	Subject        string
	Message        string
}

func SendMail(m Msg) {
	// Create a new message
	message := gomail.NewMessage()

	// Set email headers

	message.SetHeader("From", m.RecipientEmail)
	message.SetHeader("To", "anadi.sky17@gmail.com")
	message.SetHeader("Subject", m.Subject)

	// Set email body
	message.SetBody("text/plain", m.Message)

	// Set up the SMTP dialer
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "anadi.sky17@gmail.com", "vevo ehrf owcx rwuu")

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Error in dailing email smtp server", err)
		panic(err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
