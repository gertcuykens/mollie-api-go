package demomodule

import (
	"encoding/json"
	"net/http"

	mollie "github.com/gertcuykens/mollie-api-go"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func webhook(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	client := urlfetch.Client(ctx)
	err := mollie.Webhook(client, r.FormValue("id"), func(t mollie.Transaction) error {
		key := datastore.NewKey(ctx, "transaction", t.Id, 0, nil)
		_, err := datastore.Put(ctx, key, &t)
		if err != nil {
			log.Errorf(ctx, "Webhook: %s", err.Error())
		}
		b, err := json.Marshal(t)
		if err != nil {
			return err
		}
		w.Header().Set("Cache-Control", "must-revalidate")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func transaction(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	client := urlfetch.Client(ctx)
	err := mollie.Webhook(client, r.FormValue("id"), func(t mollie.Transaction) error {
		b, err := json.Marshal(t)
		if err != nil {
			return err
		}
		w.Header().Set("Cache-Control", "must-revalidate")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
