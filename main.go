package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	http.ListenAndServe(os.Getenv("SCRATCHPAD_ADDRESS"), nil)
}
