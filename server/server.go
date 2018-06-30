package main

import (
	"fmt"
	"log"
	"net/http"
)

func wasmHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request code")
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, "wasm/code.wasm")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/code.wasm", wasmHandler)
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
