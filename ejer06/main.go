package main

import (
	"fmt"
)

func main2() {
	var number int
	var boolean bool
	//fmt.Println(one(7))
	var secret int
	println("Ingrese numero secreto")
	fmt.Scan(&secret)
	number, boolean = itstwo(secret)
	fmt.Println(number)
	fmt.Println(boolean)
}

func main() {
	fmt.Println(wearealot(12, 43, 65, 78, 90, 32, 68, 21))
	fmt.Println(wearealot(12, 43, 78, 68, 21))
	fmt.Println(wearealot(12, 43, 65, 78, 68, 21))
	fmt.Println(wearealot(12, 48, 21))
	fmt.Println(wearealot(12))
}

func one(num int) (z int) {
	z = num * 2
	return
}

func itstwo(val int) (int, bool) {
	if val == 20 {
		return 1, true
	} else {
		return 2, false
	}
}

func wearealot(numbers ...int) int {
	var tot int = 0
	var value int
	var key int
	for key, value = range numbers {
		tot = tot + value
		fmt.Printf("Soy la numero %d key %d value \n", key, value)
	}
	return tot

}
