package main 

import "fmt"

func Calculate( number int ) (result int) {
	result = number+number	
	return result
}

func main(){ 

	fmt.Println("Testing package")
	fmt.Println(Calculate(3))

}