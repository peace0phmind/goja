package goja

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response struct {
	Message string `json:"message"`
}

func startTestServer() *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := Response{Message: "Hello, World!"}
		json.NewEncoder(w).Encode(response)
	})

	return httptest.NewServer(handler)
}

func TestHttpGetContentType(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	SCRIPT := fmt.Sprintf("Http.get(\"%s\").Headers['Content-Type']", server.URL)

	testScript(SCRIPT, asciiString("application/json"), t)
}

func TestHttpGetData(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	SCRIPT := fmt.Sprintf("JSON.parse(Http.get(\"%s\").Data).message", server.URL)

	testScript(SCRIPT, asciiString("Hello, World!"), t)
}
