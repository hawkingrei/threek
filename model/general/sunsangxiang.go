package general

import "github.com/hawkingrei/threek/model"

type Sunsangxiang struct {
	g model.general
}

func GenerateSunsangxiang() (g Sunsangxiang) {
	g.g.Name = "Sang sangxiang"
	g.g.Blood = 3
	return
}

func (g *Sunsangxiang) Wound() {
	g.g.Wound()
}

func (g *Sunsangxiang) Restore() {
	g.g.Restore()
}

func (g *Sunsangxiang) DoSkill(enemy *model.General) {
	g.g.Skill(enemy)
}
