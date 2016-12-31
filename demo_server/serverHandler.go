package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	mollie "github.com/gertcuykens/mollie-api-go"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func order(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	log.Infof(ctx, string(b[:]))

	v := Order{
		Context: ctx,
	}

	err = json.Unmarshal(b, &v)
	if err != nil {
		log.Errorf(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v.save()

	b, err = json.Marshal(v.Meta)
	if err != nil {
		log.Errorf(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Infof(ctx, string(b))

	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// strings.NewReader("amount=10.00&description=My first API payment&redirectUrl=https://webshop.example.org/order/12345/&metadata[order_id]=12345")
func payment(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)

	if r.FormValue("id") == "" {
		http.Error(w, "No id!", http.StatusNoContent)
		return
	}

	nr, err := strconv.ParseInt(r.FormValue("nr"), 10, 64)
	if nr > 0 && err != nil {
		http.Error(w, "No nr!", http.StatusNoContent)
		return
	}

	// event, _ := getEvent(ctx, r.FormValue("id"))
	// randoneur, _ := getRandonneur(ctx, r.FormValue("id"), nr)

	// u, _ := url.Parse(SERVER + "/status.html?id=" + r.FormValue("id") + "&nr=" + r.FormValue("nr"))
	// v := url.Values{}
	// v.Add("amount", event.Price)
	// v.Add("description", event.Title)
	// v.Add("redirectUrl", u.String())
	// v.Add("metadata["+r.FormValue("nr")+"]", r.FormValue("id"))
	// log.Debugf(ctx, "%s\n", v.Encode())

	// response, err := mollie.GetPayment(client, strings.NewReader(v.Encode()))

	response, err := mollie.GetPayment(client, r.Body)
	if err != nil {
		log.Errorf(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)

	var payment mollie.Payment
	err = json.Unmarshal(b, &payment)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if &payment.Links == nil {
		// log.Debugf(ctx, "%s", string(b[:]))
		http.Error(w, string(b[:]), http.StatusInternalServerError)
		return
	}

	// randoneur.Transaction = payment.Id
	// nr, err = setRandonneur(ctx, randoneur, r.FormValue("id"), nr)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte("{\"id\":\"" + payment.Id + "\", \"link\":\"" + payment.Links.PaymentUrl + "\"}"))
	w.Write(b)
}

func method(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)

	response, err := mollie.Method(client)
	if err != nil {
		log.Errorf(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func issuer(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)

	response, err := mollie.Issuer(client)
	if err != nil {
		log.Errorf(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

//v := string(b[:])

//c.Infof(v)

/*type method struct {
	Id string
}

var response struct {
	Data []*method
}

err = json.Unmarshal(b, &response)
if err != nil {
	c.Errorf(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}

if &response.Data == nil {
	c.Errorf(v)
	http.Error(w, v, http.StatusInternalServerError)
	return
}*/
