package code2box

import (
	"fmt"
	"testing"
)

func TestRandomBytes(t *testing.T) {
	fmt.Println(RandomBytes(10))
	fmt.Println(RandomBytes(100))
	fmt.Println(RandomBytes(1000))
}
