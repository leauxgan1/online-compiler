package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"log"
	"net/http"
	"os"
	"strconv"
)

const PORT = 8080
const BASE = "/"
const COMPILER_PATH = "../compiler/"

type CompileRequest struct {
	Code string `json:"code"`
}

type CompileResponse struct {
	Output string `json:"output"`
	Error  string `json:"error"`
}


func CompilerResponseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"COMPILED CODE RESULT")
}

func CompilerRequestHandler(w http.ResponseWriter, r *http.Request) {
	// Receive code from user and validate it
	if(r.Method != http.MethodPost) {
		log.Printf("Expected POST http method, got: %s\n",r.Method)
		fmt.Fprintf(w,"Incorrect HTTP method for this endpoint\n")
		return
	}

	var req CompileRequest 
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w,"Invalid request body", http.StatusBadRequest)
		log.Printf("Request improperly formatted, exiting\n")
		return
	}

	// Create a temporary file to be compiled by our running container
	tmpFile, err := os.CreateTemp("","main-*.go")
	if err != nil {
		http.Error(w,"Unable to process file, temporary file could not be created", http.StatusInternalServerError);
		return 
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(req.Code)); err != nil {
		http.Error(w,"Unable to write to file", http.StatusInternalServerError);
		return 
	}
	tmpFile.Close()

	cmd := exec.Command("docker", "run", "--rm", "-v", fmt.Sprintf("%s:/usr/src/app/main.go", tmpFile.Name()), "go-compiler")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		resp := CompileResponse {
			Output: stdout.String(),
			Error: stderr.String(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := CompileResponse {
		Output: stdout.String(),
		Error: "",
	}
	json.NewEncoder(w).Encode(resp)

	// fmt.Fprintf(w,"SENT CODE TO BE COMPILED")
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
	if err != nil {
		log.Fatalf("Server failed to start: %s",err)
	}
}
