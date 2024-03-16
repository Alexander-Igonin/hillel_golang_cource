package models


type Hydra struct {
	NumOfHeads int
	Shout string
}

func NewHydra(numOfHeads int) *Hydra {
	return &Hydra{
		NumOfHeads: numOfHeads,
	}
}

func (h *Hydra) LostHeads(amount int) {
	h.NumOfHeads -= amount
}
