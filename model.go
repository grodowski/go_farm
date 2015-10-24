package main

import "fmt"
import "errors"

type Species int
const (
  Dog Species = iota
  Cat
  Cow
)

type Farm struct {
  Animals []*Animal
}

type Animal struct {
  Species Species
  Name string
  Age int
}

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

func (a Animal) IsValid() (error, bool) {
  var msg string
  if a.Species < Dog && a.Species > Cow { // WAT xD
    msg = "Species is invalid"
  }
  if a.Name == "" {
    msg = "Name is invalid"
  }
  if a.Age < 0 {
    msg = "Age is invalid"
  }
  if msg != "" {
    return errors.New(msg), false
  } else {
    return nil, true
  }
}

func (f *Farm) AddAnimal(animal Animal) error {
  e, _ := animal.IsValid()
  if e != nil {
    return e
  } else {
    f.Animals = append(f.Animals, &animal)
    return nil
  }
}

func (f *Farm) String() string {
  str := "Farm Summary\n"
  for index, element := range f.Animals {
    str += fmt.Sprintf("%d[Animal Species:%s, Name:%s, Age:%d]\n", index,
      element.Species.String(), element.Name, element.Age)
  }
  return str
}

