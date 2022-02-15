package celebrityReporter

import "fmt"

type ndtv struct {
}

func NewNdtv() ndtv {
	return ndtv{}
}

func (a ndtv) Broadcast(news string) {
	fmt.Println("NDTV news-", news)
}
