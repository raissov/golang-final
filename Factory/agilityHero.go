package Factory

import (
	"bufio"
	"fmt"
	"goFinal/Builder"
	"goFinal/Facade"
	"os"
)

type agilityHero struct {
	Hero
	Facade.UpgradeInterface
}

func(s *agilityHero) GetHero() Hero {
	return s.Hero
}

func agilityHeroStats()  {
	agilityBuilder := Builder.GetBuilder("agility")
	director := Builder.NewDirector(agilityBuilder)
	agilityH := director.BuildHero()

	fmt.Printf("\nHero Damage: %d \n", agilityH.Damage)
	fmt.Printf("\nHero Maximum Health Capcity: %d \n", agilityH.MaxHealth)
}

func newAgilityHero() IHero {
	defer agilityHeroStats()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter hero name: ")
	heroName, _ := reader.ReadString('\n')
	return &agilityHero{
		Hero{
			name:      heroName,
			attribute: "agility",
			FirstEnemy: false,
			SecondEnemy: false,
			ThirdEnemy: false,
		},
	}
}
