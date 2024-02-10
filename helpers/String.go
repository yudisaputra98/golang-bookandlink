package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func Uid(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}
