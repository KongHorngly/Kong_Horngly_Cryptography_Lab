package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"Lab1/utils/crack"
)

func main() {
	targetHash := strings.ToLower("6a85dfd77d9cb35770c9dc6728d73d3f")
	wordlistPath := filepath.Join(".", "nord_vpn.txt")

	file, err := os.Open(wordlistPath)
	if err != nil {
		log.Fatalf("Failed to open wordlist at %q: %v", wordlistPath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		password := strings.TrimSpace(scanner.Text())
		if password == "" {
			continue
		}

		hashedAttempt := strings.ToLower(crack.HashMD5(password))

		fmt.Printf("[Line %d] Trying: %s -> %s\n", lineNumber, password, hashedAttempt)

		if hashedAttempt == targetHash {
			fmt.Println("\n========================================")
			fmt.Printf("[+] SUCCESS! Password found: %s\n", password)
			fmt.Printf("[+] Target Hash: %s\n", targetHash)
			fmt.Println("========================================")
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Println("\n[-] Password not found in the wordlist.")
}
