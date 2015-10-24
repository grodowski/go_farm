package main

import "fmt"
import "encoding/json"
import "bufio"
import "os"

type Species int
const (
  Dog Species = iota
  Cat
  Cow
)

var speciesNames = map[Species]string{
  Dog: "Dog",
  Cat: "Cat",
  Cow: "Cow",
}

func (s Species) String() string {
  n, ok := speciesNames[s]
  if ok {
    return n
  } else {
    return "Unknown"
  }
}

type Animal struct {
  Species Species
  Name string
  Age int
}

type Farm struct {
  Members []*Animal
}

func (f *Farm) addAnimal(species Species, name string, age int) {
  f.Members = append(f.Members, &Animal{species, name, age})
}

func (f *Farm) String() string {
  str := "Farm Summary\n"
  for index, element := range f.Members {
    str += fmt.Sprintf("%d[Animal Species:%s, Name:%s, Age:%d]\n", index,
      element.Species.String(), element.Name, element.Age)
  }
  return str
}

func loadAnimals(f *Farm) {
  for {
    var (
      species Species
      name string
      age int
    )
    fmt.Print("Species [0 - Dog, 1 - Cat, 2 - Cow]: ")
    s, _ := fmt.Scanf("%d", &species)
    fmt.Print("Name: ")
    n, _ := fmt.Scanf("%s", &name)
    fmt.Print("Age: ")
    a, _ := fmt.Scanf("%d", &age)
    if n == 0 || s == 0 || a == 0 {
      fmt.Println("Missing input! Try again 🙈")
      cont := askContinue()
      if cont {
        continue
      } else {
        return
      }
    } else {
      f.addAnimal(species, name, age)
      cont := askContinue()
      if !cont {
        return
      }
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
  file, _ := os.Create(filename)
  writer := bufio.NewWriter(file)
  defer func() {
    writer.Flush()
    file.Close()
  }()
  enc := json.NewEncoder(writer)
  enc.Encode(farm)
}

func main() {
  fmt.Println("Welcome to Uncle Doge's Farm 🐴")
  farm := new(Farm)
  loadAnimals(farm)
  fmt.Println(farm)
  writeJSON(farm, "farm.json")
}
