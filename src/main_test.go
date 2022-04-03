package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/ELPanaJose/api-deno-compiler/src/server"
)

func TestXxx(t *testing.T) {
	go server.StartServer()
	time.Sleep(500 * time.Millisecond)
	for i := 0; i < 10; i++ {
		start := time.Now()
		postBody, _ := json.Marshal(map[string]string{
			"code": "console.log(Deno.version)",
		})
		responseBody := bytes.NewBuffer(postBody)
		resp, err := http.Post("http://localhost:5000/code", "application/json", responseBody)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(body))
		duration := time.Since(start)

		log.Printf("API Response Time: %d%s\n", duration.Milliseconds(), "ms")
	}

}
