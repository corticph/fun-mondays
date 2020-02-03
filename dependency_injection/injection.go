package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %v", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Lorenzo")

}

func main() {
	myWriter := MyWriter{}
	Greet(os.Stdout, "Lorenzo")
	Greet(&myWriter, "Lorenzo")
	http.ListenAndServe(":8888", http.HandlerFunc(GreetHandler))
}

type MyWriter struct {
	s string
}

func (w *MyWriter) Write(p []byte) (n int, err error) {
	return 1, nil
}
