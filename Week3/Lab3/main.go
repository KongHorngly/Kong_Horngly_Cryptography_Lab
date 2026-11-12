package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/url"
	"regexp"
)

func hashMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {

	crackedTarget := "meowmeow"
	leakedURL := "https://google.com/?q=" + hashMD5(crackedTarget)

	parsed, err := url.Parse(leakedURL)
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}

	targetHash := parsed.Query().Get("q")
	if targetHash == "" {
		log.Fatal("[-] No q parameter found in URL. Example: https://google.com/?q=<md5>")
	}

	fmt.Printf("[+] Decoded MD5 Hash: %s\n", targetHash)

	wordlist := []string{
		"cat", "cats", "kitty", "kitten", "meow",
		"meowmeow", "purr", "purrpurr", "ilovecats", "cat123",
	}

	var crackedPassword string
	for _, word := range wordlist {
		if hashMD5(word) == targetHash {
			crackedPassword = word
			fmt.Printf("[+] Password Cracked: %s\n", crackedPassword)
			break
		}
	}

	if crackedPassword == "" {
		log.Fatal("[-] Password not found in wordlist.")
	}

	flag := fmt.Sprintf("cryptoCTF{%s}", crackedPassword)
	verifyPattern := `^cryptoCTF\{meowmeow\}$`

	matched, err := regexp.MatchString(verifyPattern, flag)
	if err != nil {
		log.Fatalf("Regex error: %v", err)
	}

	if matched {
		fmt.Printf("[+] Verified Flag:    %s\n", flag)
	} else {
		fmt.Println("[-] Flag did not match verification regex.")
	}
}
