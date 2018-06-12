package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/preichenberger/oauth2-router/redirector"
)

var _redirector *redirector.Redirector

func OauthRouterServer(w http.ResponseWriter, req *http.Request) {
	redirectUrl, err := _redirector.CreateUrl(req.URL.Query())
	if err != nil {
		switch err.(type) {
		case redirector.Error:
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
	var help bool
	var port int
	var whitelist string
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.StringVar(&whitelist, "whitelist", "*", "comma-delimited list of whitelist domains i.e '*.github.com,pizza.com'")
	flag.BoolVar(&help, "h", false, "help")
	flag.Parse()

	env_port := os.Getenv("PORT")
	if len(env_port) != 0 {
		env_port_int, err := strconv.Atoi(env_port)
		if err != nil {
			log.Fatal(err)
		}
		port = env_port_int
	}

	env_whitelist := os.Getenv("WHITELIST")
	if len(env_whitelist) != 0 {
		whitelist = env_whitelist
	}

	if help {
		println("Usage: oauth2-redirector [-port 8080] [-whitelist *.github.com,pizza.com")
		os.Exit(0)
	}

	_redirector = redirector.NewRedirector(whitelist)

	http.HandleFunc("/", OauthRouterServer)
	log.Printf("Starting OAuth2 Router on port: %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
