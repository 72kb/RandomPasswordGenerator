package randomizer

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	lowercase    = "abcdefghijklmnopqrstuvwxyz"
	uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
	allChars     = lowercase + uppercase + digits + specialChars
)

func getRandomChar(charset string) byte {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	return charset[n.Int64()]
}

func GeneratePassword(length int) string {
	if length < 12 || length > 20 {
		length = 16 // Default length if out of range
	}

	var password strings.Builder

	// Ensure at least one of each type
	password.WriteByte(getRandomChar(lowercase))
	password.WriteByte(getRandomChar(uppercase))
	password.WriteByte(getRandomChar(digits))
	password.WriteByte(getRandomChar(specialChars))

	// Fill the rest randomly
	for i := 4; i < length; i++ {
		password.WriteByte(getRandomChar(allChars))
	}

	// Convert to string and shuffle to avoid predictable patterns
	shuffledPassword := shuffleString(password.String())

	return shuffledPassword
}

func shuffleString(input string) string {
	runes := []rune(input) // Convert string to rune slice (handles Unicode)
	for i := range runes {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(len(runes))))
		runes[i], runes[j.Int64()] = runes[j.Int64()], runes[i] // Swap characters
	}
	return string(runes) // Convert runes back to string
}
