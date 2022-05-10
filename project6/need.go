package main

import (
	"fmt"

	"rsc.io/quote"
)

func sayHello(name string) string {
	var a, b, c = 100, "john", 10.8

	fmt.Printf("%d, %s, %f\n", a, b, c)

	return "Hello " + name + " ! from aliso viejo"
}

func Hello() string {
	return quote.Hello()
}
