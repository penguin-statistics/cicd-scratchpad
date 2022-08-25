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
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "#UnknownHostname"
		}
		w.Write([]byte("Hello, world! This is version " + Version + " built at " + BuildTime + ", running on " + hostname))
	})
	http.ListenAndServe(os.Getenv("SCRATCHPAD_ADDRESS"), nil)
}
