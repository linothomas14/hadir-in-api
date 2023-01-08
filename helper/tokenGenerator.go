package helper

import (
	"crypto/rand"
	"fmt"
)

func TokenGenerator() string {
	b := make([]byte, 3)
	fmt.Println(b)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
