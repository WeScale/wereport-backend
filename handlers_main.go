package main

import (
	"fmt"
	"html"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index, %q", html.EscapeString(r.URL.Path))
}
