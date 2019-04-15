package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
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
	port, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Starting server at Port %d", port)
	http.HandleFunc("/", jsonResponse)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
