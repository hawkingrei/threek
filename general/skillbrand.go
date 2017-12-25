package general

import "github.com/hawkingrei/threek/model"

func Sha(g model.Generals) {
	g.Wound()
}

func Yao(g model.Generals) {
	g.Restore()
}

func Shao(g model.Generals) {
	g.SetShan()
}
