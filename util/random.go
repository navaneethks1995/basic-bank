package util

import (
	"math/rand"
)

// func init() {
// 	rand.Seed(time.Now().UnixN())
// }

func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return randomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD", "INR"}
	return currencies[rand.Intn(len(currencies))]
}
