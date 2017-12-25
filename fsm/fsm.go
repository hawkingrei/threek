package fsm

import (
	"errors"
	"sync"

	general "github.com/hawkingrei/threek/general"
	"github.com/hawkingrei/threek/model"
)

type GameStatic = int

const (
	LoadStatic GameStatic = iota       
	LoyalistStatic         
	QuislingStatic   
	ThiefStatic            
)

type Fsm struct {
	mu sync.Mutex
	Static           GameStatic
	LordUsername     string
	LoyalistUsername string
	QuislingUsername string
	ThiefUsername    string
	Lord             model.Generals
	Loyalist         model.Generals
	Quisling         model.Generals
	Thief            model.Generals
}

func getGeneral(name string) (g model.Generals, err error) {
	switch name {
	case "caocao":
		g = general.GenerateCaocao()
		return g, err
	case "guanyu":
		g = general.GenerateGuanyu()
		return g, err
	case "sangsangxiang":
		g = general.GenerateSunsangxiang()
		return g, err
	case "zhouyu":
		g = general.GenerateZhouyu()
		return g, err
	}
	return g, errors.New("name error")
}

func NewFsm(LordUsername string,
	LoyalistUsername string,
	QuislingUsername string,
	ThiefUsername string,
	Lord string,
	Loyalist string,
	Quisling string,
	Thief string,
) (f Fsm, err error) {
	f.LordUsername = LordUsername
	f.LoyalistUsername = LoyalistUsername
	f.QuislingUsername = QuislingUsername
	f.ThiefUsername = ThiefUsername
	f.Lord, err = getGeneral(Lord)
	if err != nil {
		return f, err
	}
	f.Loyalist, err = getGeneral(Loyalist)
	if err != nil {
		return f, err
	}

	f.Quisling, err = getGeneral(Quisling)
	if err != nil {
		return f, err
	}

	f.Thief, err = getGeneral(Thief)
	if err != nil {
		return f, err
	}
	return f, err
}
