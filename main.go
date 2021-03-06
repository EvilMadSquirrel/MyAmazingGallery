package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my Amazing Gallery!</h1>")
	} else if r.URL.Path == "/contacts" {
		fmt.Fprint(w, "To get in touch please send an email to: <a href=\"mailto:minichev.s.l@gmail.com\":>minichev.s.l@gmail.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find page were you looking for :(</h1><p>Please email us if you keep being went to an invalid page</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
