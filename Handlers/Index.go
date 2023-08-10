package Handlers

import (
	"healthtracker/Templater"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	err := Templater.Templater.ExecuteTemplate(w, "index.html", "")
	if err != nil {
		return
	}
}
