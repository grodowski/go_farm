package main

import "net/http"
import "log"

type safeRequestHandler func(http.ResponseWriter, *http.Request)

type requestHandler func(http.ResponseWriter, *http.Request) error

type loggedResponse struct {
  http.ResponseWriter
  status int
}

func (l *loggedResponse) WriteHeader(status int) {
  l.status = status
  l.ResponseWriter.WriteHeader(status)
}

func logRequest(handler safeRequestHandler) safeRequestHandler {
  return func(w http.ResponseWriter, req *http.Request) {
    handler(w, req)
    log.Printf("%s %s", req.Method, req.RemoteAddr)
  }
}

func handleErrors(handler requestHandler) safeRequestHandler {
  return func(w http.ResponseWriter, req *http.Request) {
    if err := handler(w, req); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }
}
