package log

import (
	"fmt"
	"sync"
)

var m sync.Mutex

type log struct {
	msg string
}

var logObj *log

func new(msg string) *log {
	return &log{
		msg: msg,
	}
}

func (l *log) GetMsg() string {
	return l.msg
}

func GetInstance(msg string) *log {
	m.Lock()
	//defer m.Unlock()
	fmt.Println(logObj)
	if logObj == nil {
		fmt.Println("created new log obj")
		logObj = new(msg)
	} else {
		fmt.Println("used last log obj")
		logObj.msg = msg
	}
	m.Unlock()
	return logObj
}
