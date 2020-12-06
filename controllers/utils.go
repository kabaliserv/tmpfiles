package controllers

import (
	"net/http"

	"github.com/kabaliserv/tmpfiles/modules/errcode"
)

func renderError(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	render := errcode.JSONMap(code).ToJSON()
	w.Write(render)
}
