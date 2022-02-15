package celebrityReporter

import "fmt"

type celebrity struct {
	name         string
	age          uint8
	newsChannels []newsChannel
}

func NewCelebrity(name string, age uint8) *celebrity {
	return &celebrity{
		name: name,
		age:  age,
	}
}

func (c *celebrity) AddNewsChannel(othernewsChannel newsChannel) {
	c.newsChannels = append(c.newsChannels, othernewsChannel)
}

func (c *celebrity) Walk() {
	fmt.Println("celebrity walking")
}

func (c *celebrity) SpreadGossip(gossip string) {
	for _, v := range c.newsChannels {
		v.Broadcast(gossip)
	}
}
