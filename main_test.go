package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drpaij0se/api-deno-compiler/others"
)

func TestDownload(t *testing.T) {
	others.Download()
}

func TestAPI(t *testing.T) {
	url := "http://localhost:5000/code"
	payload := []byte(`{"code": "console.log('Hello, World!')"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Create a mock HTTP server to handle the request
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"ip":"8.8.8.8"}`)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	// Replace URL with the mock server's URL

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response: %v", err)
	}

	// print the response as json
	jsonResp, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		t.Fatalf("Error marshalling response: %v", err)
	}
	fmt.Println(string(jsonResp))
}

/*
const rawResponse = await fetch(
  "http://localhost:5000/code/tkwryp",
  {
    method: "GET",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
  }
);
const content = await rawResponse.json();
console.log(content);
*/
