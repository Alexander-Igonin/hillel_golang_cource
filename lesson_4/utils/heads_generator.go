package utils

import (
	"errors"
	"math/rand"
)


func InitHeads(min, max int) (int, error) {
	if min <= 0 || max <= 0 {
		return 0, errors.New("number can't be negative")
	}
	return rand.Intn(max-min) + min, nil
}

func RecoverHeads(posibility, damage int) int {
	var res int

	for i := 0; i < damage; i++ {
		switch {
		case posibility >= 0 && posibility <= 35:
			res += 0
		case posibility > 35 && posibility <= 70:
			res += 1
		case posibility > 70 && posibility <= 90:
			res += 2
		case posibility > 90 && posibility <= 100:
			res += 3
		}
	} 

	return res
}
