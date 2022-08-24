package main

import (
	"net/http"
	"os"
)

var BuildTime = "Unknown"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world! Build time: " + BuildTime))
	})
	http.ListenAndServe(os.Getenv("SCRATCHPAD_ADDRESS"), nil)
}
