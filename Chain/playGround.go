package Chain

import (
	"goFinal/Factory"
	_ "goFinal/Factory"
)

type playGround interface {
	askQuestion(*Factory.Hero)
	setNext(playGround)
}