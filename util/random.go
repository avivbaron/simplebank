package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{USD, CAD, EUR}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail Generates emails
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
