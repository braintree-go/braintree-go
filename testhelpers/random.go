package testhelpers

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"strconv"
)

func RandomInt() int64 {
	var n int64
	b := make([]byte, 8)
	rand.Read(b)
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.LittleEndian, &n)
	return n
}

func RandomString() string {
	return strconv.FormatInt(RandomInt(), 10)
}
