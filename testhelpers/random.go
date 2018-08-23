package testhelpers

import (
	"math/rand"
	"strconv"
)

func RandomInt() int64 {
	return rand.Int63()
}

func RandomString() string {
	return strconv.FormatInt(RandomInt(), 10)
}
