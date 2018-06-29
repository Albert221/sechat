package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var alphabet = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")

func RandomString(length int) string {
	output := make([]rune, length)
	for i := range output {
		output[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(output)
}