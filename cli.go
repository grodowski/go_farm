package main

import "encoding/json"
import "fmt"
import "os"

func loadAnimals(f *Farm) {
  for c := true; c; c = askContinue() {
    animal := new(Animal)
    fmt.Print("Species [1 - Dog, 2 - Cat, 3 - Cow]: ")
    fmt.Scanf("%d", &animal.Species)
    fmt.Print("Name: ")
    fmt.Scanf("%s", &animal.Name)
    fmt.Print("Age: ")
    fmt.Scanf("%d", &animal.Age)
    if e := f.AddAnimal(animal); e != nil {
      fmt.Println(e)
    } else {
      fmt.Println(animal)
    }
  }
}

func askContinue() bool {
  var cont string
  fmt.Print("Continue? [y/n]: ")
  fmt.Scanf("%s", &cont)
  return cont == "y"
}

func writeJSON(farm *Farm, filename string) {
  file, fileErr := os.Create(filename)
  if fileErr != nil {
    fmt.Println("Filesystem Error", fileErr)
    return
  }
  defer file.Close()
  if jsonErr := json.NewEncoder(file).Encode(farm); jsonErr != nil {
    fmt.Println("JSON Error", jsonErr)
  }
}

func main() {
  fmt.Println("Welcome to Uncle Doge's Farm 🐴")
  farm := new(Farm)
  loadAnimals(farm)
  fmt.Println(farm)
  writeJSON(farm, "farm.json")
}
