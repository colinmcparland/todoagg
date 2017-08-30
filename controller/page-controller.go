package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

type HomepageData struct {
	BasecampToken string
}

/*
	Handle various HTTP error codes
*/
func ErrorHandler(response http.ResponseWriter, request *http.Request, status int) {

	/*
		Check if the status is 404, in case we add future error codes to this function
	*/
	if status == http.StatusNotFound {
		http.ServeFile(response, request, "html/404.html")
	}
}

/*
	Handle a homepage request
*/
func HttpHomeHandler(response http.ResponseWriter, request *http.Request) {

	/*
		First off, 404
	*/
	if request.URL.Path != "/" {
		ErrorHandler(response, request, http.StatusNotFound)

		/*
			Break out of the function, 404'd!
		*/
		return
	}

	/*
		Check for POST/GET arguments.  If the user is trying to import, look now
	*/
	request.ParseForm()
	args := request.Form

	if len(args["integration_request"]) > 0 {

		integration_request := args["integration_request"][0]

		/*
			If we have Basecamp, redirect user to their OAuth2 flow
		*/
		if integration_request == "basecamp" {

			basecamp_auth_url := "https://launchpad.37signals.com/authorization/new?type=web_server&client_id=3ddc0b5254ca2ff483e49c3665a45b9f9d95c0d1&redirect_uri=http://localhost:8081/basecamp-oauth2/"

			http.Redirect(response, request, basecamp_auth_url, http.StatusFound)
			return
		}
	}

	t, err := template.ParseFiles("html/index.html")

	if err != nil {
		return
	}

	var data HomepageData
	c, err := request.Cookie("basecamp_token")
	fmt.Println(c, err)
	data.BasecampToken = c.Value
	t.Execute(response, data)
}

/*
	Handle a request to the main checklist page
*/

func HttpListHandler(response http.ResponseWriter, request *http.Request) {

}
