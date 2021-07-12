package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s\n", name)
}

func main() {
	Greet(os.Stdout, "QUESO")
}

// Esto se puede usar ya para escribir valores en alguna pagina por el formato que tiene
// como usar un Fprintf directo, pero testeable

// func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
//     Greet(w, "world")
// }

// func main() {
//     log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
// }
