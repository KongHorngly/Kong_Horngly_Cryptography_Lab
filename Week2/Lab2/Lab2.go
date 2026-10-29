package main
import "fmt"

func main(){
	a := 67
	b := 76
	c := 72
	if a != b {
		fmt.Println("True")
	}else{
		fmt.Println("false")
	}
	if  a > 0  && b > 0 {
		fmt.Println("True")
	}else{
		fmt.Println("False")
	}
	if a > c || b > c{
		fmt.Println("True")
	}else{
		fmt.Println("False")
	}

}