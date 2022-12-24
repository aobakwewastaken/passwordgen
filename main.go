package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	length := flag.Int("length", 16, "Length of password")
	characters := flag.String(
		"characters",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		"Allowed characters",
	)
	flag.Parse()

	password := generatePassword(*length, *characters)
	fmt.Println(password)

}

func generatePassword(length int, characters string) string {
	rand.Seed(time.Now().UnixNano())
	password := make([]rune, length)

	for i := range password {
		password[i] = rune(characters[rand.Intn(len(characters))])
	}

	return string(password)
}
