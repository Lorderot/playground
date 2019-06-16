package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		username, password, ok := request.BasicAuth()
		if !ok {
			writer.Write([]byte("Hello!"))
			return
		}
		writer.Write([]byte(fmt.Sprintf("Hello, %s! I know your password! It contains %d digits!",
			username, len(password))))
	})
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		panic(err)
	}
}
