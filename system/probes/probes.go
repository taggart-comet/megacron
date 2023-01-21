package probes

import (
	"megacron/system/functions"
	"net/http"
	"os"
)

// Serve responds to kubernetes probes
func Serve() {

	functions.Log("Servin' probes..")

	http.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	http.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	http.ListenAndServe(":"+os.Getenv("SERVER_PROBES_PORT"), nil)
}
