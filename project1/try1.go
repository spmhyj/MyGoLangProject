package main

import "fmt"

func main0101() {
	var a int = 10
	fmt.Println(a + 10)
	fmt.Println("Hello from Aliso Viejo")
}

func main0102() {
	PI := 3.14
	var r float64 = 5
	fmt.Println(PI * r)
}
func mai0103n() {
	// a, b, c := "test", 13.7, 20

	d, e := 20, 10
	d, e = e, d
	fmt.Println(d, e)
}

func main0104() {
	a := 20
	b := 3.14159
	c := 'a'
	fmt.Printf("==%05d==\n", a)
	fmt.Printf("%f\n", b)
	fmt.Println(c)
	fmt.Printf("%c\n", c)
}

func main0105() {
	var a int
	fmt.Scan(&a)
	fmt.Printf("here[%d]\n", a)
}
func main0106() {
	var s1, s2 string

	fmt.Scan(&s1, &s2)
	fmt.Println(s1, s2)
}
func main() {
	// var a, b int
	var a int
	var b int
	fmt.Printf("%p-%p\n", &a, &b)
	fmt.Scan(&a, &b)
	// fmt.Scan(&a)
	// fmt.Scan(&b)
	fmt.Println(a, b)
	fmt.Printf("%p-%p\n", &a, &b)
}
