package util

import (
	"math/rand"
	"strconv"
)

func RandString()(re string)  {
	a := rand.Intn(50000)
	s := strconv.Itoa(a)
	re = s
	return
}
