package demomodule

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/gertcuykens/httx"
	"github.com/gertcuykens/httx/appengine"
	mollie "github.com/gertcuykens/mollie-api-go"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

const SERVER = "https://mollie-api-go.appspot.com"

func init() {
	mollie.Token = "test_D3BBiC7YpALzMnXmUKqNpQSzuqdaHa"
	http.Handle("/payment.json", httx.CorsHandler(appengine.ContextHandler{payment}))
	http.Handle("/method.json", httx.CorsHandler(appengine.ContextHandler{method}))
	http.Handle("/issuer.json", httx.CorsHandler(appengine.ContextHandler{issuer}))
	http.Handle("/webhook.json", httx.CorsHandler(appengine.ContextHandler{webhook}))
	http.Handle("/transaction.json", httx.CorsHandler(appengine.ContextHandler{transaction}))
	http.Handle("/payment.csv", httx.CorsHandler(appengine.ContextHandler{hCsv}))
}

func payment(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	client := urlfetch.Client(ctx)

	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	log.Debugf(ctx, string(b[:]))

	var order Order
	err = json.Unmarshal(b, &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ur, _ := url.Parse(SERVER + "/?transaction")
	uw, _ := url.Parse(SERVER + "/webhook.json")

	v := url.Values{}
	var amount float64
	for _, e := range order.Product {
		// TODO: Verify e.Name/e.Price
		v.Add("metadata["+e.Name+"]", strconv.FormatFloat(e.Quantity, 'f', -1, 64))
		amount = amount + (e.Price * e.Quantity)
	}
	v.Add("amount", strconv.FormatFloat(amount, 'f', 2, 64))
	v.Add("method", order.Method)
	v.Add("issuer", order.Issuer)
	v.Add("description", order.Email)
	v.Add("redirectUrl", ur.String())
	v.Add("webhookUrl", uw.String())
	log.Debugf(ctx, "%s\n", v.Encode())

	response, err := mollie.GetPayment(client, strings.NewReader(v.Encode()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()
	b, err = ioutil.ReadAll(response.Body)

	var payment mollie.Payment
	err = json.Unmarshal(b, &payment)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if &payment.Links == nil {
		log.Debugf(ctx, "\n\n%s\n\n", string(b[:]))
		http.Error(w, string(b[:]), http.StatusInternalServerError)
		return
	}

	_, err = save(ctx, order.Email, payment.Id, order.Product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func hCsv(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	var data []Product
	data, err := csv2(ctx, r.FormValue("email"))
	if err != nil {
		log.Debugf(ctx, "%s", err.Error())
	}

	var list [][]string

	if data != nil {
		v := reflect.ValueOf(data[0])
		rec := make([]string, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			s := fmt.Sprintf("%v", v.Type().Field(i).Name)
			rec[i] = s
		}
		list = append(list, rec)
	}

	for _, e := range data {
		v := reflect.ValueOf(e)
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

func method(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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

func issuer(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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
