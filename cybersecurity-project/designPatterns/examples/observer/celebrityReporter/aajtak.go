package celebrityReporter

import "fmt"

type aajTak struct {
}

func NewAajTak() aajTak {
	return aajTak{}
}

func (a aajTak) Broadcast(news string) {
	fmt.Println("AajTak news-", news)
}
