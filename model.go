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

type AnimalValidator struct {
  Animal *Animal
}

func (av *AnimalValidator) IsValid() error {
  error := new(AnimalValidationError)
  if av.Animal.Species < 1 || av.Animal.Species > 3 {
    error.AddMessage("Species is invalid")
  }
  if av.Animal.Name == "" {
    error.AddMessage("Name is invalid")
  }
  if av.Animal.Age < 0 {
    error.AddMessage("Age is invalid")
  }
  if len(error.Messages) > 0 {
    return error
  }
  return nil
}

type Species int
const (
  Dog Species = iota + 1
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

func (a *Animal) String() string {
  speciesName, _ := speciesNames[a.Species]
  return fmt.Sprintf("%s %s %d", speciesName, a.Name, a.Age)
}

func (f *Farm) AddAnimal(animal *Animal) error {
  validator := &AnimalValidator{animal}
  if err := validator.IsValid(); err != nil {
    return err
  }
  f.Animals = append(f.Animals, animal)
  return nil
}

