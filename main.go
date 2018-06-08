package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/preichenberger/oauth2-router/redirector"
)

func OauthRouterServer(w http.ResponseWriter, req *http.Request) {
	redirectUrl, err := redirector.CreateUrl(req.URL.Query())
	if err != nil {
		switch err.(type) {
		case redirector.ValidationError:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("400 - %s\n", err)))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 Internal Server Error\n"))
		}

		return
	}

	http.Redirect(w, req, redirectUrl.String(), 301)
}

func main() {
	var port int
	var help bool
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.BoolVar(&help, "h", false, "help")
	flag.Parse()

	if help {
		println("Usage: oauth2-redirector [-port 8080]")
		os.Exit(0)
	}

	http.HandleFunc("/", OauthRouterServer)
	log.Printf("Starting OAuth2 Router on port: %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
