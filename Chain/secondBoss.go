package Chain

import "C"
import (
	"fmt"
	"goFinal/Factory"
)


type secondBoss struct {
	next playGround
}

func (s *secondBoss) askQuestion(h *Factory.Hero){
	if h.FirstEnemy{
		fmt.Print("Boss has fallen!")
		s.next.askQuestion(h)
		return
	}
	fmt.Println("Questions")
	h.FirstEnemy = true
	s.next.askQuestion(h)
}

func (s *secondBoss) setNext(next playGround){
	s.next = next
}