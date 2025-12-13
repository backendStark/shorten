package middleware

import "net/http"

type WrapperWiter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WrapperWiter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
