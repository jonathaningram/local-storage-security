package main

import (
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	publicDir := "public"

	http.Handle("/", http.FileServer(http.Dir(publicDir)))

	appengine.Main()
}
