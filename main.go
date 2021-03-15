package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

func AllowedKey(key string) bool {
	return key == "b1946ac92492d2347c6235b4d2611184"
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not protected!\n")
}

func Protected(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !AllowedKey(ps.ByName("key")) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
	target, err := url.Parse("http://10.1.0.9:8545")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	httputil.NewSingleHostReverseProxy(target)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/archive/:key", Protected)

	log.Fatal(http.ListenAndServe(":8081", router))
}
