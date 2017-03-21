package demomodule

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http/httptest"
	"testing"

	mollie "github.com/gertcuykens/mollie-api-go"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/log"
)

func TestMollieOrder(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	var mpayment mollie.Payment

	t.Run("Order", func(t *testing.T) {
		var jsonStr = []byte(`{
			"Email": "test@test",
			"Method": "",
			"Issuer": "",
			"Product": [{
				"Name": "test",
				"Description": "",
				"Quantity": 2.01,
				"Price": 1.99,
				"Currency": ""
			},
			{
				"Name": "test2",
				"Description": "",
				"Quantity": 4.01,
				"Price": 5.99,
				"Currency": ""
			}]
		}`)

		req := httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		payment(ctx, w, req)
		t.Logf("\n\n%d - %s\n\n", w.Code, w.Body.String())

		err = json.Unmarshal(w.Body.Bytes(), &mpayment)
		if err != nil {
			t.Error("Json Unmarshal: ", err)
		}
	})

	t.Run("Webhook", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/?id="+mpayment.Id, nil)
		w := httptest.NewRecorder()
		webhook(ctx, w, req)
		t.Logf("\n\n%d - %s\n\n", w.Code, w.Body.String())
	})

	t.Run("Csv", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/?email=", nil)
		w := httptest.NewRecorder()
		hCsv(ctx, w, req)
		t.Logf("\n\n%d - %s\n\n", w.Code, w.Body.String())
	})

	t.Run("Mail", func(t *testing.T) {
		buf := new(bytes.Buffer)

		var order Order
		order.Server = SERVER
		order.Email = "test@test"
		order.Product, err = csv2(ctx, order.Email)
		if err != nil {
			log.Debugf(ctx, "%s", err.Error())
		}

		var total float64
		fnc := template.FuncMap{
			"html": func(value interface{}) template.HTML {
				return template.HTML(fmt.Sprint(value))
			},
			"multiply": func(a float64, b float64) float64 {
				total = total + (a * b)
				return a * b
			},
			"total": func() float64 {
				return total
			},
		}

		tmpl := template.Must(template.New("mail.html").Funcs(fnc).ParseFiles("mail.html"))
		tmpl.Execute(buf, order)
		t.Logf("%v", tmpl.Name())

		mail(ctx, order.Email, buf)
	})
}

func TestMollieMethod(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	method(ctx, w, req)
	t.Logf("\n\n%d - %s\n\n", w.Code, w.Body.String())
}

func TestMollieIssuer(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	issuer(ctx, w, req)
	t.Logf("\n\n%d - %s\n\n", w.Code, w.Body.String())
}
