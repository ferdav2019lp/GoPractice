package main

import (
	"fmt"
)

func main() {

	res := interceptor(suma)(8, 4)
	fmt.Println(res)
	res2 := interceptor(resta)(8, 4)
	fmt.Println(res2)
	res3 := interceptor(multiplicacion)(8, 4)
	fmt.Println(res3)
	res4 := interceptor(division)(8, 4)
	fmt.Println(res4)

}

func interceptor(f func(int, int) int) func(int, int) int {

	return func(a, b int) int {

		fmt.Println("Ahora va a hacerse la operacion")
		return f(a, b)

	}
}

func suma(a, b int) int {
	return a + b
}

func resta(a, b int) int {
	return a - b
}
func multiplicacion(a, b int) int {
	return a * b
}
func division(a, b int) int {
	return a / b
}
