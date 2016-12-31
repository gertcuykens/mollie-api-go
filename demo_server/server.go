package server

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	mollie "github.com/gertcuykens/mollie-api-go"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const SERVER = "https://mollie-api-go.appspot.com"
const SENDER = "Demo <demo@mollie-api-go.appspotmail.com>"

func BCC(e ...string) []string {
	return []string{}
}

func init() {
	mollie.Token = "test_D3BBiC7YpALzMnXmUKqNpQSzuqdaHa"
	http.HandleFunc("/mollie.go", payment)
	http.HandleFunc("/order.go", order)
	http.HandleFunc("/method.go", method)
	http.HandleFunc("/issuer.go", issuer)
	http.HandleFunc("/webhook.go", webhook)
	http.HandleFunc("/csv.go", hCsv)
	//http.HandleFunc("/_ah/mail/", incomingMail)
	//transport := http.Transport{}
	//client := &http.Client{Transport: &transport}
}

func secureGET(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fn(w, r)
			return
		}
		if r.Header.Get("Authorization") != "secret" { // r.FormValue("p")
			http.Error(w, "Unauthorized!", http.StatusUnauthorized)
			return
		}
		fn(w, r)
	}
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		//w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		//w.Header().Set("Access-Control-Allow-Credentials", "true")
		fn(w, r)
	}
}

func incomingMail(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	defer r.Body.Close()
	var b bytes.Buffer
	if _, err := b.ReadFrom(r.Body); err != nil {
		log.Errorf(ctx, "%s", err.Error())
		return
	}
	log.Debugf(ctx, "%s\n%s", r.URL.Path, b.String())
}

func requestTest(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Debugf(ctx, "%s\n%s", r.URL.Path, string(b[:]))
	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "text/plain")
	w.Write(b)
	//http.Redirect(w, r, "", http.StatusFound)
}

func hCsv(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	/*f0 := r.FormValue("id")
	if f0 == "" {
		http.Error(w, "No id!", http.StatusNoContent)
		return
	}*/

	data := &Order{
		Context: ctx,
	}
	data.csv()

	//c.Infof("%v", data.List[0])

	var list [][]string

	if data.List != nil {
		v := reflect.ValueOf(*data.List[0])
		rec := make([]string, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			s := fmt.Sprintf("%v", v.Type().Field(i).Name)
			rec[i] = s
		}
		list = append(list, rec)
	}

	for _, e := range data.List {
		v := reflect.ValueOf(*e)
		rec := make([]string, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			s := fmt.Sprintf("%v", v.Field(i).Interface())
			rec[i] = s
		}
		list = append(list, rec)
	}

	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "text/csv; charset=utf-8")
	cc := csv.NewWriter(w)
	cc.WriteAll(list)
}

/*
func send(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	email := r.FormValue("id"),
	link := r.FormValue("link")

	entity := &Entity{
	}

	var list []*Entity
	list = append(list, entity)

	//c.Infof("%v", list[0])

	data := &Transaction{
		User: email,
		List:    list,
		Context: c,
	}

	data.save()

	msg := &app.Message{
		Sender:  "Mollie <info@mollie-api-go.appspotmail.com>",
		To:      email,
		//Bcc:,
		ReplyTo: info@mollie.com,
		Subject: "Billing",
		HTMLBody: fmt.Sprintf(template, link),
		Headers: mail.Header{
			//"Content-Type": []string{"text/html; charset=UTF-8"},
			"On-Behalf-Of": []string{"iinfo@mollie.com"},
		},
	}

	//if code := r.FormValue("code"); code != "12345" {
	//   c.Errorf("Wrong code: %v", code)
	//   http.Error(w, err.Error(), http.StatusUnauthorized)
	//   return
	//}else

	if err := app.Send(c, msg); err != nil {
		c.Errorf("Error sending mail: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// c.Infof("Mail send: %v", msg)

}

const template = `<!DOCTYPE html>
<html>
<head>
  <title>Transaction</title>
  <style>
    table,
    td {
      border-collapse: collapse;
      padding: 0;
      margin: 0;
    }
  </style>
</head>
<body>
	<a href="%s">billing</a>
</body>
</html>`
*/
