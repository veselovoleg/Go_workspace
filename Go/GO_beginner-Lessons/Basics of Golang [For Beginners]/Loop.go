package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += i
	}
	fmt.Println("Сумма равна ", sum)
	
}