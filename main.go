package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	http.ListenAndServe(":9000", r)
	// http.HandleFunc("/", handlerFun)
	// fmt.Println("Start listening to port 3001...")
	// http.ListenAndServe(":3001", nil)
}

func handlerFun(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Hello Golang from local host...</h1>")
}
