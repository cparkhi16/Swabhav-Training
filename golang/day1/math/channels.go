package math

import (
	"sync"
	"time"
)

type Message struct {
	chats   []string
	friends []string
}

func GetFriendsList(ch chan *Message, wh *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	ch <- &Message{
		friends: []string{"RAJ",
			"RAM",
			"Ramesh",
		},
	}
	wh.Done()
}

func GetFriendsChat(ch chan *Message, wh *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	ch <- &Message{
		chats: []string{"Tom",
			"Jerry",
		},
	}
	wh.Done()
}
