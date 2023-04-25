package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Controller struct {
	persons *PersonRepository
}

func NewController(persons *PersonRepository) *Controller {
	return &Controller{persons: persons}
}

type Person struct {
	Name string
	Age  int
}

type PersonRepository []Person

type SaveRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/index":
		c.index(w, req)
	case "/create":
		c.create(w, req)
	default:
		c.index(w, req)
	}
}

func (c *Controller) create(w http.ResponseWriter, req *http.Request) {
	var r SaveRequest
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p := Person{
		Name: r.Name,
		Age:  r.Age,
	}
	*c.persons = append(*c.persons, p)
	fmt.Fprintf(w, "Create entity, data: %+v", r)
}

func (c *Controller) index(w http.ResponseWriter, req *http.Request) {
	if len(*c.persons) == 0 {
		http.Error(w, "No one person found", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Persons:")
	for _, p := range *c.persons {
		fmt.Fprintf(w, "%+v\n", p)
	}
}

func main() {
	repo := &PersonRepository{}
	c := NewController(repo)
	fmt.Println("Start serve...")
	log.Fatal(http.ListenAndServe("localhost:8080", c))
}
