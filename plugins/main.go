package main

import (
	"fmt"
	"net/http"
	"plugin"
	"strconv"
	"time"

	"github.com/Pungyeon/plugin/model"
)

func loadPlugin(path string) model.Processor {
	plug, err := plugin.Open(path)
	if err != nil {
		panic(err)
	}
	proc, err := plug.Lookup("Processor")
	if err != nil {
		panic(err)
	}
	var ok bool
	processor, ok = proc.(model.Processor)
	if !ok {
		panic("oh no")
	}
	return processor
}

func reload(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("value")
	value, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(err)
	}
	loadPlugin("./adder/adder.o")
	processor.WithOptions(model.Options{
		Value: int(value),
	})
}

func change(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	loadPlugin(path)
	w.WriteHeader(http.StatusOK)
}

var processor model.Processor

func main() {
	loadPlugin("./adder/adder.o")
	values := []int{1, 2, 3}
	fmt.Println(values)

	go func() {
		for {
			fmt.Println(processor.Process(values))
			time.Sleep(time.Second)
		}
	}()

	http.HandleFunc("/change", change)
	http.HandleFunc("/reload", reload)
	http.ListenAndServe(":8080", nil)
}
