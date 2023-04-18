package random

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomUInt64(min, max uint64) uint64 {
	return min + uint64(rand.Int63n(int64(max-min+1)))
}

const s = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"

func RandomString(n int) string {
	var sb strings.Builder
	l := len(s)
	for i := 0; i < n; i++ {
		c := s[rand.Intn(l)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomUsername() string {
	return RandomString(9)
}

func RandomEmail() string {
	return RandomUsername() + "@email.com"
}
