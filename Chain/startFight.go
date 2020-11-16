package Chain

func StartFight() {
	first := &FirstBoss{}
	second := &secondBoss{}
	third := &thirdBoss{}

	first.setNext(second)
	second.setNext(third)

}
