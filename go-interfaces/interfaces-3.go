package main

type Employee interface {
	Language() string
	Age() int
	Random() (string, error)
}

type Engineer struct {
	Name string
}

func (e Engineer) Language() string {
	return e.Name + "programs in Go."
}

func (e Engineer) Age() int {
	if e.Name == "Elliot" {
		return 24
	} else {
		return 30
	}
}

func (e Engineer) Random() (string, error) {
	return "", "Go is"
}

func main() {

	// This will throw an error

	var programmers []Employee

	elliot := Engineer{Name: "Elliot"}
	// Engineer does not implement the Employee interface
	// you'll need to implement Age() and Random()
	programmers = append(programmers, elliot)
}
