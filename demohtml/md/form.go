package md

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gertcuykens/httx"
	"github.com/gertcuykens/httx/appengine"
	"golang.org/x/net/context"
)

func init() {
	http.Handle("/form.md", httx.GzipHTTP(appengine.ContextHandler{form}))
}

func form(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	type order struct {
		Email   string
		Method  string
		Issuer  string
		Product []struct {
			Name        string
			Description string
			Quantity    float64
			Price       float64
			Currency    string
		}
	}

	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)

	var o order
	err = json.Unmarshal(b, &o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

	md := new(bytes.Buffer)
	tmpl := template.Must(template.New("form.html").Funcs(fnc).ParseFiles("md/form.html"))
	err = tmpl.Execute(md, o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sha := sha1.Sum(md.Bytes())
	etag := base64.URLEncoding.EncodeToString(sha[:])
	if strings.Contains(r.Header.Get("If-None-Match"), etag) {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.Header().Set("ETag", etag)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(md.Bytes())

}
