package fsm

import (
	"github.com/sirupsen/logrus"
)

type Command struct {
	Username string
	Card     string
	Skill    string
}

type Fsms struct {
	CommadnChan chan Command
	ExitChan    chan int
	FsmMap      map[int64]Fsm
}

func (f *Fsms) Start() {
	for {
		select {
		case c := <-f.CommadnChan:
			logrus.Info(c)
		case <-f.ExitChan:
			break
		}
	}
}
