package main

import "fmt"

func myXOR(a,b int){
	fmt.Printf("Bitwise XOR of (%d and %d) is %d\n",a,b, a^b)
}
func myNOT(a,b int){
	fmt.Printf("Bitwise NOT of %d is %d\n",a ,^a)
	fmt.Printf("Bitwise NOT of %d is %d\n",b ,^b)
}
func myOR(a,b int){
	fmt.Printf("Bitwise OR of (%d and %d) is %d\n",a,b,a|b)
}
func mySHIFT(a,b int){
	fmt.Printf("Bitwise Shift-left of (%d and %d) is %d\n",a,b,a<<b)
	fmt.Printf("Bitwise Shift-right of (%d and %d) is %d\n",a,b,a>>b)
}
func myAND(a,b int){
	fmt.Printf("Bitwise AND of (%d and %d) is %d\n",a,b,a&b)
}
func main(){
	a := 67
	b := 76
	myOR(a,b)
 	myXOR(a,b)
 	myAND(a,b)
 	myNOT(a,b)
 	mySHIFT(a,b)
}