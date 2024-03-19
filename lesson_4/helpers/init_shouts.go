package helpers


type Shouts struct {
	Shouts []string
}

func NewShouts(shouts []string) *Shouts {
	return &Shouts{
		Shouts: shouts,
	}
}
