package main

import (
	"fmt"
	"strconv"
)

var numero int
var texto string
var status bool = true

func main() {
	numero2, numero3, numero4, texto, status := 2, 3, 4, "hola loco", false

	var moneda int64 = 0

	numero2 = int(moneda)

	texto = strconv.Itoa(numero2)

	//texto = fmt.Sprintf("%d", moneda)

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)

	MostrarStatus()

}

func MostrarStatus() {

	fmt.Println(status)
}
