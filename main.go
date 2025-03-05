package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Define a command-line flag for password length
	length := flag.Int("length", 16, "Length of the generated password (12-20 characters)")
	flag.Parse()

	// Validate length
	if *length < 12 || *length > 20 {
		log.Println("Invalid length! Using default length: 16")
		*length = 16
	}

	// Generate password
	password := randomizer.generatePassword(*length)
	fmt.Println("Generated Password:", password)

	// Allow saving to file (optional)
	if len(flag.Args()) > 0 {
		filename := flag.Arg(0)
		err := os.WriteFile(filename, []byte(password), 0644)
		if err != nil {
			log.Fatalf("Failed to save password: %v", err)
		}
		fmt.Println("Password saved to:", filename)
	}
}
