package general

import "github.com/hawkingrei/threek/model"

type Sunsangxiang struct {
	g model.General
}

func GenerateSunsangxiang() (g Sunsangxiang) {
	g.g.Name = "Sang sangxiang"
	g.g.Blood = 3
	return
}

func (g Sunsangxiang) Wound() {
	g.g.Wound()
}

func (g Sunsangxiang) Restore() {
	g.g.Restore()
}

func (g Sunsangxiang) DoSkill(enemy *model.General) {
	g.g.Skill(enemy)
}

func (g Sunsangxiang) CancelShan() {
	g.g.CancelShan()
}

func (g Sunsangxiang) IsShan() bool {
	return g.g.IsShan()
}

func (g Sunsangxiang) SetShan() {
	g.g.SetShan()
}
