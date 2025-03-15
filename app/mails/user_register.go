package mails

import (
	"fmt"

	"github.com/goravel/framework/contracts/mail"
	"github.com/goravel/framework/facades"
)

type UserRegister struct {
	email string
	link  string
}

func NewUserRegister(email, link string) *UserRegister {
	return &UserRegister{
		email: email,
		link:  link,
	}
}

// Attachments attach files to the mail
func (receiver *UserRegister) Attachments() []string {
	return []string{}
}

// Content set the content of the mail
func (receiver *UserRegister) Content() *mail.Content {
	return &mail.Content{
		Html: fmt.Sprintf(`
					<h1>Welcome to E-Vote!</h1>
					<p>Please verify your email address by clicking the link below:</p>
					<p>This link will expire in 24 hours.</p>
					<a href="%s" style="background-color: #4CAF50; color: white; padding: 14px 20px; text-decoration: none; border-radius: 4px;">
						Verify Email
					</a>
					<p>If you didn't create an account, please ignore this email.</p>
				`, receiver.link),
	}
}

// Envelope set the envelope of the mail
func (receiver *UserRegister) Envelope() *mail.Envelope {
	return &mail.Envelope{
		From: mail.Address{
			Address: facades.Config().GetString("MAIL_FROM_ADDRESS", "evote@rizkirmdhn.cloud"),
			Name:    facades.Config().GetString("MAIL_FROM_NAME", "Evote"),
		},
		Subject: "User Registration Confirmation",
		To:      []string{receiver.email},
	}
}

// Queue set the queue of the mail
func (receiver *UserRegister) Queue() *mail.Queue {
	return &mail.Queue{}
}
