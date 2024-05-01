package main

import (
	"fmt"
)

func main() {
	exponente := 3
	exponente2 := 3
	potencia(exponente, exponente2)

}

func potencia(val int, potenciaDe int) {
	if val > 10000 {
		return
	} else {
		fmt.Println(val)
		potencia(val*potenciaDe, potenciaDe)
	}
}
