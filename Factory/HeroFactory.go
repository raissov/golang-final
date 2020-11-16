package Factory

func GetHero(heroAttribute string) (IHero) {
	if heroAttribute == "agility"{
		return newAgilityHero()
	}
	if heroAttribute == "strength"{
		return newStrengthHero()
	}
	return nil
}