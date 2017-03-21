package md

import (
	"crypto/sha1"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gertcuykens/httx"
	"github.com/gertcuykens/httx/appengine"
	"golang.org/x/net/context"
)

func init() {
	http.Handle("/css.md", httx.GzipHTTP(appengine.ContextHandler{index}))
}

func index(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	md, err := ioutil.ReadFile("md/css.md")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sha := sha1.Sum(md)
	etag := base64.URLEncoding.EncodeToString(sha[:])
	if strings.Contains(r.Header.Get("If-None-Match"), etag) {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.Header().Set("ETag", etag)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(md)

}
