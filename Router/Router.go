package Router

import (
	"healthtracker/Handlers"
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Handlers.Index)
	mux.HandleFunc("/measurements", Handlers.Measurements)
	mux.HandleFunc("/measurements/", Handlers.Measurements)

	return mux
}
