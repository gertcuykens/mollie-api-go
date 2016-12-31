package mollie

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Payment Mollie
type Payment struct {
	Id    string
	Links struct {
		PaymentUrl  string
		RedirectUrl string
	}
}

// Transaction Mollie
type Transaction struct {
	Id              string
	Mode            string
	CreatedDatetime string
	Status          string
	PaidDatetime    string
	Amount          string
	Description     string
	Method          string
	Metadata        map[string]string
	details         struct {
		CardNumber string
	}
	Locale string
	Links  struct {
		WebhookUrl  string
		RedirectUrl string
	}
}

// Token Mollie
var Token = ""

// GetPayment api.mollie.nl/v1/payments
func GetPayment(client *http.Client, body io.Reader) (response *http.Response, err error) {
	request, _ := http.NewRequest("POST", "https://api.mollie.nl/v1/payments", body)
	request.Header.Set("Authorization", "Bearer "+Token)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return client.Do(request)
}

// GetTransaction api.mollie.nl/v1/payments/{id}
func GetTransaction(client *http.Client, id string) (response *http.Response, err error) {
	request, _ := http.NewRequest("GET", "https://api.mollie.nl/v1/payments/"+id, nil)
	request.Header.Set("Authorization", "Bearer "+Token)
	return client.Do(request)
}

// Refund api.mollie.nl/v1/payments/{id}/refunds amount=5.95
func Refund(client *http.Client, id string, amount string) (response *http.Response, err error) {
	request, _ := http.NewRequest("POST", "https://api.mollie.nl/v1/payments/"+id+"/refunds", strings.NewReader(amount))
	request.Header.Set("Authorization", "Bearer "+Token)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return client.Do(request)
}

//Balance api.mollie.nl/v1/payments/{id}/refunds
func Balance(client *http.Client, id string) (response *http.Response, err error) {
	request, _ := http.NewRequest("GET", "https://api.mollie.nl/v1/payments/"+id+"/refunds", nil)
	request.Header.Set("Authorization", "Bearer "+Token)
	return client.Do(request)
}

// Method api.mollie.nl/v1/methods?count=50&offset=0
func Method(client *http.Client) (response *http.Response, err error) {
	request, _ := http.NewRequest("GET", "https://api.mollie.nl/v1/methods?count=50&offset=0", nil)
	request.Header.Set("Authorization", "Bearer "+Token)
	return client.Do(request)
}

// Issuer api.mollie.nl/v1/issuers?count=50&offset=0
func Issuer(client *http.Client) (response *http.Response, err error) {
	request, _ := http.NewRequest("GET", "https://api.mollie.nl/v1/issuers?count=50&offset=0", nil)
	request.Header.Set("Authorization", "Bearer "+Token)
	return client.Do(request)
}

// Webhook Mollie
func Webhook(client *http.Client, id string, fn func(k string, v string)) error {

	response, err := GetTransaction(client, id)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	b, _ := ioutil.ReadAll(response.Body)

	var t Transaction
	err = json.Unmarshal(b, &t)
	if err != nil { // || t == (Transaction{})
		return err
	}

	for k, v := range t.Metadata {
		fn(k, v)
	}

	return nil

}
