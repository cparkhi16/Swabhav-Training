package main

import (
	"encoding/json"
	"fmt"
	"httpex/model"
	"net/http"
	"sync"
	"time"
)

func getJoke(i int, m *sync.Mutex) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")

	var chuck model.ChuckNorris
	err = json.NewDecoder(resp.Body).Decode(&chuck)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(chuck)
	m.Lock()
	jokes[i] = chuck
	m.Unlock()
	wg.Done()
}

func getJoke2(i int) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")

	var chuck model.ChuckNorris
	err = json.NewDecoder(resp.Body).Decode(&chuck)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(chuck)
	maap.Store(i, chuck)
	wg.Done()
}

var jokes map[int]model.ChuckNorris
var wg = sync.WaitGroup{}
var m sync.Mutex

var maap sync.Map

func main() {
	jokes = make(map[int]model.ChuckNorris, 25)
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()
	wg.Add(25)
	var i int
	for i = 0; i < 25; i++ {
		go getJoke2(i)
	}

	wg.Wait()
	fmt.Println(maap)
}
