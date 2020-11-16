package Factory

type IHero interface {
	setName(name string)
	setAttribute(attribute string)
	GetName() string
	GetAttribute() string
	SetFirstEnemy(FirstEnemy bool)
	SetSecondEnemy(SecondEnemy bool)
	SetThirdEnemy(ThirdEnemy bool)
	GetFirstEnemy() bool
	GetSecondEnemy() bool
	GetHero() Hero
}
