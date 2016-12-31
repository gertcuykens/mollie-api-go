package server

import (
	"net/http"

	mollie "github.com/gertcuykens/mollie-api-go"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func webhook(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	err := mollie.Webhook(client, r.FormValue("id"), func(k string, v string) {
		log.Debugf(ctx, "%s %s", k, v)
		// nr, _ := strconv.ParseInt(k, 10, 64)
		// randoneur, _ := getRandonneur(ctx, id, nr)
		// randoneur.TransactionStatus = t.Status
		// nr, _ = setRandonneur(ctx, randoneur, id, nr)
		// fmt.Fprintf(w, "%d", nr)
		// fmt.Fprintf(w, "%d", nr)
	})
	if err != nil {
		log.Errorf(ctx, "Webhook %s", err.Error())
	}
	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "application/json")
}
