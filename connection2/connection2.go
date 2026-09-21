// connection2 project connection2.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	fileServer := http.FileServer(http.Dir("./docHtml"))
	mux := http.NewServeMux()

	mux.Handle("/", fileServer)
	mux.HandleFunc("/form", handlesum)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}

func handlesum(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() %v", err)
		return
	}
	fmt.Fprintf(w, "Post request sucessfull")
	numberone := r.FormValue("numberone")
	numbertwo := r.FormValue("numbertwo")

	intnumberone, err := strconv.Atoi(numberone)
	if err != nil {
		panic(err)
	}
	intnumbertwo, err := strconv.Atoi(numbertwo)
	if err != nil {
		panic(err)
	}
	sum := intnumbertwo + intnumberone

	fmt.Fprintf(w, "Sum: %d", sum)
}
