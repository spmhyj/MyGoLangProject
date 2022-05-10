// repeat.go
package main

import (
	"fmt"
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
}

func run6() {
	// var amap map[string]string

	// amap := map[string]string{"a": "apple", "b": "Android"}
	amap := make(map[string]string)
	amap["a"] = "Apple"
	amap["b"] = "Android"

	fmt.Printf("%+v", amap)
}

func run5() {
	// s := fmt.Sprintf("%v%v%v", x, y, z)
	// fmt.Println(s)
	type myinttype int

	var x myinttype = 42

	fmt.Printf("type[%T]-value[%v]\n", x, x)
	y := int(x)
	fmt.Printf("type[%T]-value[%v]\n", y, y)
}

func run4() {
	var (
		a int
		b string
	)
	fmt.Print("Enter int, string: ")
	fmt.Scanf("%d,%s", &a, &b)
	fmt.Printf("a=[%d], b=[%s]\n", a, b)

}

// func run1() {
// 	var a [2]int
// 	for i := 0; i < len(a); i++ {
// 		fmt.Println(a[i])
// 	}
// 	var b [2]string
// 	for i := 0; i < len(b); i++ {
// 		fmt.Printf("[%s]\n", b[i])
// 	}
// }

func run2() {
	// got := []int{1, 3, 5}
	// want := []int{1, 3, 5, 2}
	fmt.Printf("sum: %d", run3(1, 3, 5))
}

func run3(nums ...int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}
