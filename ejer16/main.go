package main

import (
	"fmt"
	"net/http"
)

func main() {
	//aca van la ruta con "/"
	http.HandleFunc("/", home)
	fmt.Printf("Starting server at port 3000\n")
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./index.html")

}
