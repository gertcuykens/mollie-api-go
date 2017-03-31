package demomail

import (
	"net/http"

	"github.com/gertcuykens/httx"
	"github.com/gertcuykens/httx/appengine"
	"golang.org/x/net/context"
)

func init() {
	http.Handle("/", httx.CorsHandler(appengine.ContextHandler{mailFn}))
}

func mailFn(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("\"offline\""))
}
