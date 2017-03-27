package demomodule

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	mollie "github.com/gertcuykens/mollie-api-go"
	"google.golang.org/appengine/aetest"
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
