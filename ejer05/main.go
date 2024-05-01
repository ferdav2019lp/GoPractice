package main

import (
	"fmt"
)

func main() {

	i := 0
	for i < 10 {
		fmt.Printf("\n valor de i: %d", i)
		if i == 4 {
			fmt.Printf("soy el 4 jajaa")
			fmt.Printf("Le sumo a i 2 \n")
			i = i * 2
			continue
		}
		fmt.Println("continuo")
		fmt.Println("soy i", i)
		i++

	}
}
