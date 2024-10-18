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

	SCRIPT := fmt.Sprintf("let msg = 1; Http.get(\"%s\", {timeout: 1000}).then(function(x) { msg = x.StatusText }); msg", server.URL)

	testScript(SCRIPT, asciiString("Hello, World!"), t)
}

func TestMyNewTest(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	SCRIPT := fmt.Sprintf("let msg, err; Http.get(\"%s\", {responseType: 'json'}).then(x => msg = x.Data.message).catch(e => err = e);  msg", server.URL)
	rt := New()
	v, _ := rt.RunScript("", SCRIPT)
	println(v.String())
	println(rt.Get("msg").toString().String())
	println(rt.Get("err").toString().String())
}

func TestHttpConfigTest(t *testing.T) {
	SCRIPT := `Http.get("https://baidu.com?abc=123", {
		params: {
			ttt: "abc efg"
		}
	})`

	testScript(SCRIPT, asciiString("Hello, World!"), t)
}
