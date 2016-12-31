package server

import (
	"bytes"
	"fmt"
	"html/template"

	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	gmail "google.golang.org/appengine/mail"
)

func mail(ctx context.Context, d struct {
	Nr    int64
	ID    string
	Order Order
	User  User
}) error {

	buf := new(bytes.Buffer)
	t, _ := template.New("mail").Funcs(template.FuncMap{
		"html": func(value interface{}) template.HTML {
			return template.HTML(fmt.Sprint(value))
		},
	}).Parse(tpl)
	t.Execute(buf, d)

	msg := &gmail.Message{
		Sender: SENDER,
		To:     []string{d.User.Email},
		// Bcc:     BCC(d.Event.Organizer),
		// ReplyTo: d.Event.Organizer,
		// Subject: "[" + d.Event.Title + " " + d.Event.Date + "] Bevestiging inschrijving " + d.User.Lastname,
		// Body:    message,
		HTMLBody: buf.String(),
		// Headers: mail.Header{
		// 	//"Content-Type": []string{"text/html; charset=utf-8"},
		// 	"On-Behalf-Of": []string{"info@randonneurs.nl"},
		// },
	}

	// log.Debugf(ctx, "%v", msg)
	err := gmail.Send(ctx, msg)
	if err != nil {
		log.Errorf(ctx, "Error sending mail: %s", err.Error())
	}
	return err
}
