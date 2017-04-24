package demohtml

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
	http.Handle("/", httx.GzipHTTP(appengine.ContextHandler{push1}))
}

func push1(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	md, err := ioutil.ReadFile("md/index.html")
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

	// var push = [][]string{
	// 	[]string{"favicon.ico", "image/x-icon"},
	// 	[]string{"manifest.json", "application/javascript"},
	// 	[]string{"md/css.md", "text/html; charset=utf-8"},
	// }
	// push2(w, push)

}

// func push2(w http.ResponseWriter, push [][]string) {
// 	for k := range push {
// 		f, err := ioutil.ReadFile(push[k][0])
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		sha := sha1.Sum(f)
// 		etag := base64.URLEncoding.EncodeToString(sha[:])
// 		w.Header().Set("ETag", etag)
// 		w.Header().Set("Content-Type", push[k][1])
// 		if p, ok := w.(http.Pusher); ok {
// 			p.Push(push[k][0], nil)
// 		}
// 	}
// }
