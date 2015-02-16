package server

import (
	"io"
	"net/http"
	"strings"
)

type mollie struct {
	token    string
	id       string
	body     io.Reader
	amount   string
	client   *http.Client
	response *http.Response
}

func (p *mollie) order() (err error) {
	request, _ := http.NewRequest("POST", "https://api.mollie.nl/v1/payments", p.body)
	request.Header.Set("Authorization", "Bearer "+p.token)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	p.response, err = p.client.Do(request)
	return err
}

func (p *mollie) status() (err error) {
	request, _ := http.NewRequest("GET", "https://api.mollie.nl/v1/payments/"+p.id, nil)
	request.Header.Set("Authorization", "Bearer "+p.token)
	p.response, err = p.client.Do(request)
	return err
}

func (p *mollie) refund() (err error) { //amount=5.95
	request, _ := http.NewRequest("POST", "https://api.mollie.nl/v1/payments/"+p.id+"/refunds", strings.NewReader(p.amount))
	request.Header.Set("Authorization", "Bearer "+p.token)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	p.response, err = p.client.Do(request)
	return err
}

func (p *mollie) balance() (err error) {
	request, _ := http.NewRequest("GET", "https://api.mollie.nl/v1/payments/"+p.id+"/refunds", nil)
	request.Header.Set("Authorization", "Bearer "+p.token)
	p.response, err = p.client.Do(request)
	return err
}
