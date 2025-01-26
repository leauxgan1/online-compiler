package main

import (
	"fmt"
	"log"
	"strconv"
	"net/http"
)

const PORT = 8080
const BASE = "/"

func CompilerResponseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"COMPILED CODE RESULT")
}

func CompilerRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"SENT CODE TO BE COMPILED")
}

func CompilerRunHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"SENT RUN EVENT TO COMPILER")
}

func main() {
	http.HandleFunc("/result",CompilerResponseHandler)
	http.HandleFunc("/compile",CompilerRequestHandler)
	http.HandleFunc("/run",CompilerRunHandler)
	log.Printf("Started server running on localhost:%d%s",PORT,BASE)
	err := http.ListenAndServe(":" + strconv.Itoa(PORT),nil)
	if(err != nil) {
		log.Fatalf("Server failed to start: %s",err)
	}
}
