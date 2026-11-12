package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	
	"golang.org/x/crypto/sha3"
)


func proofMe(textA,textB string){


//	Hashing MD5
	hashMD5_A:= md5.Sum([]byte(textA))
	hashMD5_B := md5.Sum([]byte(textB))

if(hashMD5_A == hashMD5_B){
	fmt.Println("Hash(MD5): Output A / Output B => Match ");
}else{
	fmt.Println("Hash(MD5): Output A / Output B => Not Match");
}

//	Hashing Sha1
	hashSha_A:= sha1.Sum([]byte(textA))
	hashSha_B := sha1.Sum([]byte(textB))
if(hashSha_A == hashSha_B){
	fmt.Println("Hash(Sha1): Output A / Output B => Match ");
}else{
	fmt.Println("Hash(Sha1): Output A / Output B => Not Match ");
}

//	Hashing Sha256
	hashSha256_A:= sha256.Sum256([]byte(textA))
	hashSha256_B := sha256.Sum256([]byte(textB))
if(hashSha256_A == hashSha256_B){
	fmt.Println("Hash(Sha256): Output A / Output B => Match ");
}else{
	fmt.Println("Hash(Sha256): Output A / Output B => Not Match ");
}
//	Hashing Sha512
	hashSha512_A:= sha512.Sum512([]byte(textA))
	hashSha512_B := sha512.Sum512([]byte(textB))
if(hashSha512_A == hashSha512_B){
	fmt.Println("Hash(Sha512): Output A / Output B => Match ");
}else{
	fmt.Println("Hash(Sha512): Output A / Output B => Not Match ");
}
//	Hashing Sha3
	hashSha3_A := sha3.Sum256([]byte(textA))
	hashSha3_B := sha3.Sum256([]byte(textB))
if hashSha3_A == hashSha3_B {
		fmt.Println("Hash(Sha3-256): Output A / Output B => Match")
} else {
		fmt.Println("Hash(Sha3-256): Output A / Output B => Not Match")
}
}

func main() {
	var textA string;
	var textB string;
fmt.Print("Enter The textA: ")
	fmt.Scanln(&textA)
fmt.Print("Enter The textB: ")
	fmt.Scanln(&textB)

proofMe(textA,textB)

}
