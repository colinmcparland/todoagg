package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/*
	Struct to parse Basecamps JSON reply to a token we can use
*/
type BasecampAccessToken struct {
	AccessToken string `json:"access_token"`
}

/*
	Function to handle the Basecamp 3 oAuth2 flow
*/
func BasecampOauthHandler(response http.ResponseWriter, request *http.Request) {

	/*
		Check to see if there is a code.  If so, send it off to get a token
	*/
	request.ParseForm()
	args := request.Form

	/*
		If there is a code in the URL, use it to get an access token from Basecamp.
	*/
	if len(args["code"]) > 0 {

		/*
			TODO:  Add URLs and keys into ENV Variables
		*/
		code := args["code"][0]
		posturl := bytes.NewBufferString("https://launchpad.37signals.com/authorization/token?")

		posturl.WriteString("type=web_server&client_id=3ddc0b5254ca2ff483e49c3665a45b9f9d95c0d1&redirect_uri=http://localhost:8081/basecamp-oauth2/&client_secret=008b3a77b37ad50cf9cb59e430cc70928620cab3&code=")
		posturl.WriteString(code)

		/*
			Use the URL we just built to send a POST request
		*/
		resp, err := http.Post(posturl.String(), "body/type", bytes.NewBuffer([]byte("")))

		if err != nil {
			return
		}

		/*
			Parse the JSON response and extract the access token from it
		*/
		var access_token BasecampAccessToken
		err = json.NewDecoder(resp.Body).Decode(&access_token)

		if err != nil {
			return
		}

		if len(access_token.AccessToken) > 0 {

			/*
				Store the token in a cookie so it can be used for future imports without worrying about 2 week TTL, since we can set it in the cookie header
			*/

			expiration := time.Now().Add(14 * 24 * time.Hour)
			cookie := http.Cookie{Name: "basecamp_token", Value: string(access_token.AccessToken), Expires: expiration, Path: "/"}
			http.SetCookie(response, &cookie)
		}

		fmt.Println(request.Cookie("basecamp_token"))

	}
}
