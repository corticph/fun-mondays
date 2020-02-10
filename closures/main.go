package main

import (
	"log"
	"net/http"
)

func logmw(fn func() error) {
	log.Println("Invoking something")
	fn()
}

func one(fn func()) {
	fn()
}

// type Animal interface {
// 	Speak()
// }

// type Dog struct {
// 	age int
// }

// func (dog *Dog) Speak() {
// 	log.Println("Bark:", dog.age)
// }

// type Human struct {
// 	age int
// }

// func (human *Human) Speak() {
// 	log.Println("Hello there:", human.age)
// }

// func hello(animal Animal) {
// 	animal.Speak()
// }

type HTTPFunction func(w http.ResponseWriter, r *http.Request)

func httpNextMW(fn func() error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		if err := fn(); err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
	}
}

func httpMW(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		fn(w, r)
	}
}

func httpLogMW(w http.ResponseWriter, r *http.Request, fn HTTPFunction) {
	log.Println(r.URL)
	fn(w, r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

var constValue = func() func() int {
	value := 0
	return func() int {
		value++
		return value
	}
}

func circularPointer(limit int) func() int {
	value := 0
	return func() int {
		value++
		if value == limit {
			value = 0
		}
		return value
	}
}

func main() {
	// var _ = func() {
	// 	log.Println("Hello from Lambda")
	// }

	// var somerr = func() error {
	// 	log.Println("inside the function")
	// 	return nil
	// }

	// logmw(somerr)

	// httpLogMW(w, r, indexHandler)

	// inc := constValue()

	// log.Println(inc())
	// log.Println(inc())
	// log.Println(inc())

	http.HandleFunc("/", httpMW(indexHandler))

	cp := circularPointer(2)

	log.Println(cp())
	log.Println(cp())
	log.Println(cp())
}
