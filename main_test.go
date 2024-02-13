package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drpaij0se/api-deno-compiler/src/others"
)

func TestDownload(t *testing.T) {
	others.Download()
}

func TestAPI(t *testing.T) {
	url := "http://localhost:5000/code"
	payload := []byte(`{"code": "console.log(await fetch(\"https://ip-api.com/\"))"}`)

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

	fmt.Println("Response:", string(body))
}
