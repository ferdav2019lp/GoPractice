package main

import (
	"fmt"
	"time"

	us "Practicas/ejer09/user"
)

type pepe struct {
	us.Usuario
}

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "jose", time.Now(), true)
	fmt.Println(u.Usuario)
}
