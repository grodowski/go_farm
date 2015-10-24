package main

import "fmt"

type AnimalValidationError struct {
  Messages []string
}

func (e *AnimalValidationError) AddMessage(msg string) {
  e.Messages = append(e.Messages, msg)
}

func (e *AnimalValidationError) Error() string {
  return fmt.Sprint(e.Messages)
}

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

func (a *Animal) IsValid() (error, bool) {
  validation := new(AnimalValidationError)
  if a.Species < 1 || a.Species > 3 {
    validation.AddMessage("Species is invalid")
  }
  if a.Name == "" {
    validation.AddMessage("Name is invalid")
  }
  if a.Age < 0 {
    validation.AddMessage("Age is invalid")
  }
  if len(validation.Messages) > 0 {
    return validation, false
  } else {
    return nil, true
  }
}

func (f *Farm) AddAnimal(animal *Animal) error {
  e, _ := animal.IsValid()
  if e != nil {
    return e
  } else {
    f.Animals = append(f.Animals, animal)
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

