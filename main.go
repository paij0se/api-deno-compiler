package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/acarl005/stripansi"
	"github.com/drpaij0se/api-deno-compiler/others"
	"github.com/drpaij0se/api-deno-compiler/others/database"
	"github.com/gorilla/mux"
	"github.com/zhexuany/wordGenerator"
)

type code struct {
	Code string
}

func postCode(w http.ResponseWriter, r *http.Request) {
	var inputCode code
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}

	json.Unmarshal([]byte(reqBody), &inputCode)

	input := inputCode.Code
	// check if the request is empty
	if input == "" {
		json.NewEncoder(w).Encode("Error,empty input")
	} else {
		client := database.ConnectToDB()
		id := database.DbInsert(client, input)
		fmt.Println("input:", input)
		program := wordGenerator.GetWord(12) + ".ts"
		noBackQuote := strings.ReplaceAll(program, "`", "p")
		f, err := os.Create(noBackQuote)
		if err != nil {
			fmt.Println("some error creating the archive", err)
		}
		defer f.Close()
		_, err2 := f.WriteString(input)
		if err2 != nil {
			fmt.Println("error writing the archive", err2)
		}
		// execute deno and kill the process
		fmt.Println("archive edited")
		var stdout, stderr bytes.Buffer
		cmd := exec.Command("sh", "-c", `
	./deno run --allow-net --no-check `+noBackQuote+`&`+` sleep 1;kill $! 2>&1`)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		peo := cmd.Run()
		if peo != nil {
			fmt.Println(err)
		}
		// capture the stderr and stdout
		executedOut := stdout.String() + stderr.String()
		noAnsii := stripansi.Strip(executedOut)
		coolOut := strings.ReplaceAll(noAnsii, "sh: 2: kill: No such process", "")
		// delete the archive
		err = os.Remove(noBackQuote)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("archive deleted")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		// send the id and the output
		json.NewEncoder(w).Encode(map[string]string{"id": id, "output": coolOut})
	}
}

func GetCode(w http.ResponseWriter, r *http.Request) {
	// get the id from the request
	vars := mux.Vars(r)
	id := vars["id"]
	// connect to the database
	client := database.ConnectToDB()
	code := database.DbGet(client, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(code)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "make a http post request, more info in https://github.com/drpaij0se/api-deno-compiler")
}

func main() {
	// Download Deno
	others.Download()
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", indexRoute)
	r.HandleFunc("/code/{id}", GetCode).Methods("GET")
	r.HandleFunc("/code", postCode).Methods("POST")
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "5000"
	}
	fmt.Printf("Api on port: %s", port)
	http.ListenAndServe(":"+port, r)
}
