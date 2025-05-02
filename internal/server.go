package internal

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func addDefaultRoute() {
	// This function will add the default routes to the server
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/Card-Page", CardPagehandler)
	http.HandleFunc("/filters", HomePageSwitch)
}

func CreateAndListenServer(port int) {
	// This function will create the server and listen on the specified port
	addDefaultRoute()

	// This will serve the static files
	fs := http.FileServer(http.Dir("frontend/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// This starts the server
	fmt.Print("Server was Started on port ", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), logRequest(http.DefaultServeMux)))
}

func logRequest(handler http.Handler) http.Handler {
	// This function prints the request to the console
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
