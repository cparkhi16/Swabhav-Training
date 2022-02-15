package main

import "news/celebrityReporter"

func main() {
	xyz := celebrityReporter.NewCelebrity("xyz", 45)
	xyz.AddNewsChannel(celebrityReporter.NewAajTak())
	xyz.AddNewsChannel(celebrityReporter.NewNdtv())
	xyz.SpreadGossip("xyz going to airport")
}
