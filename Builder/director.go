package Builder

type director struct {
	builder heroBuilder
}

func NewDirector(h heroBuilder) *director {
	return &director{
		builder: h,
	}
}

func (d *director) SetBuilder(h heroBuilder) {
	d.builder = h
}

func (d *director) BuildHero() HeroFinal {
	d.builder.setDamage()
	d.builder.setMaxHealth()
	return d.builder.getHero()
}