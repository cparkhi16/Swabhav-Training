package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

type ChuckNorriss struct {
	Categories []string `json:"categories"`
	CreatedAt  string   `json:"created_at"`
	IconURL    string   `json:"icon_url"`
	ID         string   `json:"id"`
	UpdatedAt  string   `json:"updated_at"`
	URL        string   `json:"url"`
	Value      string   `json:"value"`
}

//var q = make(map[int]ChuckNorriss)

func getQuotes(i int) {
	defer wg.Done()
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		log.Fatalln(err)
	}
	var c ChuckNorriss
	er := json.NewDecoder(resp.Body).Decode(&c)
	if er != nil {
		log.Fatal(er)
	}
	//q[i] = c
	var m sync.Map
	m.Store(i, c)
	//fmt.Println(m.Load(i))
}

func main() {
	now := time.Now()
	fmt.Println("Start time ", now)
	defer func() {
		fmt.Println(time.Since(now))
	}()
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go getQuotes(i)
	}
	wg.Wait()
	//fmt.Println("Map q ", q)
}
