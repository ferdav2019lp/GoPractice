package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func leer1(ruta string) {

	archivo, err := ioutil.ReadFile(ruta)
	if err != nil {
		fmt.Println("Hubo un error en mi archivo")
	} else {
		fmt.Println("ahora voy a leer mi archivo")
		fmt.Println(string(archivo))
	}

}

func main() {

	//leer1()
	//leer2()
	//grabar1()
	miRutaDeArch := "./miarchivo.txt"
	//miRutaDeArch2 := "./minuevoarchivo.txt"
	//leer2(miRutaDeArch)
	//leer1(miRutaDeArch2)
	grabar2(miRutaDeArch)

}

func leer2(ruta string) {
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Hubo un error en mi archivo")
	} else {
		//fmt.Println("ahora voy a leer mi archivo")
		//fmt.Println(string(archivo)
		scaner := bufio.NewScanner(archivo)
		for scaner.Scan() {
			registro := scaner.Text()
			fmt.Printf("Linea >" + registro + "\n")
		}
	}
	archivo.Close()
}

func grabar1(ruta string) {
	archivo, err := os.Create(ruta)
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		//archivo.WriteString("Este es mi nuevo archivo de prueba pisado")
		fmt.Fprintln(archivo, "Este es mi nuevo archivo de prueba superrepisado")
	}
	archivo.Close()

}

func grabar2(ruta string) {
	archivo, err := os.OpenFile(ruta, os.O_RDWR|os.O_APPEND, 7777)
	if err != nil {
		fmt.Println("Hubo error")
	} else {
		if appendx(archivo, "\nEsta es la linea mas top del mundo") == false {
			fmt.Println("Hubo error")
		} else {
			fmt.Println("Ya escribi 2")
		}

	}
	archivo.Close()
}

func appendx(arch *os.File, texto string) bool {
	_, err := arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo error")
		return false
	} else {
		fmt.Println("ya escribi 1")
		return true
	}
}
