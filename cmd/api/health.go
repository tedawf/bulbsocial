package main

import "net/http"

func (app *application) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
