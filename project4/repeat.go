// repeat.go
package main

import (
	"fmt"
	"os"
	"strings"
	// "reflect"
)

func Repeat(s string, cnt int) string {
	rep := ""
	for i := 0; i < cnt; i++ {
		rep += s
	}
	return rep
}

func sumArrayEle(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func sumArrayEle2(nums []int) int {
	sum := 0
	for _, number := range nums {
		sum += number
	}
	return sum
}

func sumAll(nums ...[]int) []int {
	var all []int
	for _, s := range nums {
		tempSum := sumArrayEle2(s)
		all = append(all, tempSum)
	}
	return all
}

var (
	x int    = 42
	y string = "James Bond"
	z bool   = true
)

func main() {
	fmt.Println("start TDD...")
	// run2()
	// run4()
	// run5()
	// run6()
	// run7()
	// run8()
	// run9()
	// run10()
	// run11()
	// run12()
	// run13()
	// run14()
	// run15()
	// run16()
	// runa()
	// runb()
	// runc()
	// rund()
	// rune()
	runf()
}

func runf() {
	const (
		usage    = "Usage: [username] [password]"
		errUser  = "Access denied for %q. \n"
		errPwd   = "Invalid passwd for %q.\n"
		accessOK = "Access granted to %q. \n"
	)

	users := map[string]string{"jack": "1888", "inanc": "1879"}

	if len(os.Args) < 3 {
		fmt.Println(usage)
		return
	}

	u, p := os.Args[1], os.Args[2]

	passwd, found := users[u]

	if !found {
		fmt.Printf(errUser, u)
		return
	} else if found {
		if p != passwd {
			fmt.Printf(errPwd, u)
			return
		} else {
			fmt.Printf(accessOK, u)
		}
	}

}

func rune() {
	vowels := "aeiouyw"

	if len(os.Args) == 1 || len(os.Args[1]) > 1 {
		fmt.Println("give me a letter")
	} else {
		idx := strings.IndexAny(vowels, os.Args[1])

		if idx >= 0 {
			if idx >= 5 {
				fmt.Printf("%q is sometimes a vowel, sometimes not\n", os.Args[1])
			} else {
				fmt.Printf("%q is a vowel\n", os.Args[1])
			}
		} else {
			fmt.Printf("%q is a consonant\n", os.Args[1])
		}

	}
}

// func rund() {
// 	n := len(os.Args) - 1

// 	switch n {
// 	case 0:
// 		fmt.Printf("Give me args")
// 	case 1:
// 		fmt.Printf("There is one: %+q", os.Args[1])
// 	case 2:
// 		fmt.Printf(`There is two: "%s %s" `+"\n", os.Args[1], os.Args[2])
// 	case 3:
// 		fmt.Printf("There is three: %+q", os.Args)
// 	case 4:
// 		fmt.Printf("There is four: %+q", os.Args)
// 	case 5:
// 		fmt.Printf("There is five: %+q", os.Args)
// 	}
// }

// func runc() {
// 	isSphere, radius := true, 200
// 	if isSphere && radius >= 200 {
// 		fmt.Println("It's a big sphere.")
// 	} else {
// 		fmt.Println("I don't know.")
// 	}
// }

// func runb() {
// 	age := 10
// 	if age >= 60 {
// 		fmt.Println("getting old")
// 	} else if age >= 30 && age < 60 {
// 		fmt.Println("getting wise")
// 	} else if age >= 20 && age < 30 {
// 		fmt.Println("getting adulthood")
// 	} else if age >= 10 && age < 20 {
// 		fmt.Println("Yound blood")
// 	} else {
// 		fmt.Println("Booting up")
// 	}

// }

// func runa() {
// 	tf := false
// 	fmt.Printf("Thess are %t claims", tf)
// 	t := 29.5
// 	fmt.Printf("Temperature is %.1f degrees\n", t)
// 	fmt.Printf("Temperature is %.1f degrees, and is normal %.2[1]f\n", t)
// 	s := "Hello World!"
// 	fmt.Printf("%q - %+[1]q", s)
// }

// func run16() {
// 	planet := "venus"
// 	distance := 261

// 	fmt.Printf("%v is %v away. Think! %[2]v kms! %[1]v OMG\n", planet, distance)
// }

// func run15() {

// 	// Now, use iota and initialize the following constants
// 	// "automatically" to 1, 2, and 3 respectively.
// 	const (
// 		_      = iota
// 		Spring = iota * 3
// 		Summer
// 		Fall
// 		Winter
// 	)

// 	// This should print: 1 2 3
// 	// Not: 0 1 2
// 	fmt.Println(Winter, Spring, Summer, Fall)
// }

// func run14() {
// 	const _ = iota

// 	const (
// 		a = iota
// 		b
// 		c
// 	)
// 	const (
// 		d = iota
// 		e
// 		f
// 	)

// 	fmt.Println(a, b, c, d, e, f)
// }

// func run13() {
// 	const (
// 		Nov = 11 - iota
// 		Oct
// 		Sep
// 	)

// 	fmt.Println(Sep, Oct, Nov)
// 	// const (
// 	// 	abc = iota + 10
// 	// 	_
// 	// 	def
// 	// 	ghi
// 	// )
// 	// fmt.Println(abc, def, ghi)
// }
// func run12() {
// 	// name := os.Args[1]
// 	// msg := `hi ` + name + `
// 	// How are you?`
// 	// fmt.Println(msg)
// 	json := `
// 		{
// 			"items": [{
// 				"item":{
// 					"name": "Teddy Bear"
// 				}
// 			}]
// 		}
// 	`
// 	fmt.Println(json)

// }

// func run11() {
// 	name := "inanç           "
// 	fmt.Println(utf8.RuneCountInString(strings.TrimSpace(name)))
// 	// 	msg := `

// 	// 	The weather looks good.
// 	// I should go and play.
// 	// 	`
// 	msg := `
// 	test   `

// 	fmt.Printf("[%s]", strings.TrimSpace(msg))
// }

// func run10() {
// 	n := len(os.Args[1])
// 	sRep := strings.Repeat("!", n)

// 	fmt.Printf("%s%s%s", sRep, strings.ToUpper(os.Args[1]), sRep)
// }

// func run9() {
// 	s := "I am 蒋颢"
// 	fmt.Printf("the length of [%s] is %d, and the number of chars %d", s, len(s), utf8.RuneCountInString(s))
// 	// fmt.Printf("Hi %s \n How are you? \n", os.Args[1])
// 	// fmt.Println(os.Args[0])
// 	// count := len(os.Args) - 1
// 	// fmt.Printf("There are %d names.\n", count)
// 	// fmt.Println(5.5)
// 	// a, b := 10, 5.5
// 	// a = int(b)
// 	// fmt.Println(float64(a) + b)
// }

// func run8() {

// 	var (
// 		planet string
// 		isTrue bool
// 		temp   float64
// 	)

// 	planet, isTrue, temp = "Mars", true, 19.5

// 	fmt.Printf("Air is good on %s\n", planet)
// 	fmt.Printf("It's %t\n", isTrue)
// 	fmt.Printf("It is %.1f degrees\n", temp)

// 	red, blue := "red", "blue"
// 	blue, red = red, blue
// 	fmt.Printf("%s %s", red, blue)

// 	// var (
// 	// 	lang    string
// 	// 	version int
// 	// )
// 	// lang, version = "go", 2
// 	// fmt.Println(lang, "version", version)
// 	// n := 0.
// 	// n = 3.14 * 2
// 	// fmt.Println(n)
// }

// func run7() {
// 	fmt.Printf("%+v\n", os.Args)
// }
// func run6() {
// 	// var amap map[string]string

// 	// amap := map[string]string{"a": "apple", "b": "Android"}
// 	amap := make(map[string]string)
// 	amap["a"] = "Apple"
// 	amap["b"] = "Android"

// 	fmt.Printf("%#v", amap)
// }

// func run5() {
// 	// s := fmt.Sprintf("%v%v%v", x, y, z)
// 	// fmt.Println(s)
// 	type myinttype int

// 	var x myinttype = 42

// 	fmt.Printf("type[%T]-value[%v]\n", x, x)
// 	y := int(x)
// 	fmt.Printf("type[%T]-value[%v]\n", y, y)
// }

// func run4() {
// 	var (
// 		a int
// 		b string
// 	)
// 	fmt.Print("Enter int, string: ")
// 	fmt.Scanf("%d,%s", &a, &b)
// 	fmt.Printf("a=[%d], b=[%s]\n", a, b)

// }

// // func run1() {
// // 	var a [2]int
// // 	for i := 0; i < len(a); i++ {
// // 		fmt.Println(a[i])
// // 	}
// // 	var b [2]string
// // 	for i := 0; i < len(b); i++ {
// // 		fmt.Printf("[%s]\n", b[i])
// // 	}
// // }

// func run2() {
// 	// got := []int{1, 3, 5}
// 	// want := []int{1, 3, 5, 2}
// 	fmt.Printf("sum: %d", run3(1, 3, 5))
// }

// func run3(nums ...int) int {
// 	sum := 0
// 	for _, i := range nums {
// 		sum += i
// 	}
// 	return sum
// }
