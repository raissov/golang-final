package Builder

type agilityBuilder struct{
	damage int
	maxHealth int
}

func (a *agilityBuilder) setDamage()  {
	a.damage = 20
}


func (a *agilityBuilder) setMaxHealth()  {
	a.maxHealth = 800
}


func (a *agilityBuilder) getHero() HeroFinal {
	return HeroFinal{
		Damage:    a.damage,
		MaxHealth: a.maxHealth,
	}
}