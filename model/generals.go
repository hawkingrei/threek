package model

type General struct {
	Blood             int
	Name              string
	Shan              bool
	Skill             func(g *General)
	Additionaleffects []func(g General)
}

func (g *General) Wound() {
	g.Blood = g.Blood - 1
}

func (g *General) Restore() {
	g.Blood = g.Blood + 1
}

func (g *General) DoSkill(enemy *General) {
	g.Skill(enemy)
}

func (g *General) SetShan() {
	g.Shan = true
}

func (g *General) CancelShan() {
	g.Shan = false
}

func (g *General) IsShan() bool {
	return g.Shan
}

type Generals interface {
	// GetUser gets a user by unique Username.
	//DoSkil(enemy *General)
	Wound()
	Restore()
	//DoAdditionalEffects(index int, Generals interface{})

	SetShan()
	CancelShan()
	IsShan() bool
}
