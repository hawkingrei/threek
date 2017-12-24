package model

type general struct {
	Blood             int
	Name              string
	Shan              bool
	Skill             func(g *general)
	Additionaleffects []func(g general)
}

func (g *general) Wound() {
	g.Blood = g.Blood - 1
}

func (g *general) Restore() {
	g.Blood = g.Blood + 1
}

func (g *general) DoSkill(enemy *general) {
	g.Skill(enemy)
}

func (g *general) SetShan() {
	g.Shan = true
}

func (g *general) CancelShan() {
	g.Shan = false
}

func (g *general) IsShan() bool {
	return g.Shan
}

type Generals interface {
	// GetUser gets a user by unique Username.
	DoSkil(enemy *general)
	Wound()
	Restore()
	DoAdditionalEffects(index int, Generals interface{})

	SetShan()
	CancelShan()
	IsShan() bool
}
