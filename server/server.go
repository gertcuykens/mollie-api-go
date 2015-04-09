package server

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
)

//const TOKEN = "000"
//var client *http.Client

func init() {
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
	//a := r.FormValue("p") //r.Header.Get("Authorization")
	//if a != TOKEN {
	//	http.Error(w, "Unauthorized!", http.StatusUnauthorized)
	//	return
	//}
}

func incomingMail(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	defer r.Body.Close()
	var b bytes.Buffer
	if _, err := b.ReadFrom(r.Body); err != nil {
		c.Errorf("Error reading body: %v", err)
		return
	}
	c.Infof("Received mail: %v", b)
}

func webhook(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	p := mollie{
		token:  "test_D3BBiC7YpALzMnXmUKqNpQSzuqdaHa",
		id:     r.FormValue("id"),
		client: client,
	}
	p.status()

	defer p.response.Body.Close()
	b, _ := ioutil.ReadAll(p.response.Body)
	c.Infof(string(b[:]))

	type detail struct {
		CardNumber string
	}

	type link struct {
		WebhookUrl  string
		RedirectUrl string
	}

	var response struct {
		Id              string
		Mode            string
		CreatedDatetime string
		Status          string
		PaidDatetime    string
		Amount          string
		Description     string
		Method          string
		Metadata        interface{}
		details         detail
		Locale          string
		Links           link
	}

	err := json.Unmarshal(b, &response)
	if err != nil {
		c.Errorf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for k, v := range response.Metadata.(map[string]interface{}) {
		i, _ := strconv.ParseInt(k, 10, 64)
		root := datastore.NewKey(c, "user", v.(string), 0, nil)
		key := datastore.NewKey(c, "order", "", i, root)
		e := new(Entity)
		err := datastore.Get(c, key, e)
		if err != nil {
			c.Infof("Error datastore: %v", err)
			return
		}
		e.Status = response.Status
		e.Transaction = response.Id
		_, err = datastore.Put(c, key, e)
		if err != nil {
			c.Infof("Error datastore: %v", err)
			return
		}

	}

	c.Infof("%v", response.Metadata.(map[string]interface{}))

	//defer r.Body.Close()
	//b, _ := ioutil.ReadAll(r.Body)
	//c.Infof(string(b[:]))

	//w.Header().Set("Cache-Control", "must-revalidate")
	//w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte(fmt.Sprint(r.Form)))
	//http.Error(w, "requestTest", http.StatusInternalServerError)
	//http.Redirect(w, r, o.Links[0].Url, http.StatusFound)
}

func hCsv(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	/*f0 := r.FormValue("id")
	if f0 == "" {
		http.Error(w, "No id!", http.StatusNoContent)
		return
	}*/

	data := &Order{
		Context: c,
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
