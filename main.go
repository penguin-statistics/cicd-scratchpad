package main

import (
	"net/http"
	"os"
)

var (
	Version   = "Unknown"
	BuildTime = "Unknown"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world! This is version " + Version + " built at " + BuildTime))
	})
	http.ListenAndServe(os.Getenv("SCRATCHPAD_ADDRESS"), nil)
}
