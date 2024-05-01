package main

import (
	"fmt"
)

var num int
var descripcion string
var Calculo func(num1 int, num2 int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo num1 + num2 = %d \n", Calculo(10, 28))
	Calculo = func(num1 int, num2 int) int {
		return num1 * num2
	}

	fmt.Printf("Multiplico num1 * num2= %d \n", Calculo(10, 28))

	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Divido num1 / num2= %d \n", Calculo(15, 5))
	//var op func(num1 int, num2 int) int = operaciones(5, 9)
	//fmt.Printf("Resultado de operaciones %d", op(5, 9))
	opis()

	var tablaDel int = 3
	miTabla := tabla(tablaDel)
	i := 0
	for i < 10 {
		fmt.Println(miTabla())
		i++
	}

	araiis()

}

//closures

func tabla(valor int) func() int {
	var val int = valor
	var sec int = 0
	return func() int {
		//fmt.Println("Soy sec!", sec)
		var tab int = val * sec
		sec++
		return tab
	}
}

//func operaciones(num1 int, num2 int) func {
//return func() int {
//return num1 - num2

//}
//}

func opis() {
	resultado := func() int {
		var a int = 32
		var b int = 21

		return a + b
	}
	fmt.Println("Soy opis", resultado())
}

func araiis() {
	var coleccion [10]int = [10]int{7, 4, 6, 1, 5, 8, 9, 3, 2, 4}

	for i := 0; i < len(coleccion); i++ {
		fmt.Println(coleccion[i])

	}

	pedacito := coleccion[4:8]

	fmt.Println(pedacito)

	var matriz [10][5]int
	matriz[9][4] = 54
	matriz[5][2] = 23

	fmt.Println(matriz)

	miCreacion := make([]int, 5, 7)

	fmt.Println(len(miCreacion), cap(miCreacion))

	//for i := 0; i < len(miCreacion); i++ {
	//miCreacion = append(miCreacion, i)
	//}

	//fmt.Println(len(miCreacion), cap(miCreacion))

}
