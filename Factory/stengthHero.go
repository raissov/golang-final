package Factory

import (
	"bufio"
	"fmt"
	"goFinal/Builder"
	"goFinal/Facade"
	"os"
)

type strengthHero struct {
	Hero
	Facade.UpgradeInterface
}

func(s *strengthHero) GetHero() Hero {
	return s.Hero
}

func strengthHeroStats()  {
	strengthBuilder := Builder.GetBuilder("strength")
	director := Builder.NewDirector(strengthBuilder)
	strengthH := director.BuildHero()

	fmt.Printf("\nHero Damage: %d \n", strengthH.Damage)
	fmt.Printf("\nHero Maximum Health Capcity: %d \n", strengthH.MaxHealth)
}

func newStrengthHero() IHero {
	defer strengthHeroStats()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter hero name: ")
	heroName, _ := reader.ReadString('\n')
	return &strengthHero{
		Hero{
			name:      heroName,
			attribute: "strength",
			FirstEnemy: false,
			SecondEnemy: false,
			ThirdEnemy: false,
		},
	}
}

