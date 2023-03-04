package main

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Version   = "Unknown"
	BuildTime = "Unknown"

	HTTPStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "scratchpad_http_status",
			Help: "HTTP status codes",
		},
		[]string{"code"},
	)
)

func main() {
	prometheus.MustRegister(HTTPStatus)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "#UnknownHostname"
		}

		greetings := "Hello, world!"

		if r.URL.Query().Get("mockerror") != "" {
			w.WriteHeader(http.StatusInternalServerError)
			greetings = "Mocked error!"
			HTTPStatus.WithLabelValues("500").Inc()
		} else {
			w.WriteHeader(http.StatusOK)
			HTTPStatus.WithLabelValues("200").Inc()
		}

		w.Write([]byte(greetings + " This is version " + Version + " built at " + BuildTime + ", running on " + hostname))
	})

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(os.Getenv("SCRATCHPAD_ADDRESS"), nil)
}
