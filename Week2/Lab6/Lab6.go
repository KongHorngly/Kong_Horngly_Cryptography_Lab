
package main

import "fmt"

func xorEncrypt(text string, key byte) string {
	result := ""

	for i := 0; i< len(text); i++ {
		result +=string(text[i] ^ key)
	}

	return result
}

func main() {
	var get_text string
	var get_key byte

	fmt.Print("Enter plaintext: ")
	fmt.Scanln(&get_text)

	fmt.Print("Enter key (one character): ")
	fmt.Scanf(" %c", &get_key)

	encrypted := xorEncrypt(get_text, get_key)					
	fmt.Println("Encrypted text:", []byte( encrypted ))	 // Encrypt message
	decrypted := xorEncrypt(encrypted, get_key)
	fmt.Println(" Decrypted text:", decrypted)			//decrypt the message
}