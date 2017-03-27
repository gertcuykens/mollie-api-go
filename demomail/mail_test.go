package demomail

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestMail(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	t.Run("Mail", func(t *testing.T) {
		buf := new(bytes.Buffer)

		var mail = struct {
			Server  string
			Email   string
			Product []struct {
				Name     string
				Quantity float64
				Price    float64
			}
		}{
			Server: "server",
			Email:  "test@test",
			Product: []struct {
				Name     string
				Quantity float64
				Price    float64
			}{
				struct {
					Name     string
					Quantity float64
					Price    float64
				}{
					Name:     "p1",
					Quantity: 1,
					Price:    1,
				},
				struct {
					Name     string
					Quantity float64
					Price    float64
				}{
					Name:     "p2",
					Quantity: 2,
					Price:    2,
				},
			},
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
		tmpl.Execute(buf, mail)
		t.Logf("%v", tmpl.Name())

		send(ctx, mail.Email, buf)
	})

}
