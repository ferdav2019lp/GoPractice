package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan time.Duration)
	go bucle(canal)
	fmt.Println("Estoy aca")
	tiempoDiferencia := <-canal
	fmt.Println(tiempoDiferencia)

}

func bucle(channel chan time.Duration) {

	horaComienzo := time.Now()
	for i := 0; i < 100000000000; i++ {

	}
	horaFin := time.Now()
	channel <- horaFin.Sub(horaComienzo)

}
