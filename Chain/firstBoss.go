package Chain

import "C"
import (
	"fmt"
	"goFinal/Factory"
)


type FirstBoss struct {
	next playGround
}

func (f *FirstBoss) AskQuestion(h *Factory.Hero){
	if h.FirstEnemy{
		fmt.Print("Boss has fallen!")
		f.next.askQuestion(h)
		return
	}
	fmt.Println("Questions")
	h.FirstEnemy = true
	f.next.askQuestion(h)
}

func (f *FirstBoss) setNext(next playGround){
	f.next = next
}