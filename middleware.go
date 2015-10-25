package main

import "net/http"

type safeRequestHandler func(http.ResponseWriter, *http.Request)

type requestHandler func(http.ResponseWriter, *http.Request) error

func handleErrors(handler requestHandler) safeRequestHandler {
  return func(w http.ResponseWriter, req *http.Request) {
    if err := handler(w, req); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }
}
