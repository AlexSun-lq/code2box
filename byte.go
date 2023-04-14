package code2box

import (
	"crypto/rand"
	"io"
)

func RandomBytes(num int64) []byte {
	b := make([]byte, num)
	io.ReadFull(rand.Reader, b)
	return b
}
