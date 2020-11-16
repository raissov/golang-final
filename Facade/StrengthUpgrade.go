package Facade

import (
	"goFinal/Builder"
	_ "goFinal/Builder"
	"goFinal/Factory"
)

type StrengthUpgrade struct {
	Factory.Hero
}
func (h *StrengthUpgrade) UpgradeStats(k *Builder.HeroFinal) {

	if h.GetFirstEnemy() {
		k.AddHP(200)
		k.AddDamage(5)
	}
	if h.GetSecondEnemy(){
		k.AddHP(250)
		k.AddDamage(10)
	}
	if h.GetThirdEnemy(){
		k.AddHP(300)
		k.AddDamage(15)
	}
}