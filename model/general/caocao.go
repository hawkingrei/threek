package general

import "github.com/hawkingrei/threek/model"

type Caocao struct {
	g model.General
}

func GenerateCaocao() (g Caocao) {
	g.g.Name = "Caocao"
	g.g.Blood = 3
	return
}

func (g *Caocao) Wound() {
	g.g.Wound()
}

func (g *Caocao) Restore() {
	g.g.Restore()
}

func (g *Caocao) DoSkill(enemy *model.General) {
	g.g.Skill(enemy)
}
