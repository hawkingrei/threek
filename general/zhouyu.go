package general

import "github.com/hawkingrei/threek/model"

type Zhouyu struct {
	g model.General
}

func GenerateZhouyu() (g Zhouyu) {
	g.g.Name = "Zhou yu"
	g.g.Blood = 3
	return
}

func (g Zhouyu) Wound() {
	g.g.Wound()
}

func (g Zhouyu) Restore() {
	g.g.Restore()
}

func (g Zhouyu) DoSkill(enemy *model.General) {
	g.g.Skill(enemy)
}

func (g Zhouyu) CancelShan() {
	g.g.CancelShan()
}

func (g Zhouyu) IsShan() bool {
	return g.g.IsShan()
}

func (g Zhouyu) SetShan() {
	g.g.SetShan()
}
