package main

import "fmt"

//var estado bool

func main2() {
	estado := true
	if estado = false; estado == true {
		fmt.Println(estado)
	} else if estado == false {
		fmt.Println("soy false")
	}

}

func main() {
	switch numero := 5; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Mayor a 5")
	}

}
