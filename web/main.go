package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFun)
	fmt.Println("Start listening to port 3000...")
	http.ListenAndServe(":3000", nil)
}

func handlerFun(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Hello Golang from local host...</h1>")
}
