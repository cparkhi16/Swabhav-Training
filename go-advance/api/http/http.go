package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ChuckNorris struct {
	Categories []string `json:"categories"`
	CreatedAt  string   `json:"created_at"`
	IconURL    string   `json:"icon_url"`
	ID         string   `json:"id"`
	UpdatedAt  string   `json:"updated_at"`
	URL        string   `json:"url"`
	Value      string   `json:"value"`
}

func main() {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	/*body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}*/
	//Convert the body to type string
	//sb := string(body)
	//log.Printf(sb)

	var c ChuckNorris
	er := json.NewDecoder(resp.Body).Decode(&c)
	if er != nil {
		log.Fatal(er)
	}
	/*e := json.Unmarshal([]byte(body), &c)
	if e != nil {
		log.Fatal(e)
	}*/
	fmt.Println("Decoded data ", c.IconURL)

}
