package main

import (
	"github.com/user/aggtodo/controller"
	"log"
	"net/http"
)

func main() {

	/*
		Serve up the home page
	*/

	http.HandleFunc("/", controller.HttpHomeHandler)

	/*
		Serve up the list page
	*/

	http.HandleFunc("/my-list", controller.HttpListHandler)

	/*
		URL to handle OAuth2 responses
	*/

	http.HandleFunc("/basecamp-oauth2/", controller.BasecampOauthHandler)

	/*
		Since we're refrencing from the GOROOT, lets strip the /css/ parts out of the file paths so we're not double-refrencing them from our HTML files
	*/

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
