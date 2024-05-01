package main

import "fmt"

type serVivo interface {
	estaVivo() string
	vivir()
	cazar()
}

type humano interface {
	comer()
	//estaComiendo()
	coger()
	dormir()
	respirar()
	meditar()
	entrenar()
	sexo() string
}

type hombre struct {
	nombre     string
	edad       int
	esVaron    bool
	altura     int
	peso       float32
	respirando bool
	comiendo   bool
	meditando  bool
	durmiendo  bool
	entrenando bool
	cogiendo   bool
	vivo       bool
	cazando    bool
}

type mujer struct {
	hombre
}

type animal struct {
	edad       int
	raza       int
	peso       float32
	respirando bool
	durmiendo  bool
	comiendo   bool
	vivo       bool
	esMacho    bool
	cazando    bool
}

type perro struct {
	animal
}

// Implemento la interfaz humano en hombre
func (h *hombre) comer() {
	h.comiendo = true
}

func (h *hombre) coger() {
	h.cogiendo = true
}
func (h *hombre) dormir() {
	h.durmiendo = true
}
func (h *hombre) meditar() {
	h.meditando = true
}
func (h *hombre) respirar() {
	h.respirando = true
}
func (h *hombre) entrenar() {
	h.entrenando = true
}
func (h *hombre) sexo() string {
	if h.esVaron {
		return "Es varon"
	} else {
		return "es mujer"
	}
}

func (h *hombre) cazar() {
	h.cazando = true
	fmt.Println("Mi hombre es cazando")
}

//Implento interfaz serVivo en hombre

func (h *hombre) estaVivo() string {
	if h.vivo {
		return h.nombre + " esta vivo"
	} else {
		return h.nombre + " esta muerto"
	}

}

func (h *hombre) vivir() {
	h.vivo = true
	fmt.Println("Mi hombre esta viviendo")
}

func (m *mujer) vivir() {
	m.vivo = true
	fmt.Println(("Mi mujer esta viviendo"))
}

//Implemento interfaz serVivo en animal

func (a *animal) estaVivo() string {
	if a.vivo {
		return "mi aninal esta vivo"
	} else {
		return "mi animal esta muerto"
	}
}
func (a *animal) vivir() {
	a.vivo = true
	fmt.Println("Mi animal esta viviendo")
}
func (a *animal) comer() {
	a.comiendo = true

}
func (a *animal) cazar() {
	a.cazando = true
	fmt.Println("Mi animal esta cazando")
}

//Funcion a parte
func (a *animal) sexo() string {
	if a.esMacho {
		return "Es macho"
	} else {
		return "es hembra"
	}
}

//Getter & Setter de hombre

func (h *hombre) estaComiendo() string {
	if h.comiendo {
		return "mi hombre esta comiendo"
	} else {
		return "mi hombre no esta comiendo"
	}
}

//Metodo polimorfico de la interfaz

func humanoComiendo(hum humano) {
	hum.comer()
	fmt.Println("Mi humano esta comiendo")
}

func esServivo(serv serVivo) string {
	return serv.estaVivo()
}
func servivoCazando(serv serVivo) {
	serv.cazar()
}

//main
func main() {
	//perrito := new(perro)
	//perrito.vivir()
	juana := new(mujer)
	juana.nombre = "juana"
	juana.vivir()
	//juana.esVaron = true
	//perrito.vivo = true
	humanoComiendo(juana)
	//esServivo(perrito)
	//fmt.Println(esServivo(perrito))
	fmt.Println((esServivo(juana)))
	//servivoCazando(perrito)

	//perrito.cazar()

	fmt.Printf("el ser %s y %s y %s ", juana.estaVivo(), juana.sexo(), juana.estaComiendo())

}
