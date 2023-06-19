package common

import (
	"math/rand"
)

// Crea una key con el formato xxx-xxx-xxx
func CreateKey() string {
	const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const keyLength = 3

	key := ""

	for i := 0; i < keyLength; i++ {
		for j := 0; j < keyLength; j++ {
			key += string(letters[rand.Intn(len(letters))])
		}

		if i < keyLength - 1 {
			key += "-"
		}
	}

	return key
}
