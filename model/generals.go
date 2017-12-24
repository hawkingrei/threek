package model

type General struct {
	Blood             int
	Name              string
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

type generals interface {
	// GetUser gets a user by unique Username.
	DoSkil(enemy *General)
	Wound()
	Restore()
	DoAdditionalEffects(index int, generals interface{})
}
