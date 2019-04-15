package main

import (
	"net/http"
  "fmt"
)


func jsonResponse(rw http.ResponseWriter, req *http.Request) {
  response := []byte(`
      {
        "message": "Hello World!!"
      }
  `)

  defer func() {
      rw.Header().Set("Content-Type", "application/json")
      fmt.Fprint(rw, string(response))
  }()
}

func main() {
  http.HandleFunc("/", jsonResponse)
  http.ListenAndServe(":8080", nil)
}
