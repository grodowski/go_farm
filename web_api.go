package main

import "encoding/json"
import "log"
import "net/http"

type FarmAPIController struct {
  Farm *Farm
}

func NewFarmAPIController() *FarmAPIController {
  return &FarmAPIController{new(Farm)}
}

func (f *FarmAPIController) index(w http.ResponseWriter, req *http.Request) error {
  return json.NewEncoder(w).Encode(f.Farm.Animals)
}

func (f *FarmAPIController) create(r http.ResponseWriter, req *http.Request) error {
  dec := json.NewDecoder(req.Body)
  animal := new(Animal)
  if jsonErr := dec.Decode(animal); jsonErr != nil {
    return jsonErr
  }
  if farmErr := f.Farm.AddAnimal(animal); farmErr != nil {
    return farmErr
  }
  return nil
}

func (f *FarmAPIController) routes() safeRequestHandler {
  return func (w http.ResponseWriter, req *http.Request) {
    res := &loggedResponse{ResponseWriter: w, status: 200}
    switch req.Method {
    case "POST":
      logRequest(handleErrors(f.create))(res, req)
    case "GET":
      logRequest(handleErrors(f.index))(res, req)
    default:
      http.Error(w, "No Route Found", http.StatusNotFound)
    }
  }
}

func main() {
  api := NewFarmAPIController()
  http.HandleFunc("/animals", api.routes())
  log.Fatalf("Error starting server: %v", http.ListenAndServe(":3003", nil))
}
