package main

import "net/http"

func main() {
	var err error

	router := http.NewServeMux()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	if err = http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
