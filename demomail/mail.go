package demomail

import (
	"bytes"

	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	gmail "google.golang.org/appengine/mail"
)

const SENDER = "Demo <demo@mollie-api-go.appspotmail.com>"

// func BCC(e ...string) []string {
// 	return []string{}
// }

func send(ctx context.Context, email string, buf *bytes.Buffer) error {

	msg := &gmail.Message{
		Sender: SENDER,
		To:     []string{email},
		// Bcc:     BCC(),
		// ReplyTo: "",
		Subject: "Test Mollie Order",
		// Body:    message,
		HTMLBody: buf.String(),
		// Headers: mail.Header{
		// 	"Content-Type": []string{"text/html; charset=utf-8"},
		// 	"On-Behalf-Of": []string{"demo@mollie-api-go.appspotmail.com"},
		// },
	}

	log.Debugf(ctx, "\n\n%v\n\n", msg)

	err := gmail.Send(ctx, msg)
	if err != nil {
		log.Errorf(ctx, "Error sending mail: %s", err.Error())
	}
	return err
}
