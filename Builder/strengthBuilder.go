package Builder

type strengthBuilder struct{
	damage int
	maxHealth int
}

func(s *strengthBuilder) setDamage()  {
	s.damage = 18
}

func(s *strengthBuilder) setMaxHealth()  {
	s.maxHealth = 1000
}


func (s *strengthBuilder) getHero() HeroFinal {
	return HeroFinal{
		Damage:    s.damage,
		MaxHealth: s.maxHealth,
	}
}