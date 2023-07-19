package main

import "fmt"

// interface
type Person interface {
	Walk() string
}

// interface implementation
type PersonImpl struct {
	Name string
}

func (person *PersonImpl) Walk() string {
	return person.Name + "sedang berjalan"
}

// general function
func GoToHome(person Person) {
	fmt.Println(person.Walk(), "pulang")
}

func main() {
	joko := PersonImpl{Name: "Joko"}
	GoToHome(&joko)

	budi := &PersonImpl{Name: "Budi"}
	GoToHome(budi)
}
