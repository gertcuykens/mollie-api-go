package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"appengine"
	"appengine/urlfetch"
)

func payment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "application/json")

	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	p := mollie{
		token:  "test_D3BBiC7YpALzMnXmUKqNpQSzuqdaHa",
		body:   r.Body,
		client: client,
	}

	//p.body = strings.NewReader("amount=10.00&description=My first API payment&redirectUrl=https://webshop.example.org/order/12345/&metadata[order_id]=12345")

	err := p.order()
	if err != nil {
		c.Errorf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer p.response.Body.Close()
	b, err := ioutil.ReadAll(p.response.Body)
	v := string(b[:])

	//c.Infof(v)

	type link struct {
		PaymentUrl  string
		RedirectUrl string
	}

	var response struct {
		Id    string
		Links link
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		c.Errorf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if &response.Links == nil {
		c.Errorf(v)
		http.Error(w, v, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("{\"id\":\"" + response.Id + "\", \"link\":\"" + response.Links.PaymentUrl + "\"}"))
}

func order(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "application/json")

	c := appengine.NewContext(r)

	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	c.Infof(string(b[:]))

	v := Order{
		Context: c,
	}

	err = json.Unmarshal(b, &v)
	if err != nil {
		c.Errorf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v.save()

	b, err = json.Marshal(v.Meta)
	if err != nil {
		c.Errorf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Infof(string(b))

	w.Write(b)

}
