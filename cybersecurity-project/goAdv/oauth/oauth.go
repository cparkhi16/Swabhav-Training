package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	clientID     = "1058963771578-gqagesvt69mmqmv8bk9ecni04pko6frc.apps.googleusercontent.com"
	clientSecret = "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"

	googleConfig *oauth2.Config

	stateString = "xyz"
)

func init() {
	googleConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		RedirectURL:  "http://localhost:8080/redirect",
		Endpoint:     google.Endpoint,
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/redirect", redirectHandler)
	fmt.Println("Starting server")
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var html = `<html>
	<body>
	 <a href="/login"> Log in for google</a>
	</body>
</html>`
	fmt.Fprintf(w, html)
}

// google's oauth server
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login handler called")
	url := googleConfig.AuthCodeURL(stateString)
	fmt.Println("url", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// after google login
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("redirect handler called")
	contents, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Errorf("response error %s", err.Error())
		return
	}
	fmt.Fprintf(w, "content %s\n,", contents)
}

func getUserInfo(state, code string) ([]byte, error) {
	if state != stateString {
		return nil, fmt.Errorf("Invalid state")
	}
	token, err := googleConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("Error %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("response error %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read error %s", err.Error())
	}
	return contents, nil
}
