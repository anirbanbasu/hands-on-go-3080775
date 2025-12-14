// interfaces/basics/begin/main.go
package main

import "fmt"

// define a humanoid interface with speak and walk methods returning string
type humanoid interface {
	speak()
	walk()
}

// define a person type that implements humanoid interface
type person struct{ name string }

func (p person) speak() {
	fmt.Printf("%s is speaking...\n", p.name)
}

func (p person) walk() {
	fmt.Printf("%s is walking...\n", p.name)
}

func (p person) String() string {
	return fmt.Sprintf("Person[name=%s]", p.name)
}

// implement the Stringer interface for the person type

// define a dog type that can walk but not speak
type dog struct{ name string }

func (d dog) walk() {
	fmt.Printf("%s, the dog, is walking...\n", d.name)
}

func main() {
	// invoke with a person
	p := person{name: "Alice"}
	doHumanThings(p)

	// can we invoke with a dog?
	doDogThings(dog{name: "Doxie"})

	fmt.Println(p)
}

func doHumanThings(h humanoid) {
	h.speak()
	h.walk()
}

func doDogThings(d dog) {
	d.walk()
}
