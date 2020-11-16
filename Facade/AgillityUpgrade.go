package Facade

import (
	"goFinal/Builder"
	"goFinal/Factory"
)

type AgilityUpgrade struct {
	Factory.Hero
}
func (s *AgilityUpgrade) UpgradeStats(h *Builder.HeroFinal) {

	if s.GetFirstEnemy() {
		h.AddHP(100)
		h.AddDamage(15)
	}
	if s.GetSecondEnemy(){
		h.AddHP(150)
		h.AddDamage(20)
	}
	if s.GetThirdEnemy(){
		h.AddHP(200)
		h.AddDamage(25)
	}
}