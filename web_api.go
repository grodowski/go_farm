package main

import "encoding/json"
import "fmt"
import "log"
import "net/http"

type FarmAPI struct {
  Farm *Farm
}

func NewFarmAPI() *FarmAPI {
  f := new(FarmAPI)
  f.Farm = new(Farm)
  return f
}

func (f *FarmAPI) CreateAnimal(w http.ResponseWriter, req *http.Request) {
  dec := json.NewDecoder(req.Body)
  animal := new(Animal)
  if jsonErr := dec.Decode(animal); jsonErr == nil {
    fmt.Printf("%+v", animal)
    if farmErr := f.Farm.AddAnimal(*animal); farmErr == nil {
      w.WriteHeader(http.StatusCreated)
    } else {
      log.Printf("%+v", farmErr)
      w.WriteHeader(http.StatusBadRequest)
    }
  } else {
    log.Printf("%+v", jsonErr)
    w.WriteHeader(http.StatusBadRequest)
  }
}

func (f *FarmAPI) GetAnimals(w http.ResponseWriter, req *http.Request) {
  enc := json.NewEncoder(w)
  enc.Encode(f.Farm.Animals)
}

func (f *FarmAPI) Router() func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    log.Printf("%#v", req) // happy print debugging :)
    switch req.Method {
    case "POST":
      f.CreateAnimal(w, req)
    case "GET":
      f.GetAnimals(w, req)
    }
  }
}

func main() {
  api := NewFarmAPI()
  http.HandleFunc("/animals", api.Router())
  err := http.ListenAndServe(":3003", nil)
  if err != nil {
    log.Fatalln("ListenAndServe: ", err)
  }
}
