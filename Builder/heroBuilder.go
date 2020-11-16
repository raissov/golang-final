package Builder

type heroBuilder interface {
	setDamage()
	setMaxHealth()
	getHero() HeroFinal
}

func GetBuilder(builderType string) heroBuilder  {
	if builderType == "agility" {
		return &agilityBuilder{}
	}
	if builderType == "strength" {
		return &strengthBuilder{}
	}
	return nil
}
