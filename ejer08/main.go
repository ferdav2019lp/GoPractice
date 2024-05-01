package main

import "fmt"

func main() {
	campeonato := map[string]int{
		"boca":    32,
		"river":   32,
		"velez":   30,
		"newells": 34,
	}

	campeonato["lanus"] = 32

	delete(campeonato, "velez")

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s tiene %d puntos en el campeonato \n", equipo, puntaje)
	}
	equipo := "Rosario"
	puntaje, existe := campeonato[equipo]

	fmt.Printf("El equipo %s existe %t y tiene %d puntos", equipo, existe, puntaje)

}
