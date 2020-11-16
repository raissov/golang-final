package Factory

type Hero struct {
	name string
	attribute string
	FirstEnemy  bool
	SecondEnemy bool
	ThirdEnemy  bool
}

func (h *Hero) setName(name string)  {
	h.name = name
}

func (h *Hero) GetName() string {
	return h.name
}

func (h *Hero) setAttribute(attribute string)  {
	h.attribute = attribute
}

func (h *Hero) GetAttribute() string {
	return h.attribute
}

func (h *Hero) SetFirstEnemy(status bool) {
	h.FirstEnemy = status
}

func (h *Hero) SetSecondEnemy(status bool) {
	h.SecondEnemy = status
}

func (h *Hero) SetThirdEnemy(status bool) {
	h.ThirdEnemy = status
}

func (h *Hero) GetFirstEnemy() bool  {
	return h.FirstEnemy
}

func (h *Hero) GetSecondEnemy() bool  {
	return h.SecondEnemy
}

func (h *Hero) GetThirdEnemy() bool  {
	return h.ThirdEnemy
}


