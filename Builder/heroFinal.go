package Builder

type HeroFinal struct {
	Damage      int
	MaxHealth   int
}
func (h *HeroFinal) AddHP(hp int){
	h.MaxHealth = h.MaxHealth+hp
}
func (h *HeroFinal) AddDamage(damage int) {
	h.Damage = h.Damage + damage
}