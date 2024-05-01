package main

import (
	"fmt"
	"log"
	"os"
)

func probarDefer(ruta string) {
	archivo, err := os.Open(ruta)

	defer func() {
		fmt.Println("Antes de irme cierro el archivo")
		archivo.Close()
	}()

	if err != nil {
		fmt.Println("Hubo error")
		//os.Exit(1) el exit me salta el defer!!
	} else {
		fmt.Println("archivo abierto correctamente")
	}

}

func probarPanic() {
	gato := true
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error de panic %v\n", reco)
		}
		//archivo.Close()
	}()
	if gato == true {
		//fmt.Println("Esto va a explotar")
		panic("Esto explota")
	}
}

func main() {
	//miRuta := "./miarchivolocal.txt"
	//probarDefer(miRuta)
	probarPanic()

}
