package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main01() {
	fmt.Println("hello world!")
	a := "上海"
	fmt.Printf("length of %s is %d\n", a, utf8.RuneCountInString(a))
	fmt.Println(sayHello("Aliso Viejo"))

	fmt.Printf("args length %d, %+v\n", len(os.Args), os.Args)
	for _, v := range os.Args {
		fmt.Println(v)
	}

}

func main() {
	a := 12345
	fmt.Printf("(3d)%3d\n", a)
	fmt.Printf("(010d right pad)%010d\n", a)
	fmt.Printf("(0-10d, left pad)[%-010d]\n", a)
	b := 12.5
	fmt.Printf("%T\n", b)

	fmt.Println(Hello())
}
