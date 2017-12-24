package general

import "github.com/hawkingrei/threek/model"

type Guanyu struct {
	g model.general
}

func GenerateGuanyu() (g Guanyu) {
	g.g.Name = "Guan yu"
	g.g.Blood = 4
	return
}

func (g *Guanyu) Wound() {
	g.g.Wound()
}

func (g *Guanyu) Restore() {
	g.g.Restore()
}

func (g *Guanyu) DoSkill(enemy *model.General) {
	g.g.Skill(enemy)
}
