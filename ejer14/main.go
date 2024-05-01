package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	pal := "un elefante se columbiaba sobre la tela de una arania"

	go deletreolento(pal)
	fmt.Println("en segundo plano esta deletreando")
	var t string
	fmt.Scanln(&t)
	//fmt.Println()

}

func deletreolento(palabra string) {

	letras := strings.Split(palabra, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}

}
