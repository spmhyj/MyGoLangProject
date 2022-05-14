// hello.go
package main

import (
	"fmt"
)

const LEVEL = 5

const (
	ENGLISH = iota + LEVEL
	SPANISH
	FRRENCH
	CHINESE
)

const ENGPREFIX = "Hello "
const SPANISHPREFIX = "Hola "
const FRENCHPREFIX = "Bonjour "
const french = "french"
const spanish = "spanish"

func sayHello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	prefix := ENGPREFIX
	switch lang {
	case french:
		prefix = FRENCHPREFIX
	case spanish:
		prefix = SPANISHPREFIX
	}

	return prefix + name
}
func tForLoop() {
	var i int
	for i <= 3 {
		fmt.Println(i)
		i++
	}
}

func finitArray() {
	d := []int{2: 10, 5: 25}

	fmt.Printf("%+v", d)
}

func fDeclareStruct() {
	type person struct {
		fname string
		lname string
	}
	a := person{"aliso", "viejo"}
	fmt.Printf("%+v", a)
}

func fDeclareMap() {
	a := map[string]string{
		"fname": "Aliso",
		"lname": "Viejo",
		"age":   "10",
	}
	fmt.Printf("%+v", a)
}
func main() {
	// fmt.Println(sayHello("Chris", "english"))
	// tForLoop()
	// finitArray()
	// fDeclareStruct()
	// fDeclareMap()
	fmt.Println(ENGLISH, SPANISH, CHINESE)
}
