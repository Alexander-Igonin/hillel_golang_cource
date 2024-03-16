package utils

import "math/rand"


func ReleaseShout(shouts []string) string {
	return shouts[rand.Intn(len(shouts))]
}
