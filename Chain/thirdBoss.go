package Chain

import "C"
import (
	"fmt"
	"goFinal/Factory"
)


type thirdBoss struct {
	next playGround
}

func (t *thirdBoss) askQuestion(h *Factory.Hero){
	if h.FirstEnemy{
		fmt.Print("Boss has fallen!")
		t.next.askQuestion(h)
		return
	}
	fmt.Println("Questions")
	h.FirstEnemy = true
	t.next.askQuestion(h)
}

func (t *thirdBoss) setNext(next playGround){
	t.next = next
}