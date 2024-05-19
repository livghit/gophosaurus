package services

import "net/http"

func HandleStatic() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
}
