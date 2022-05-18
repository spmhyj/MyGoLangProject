package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const usage = `
Feet to Meters
--------------
This program converts feet into meters.
Usage:
feet [feetsToConvert]`

const usage2 = `
	Usage: 
	go run main age`

func main() {
	// run1()
	// run2()
	// run3()
	// run4()
	// run5()
	// run6()
	// run7()
	// run8()
	// run9()
	runa()
}

func runa() {

	const abc = 5

	var a [5]int
	a = [abc]int{1, 2, 3, 4, 5}

	b := [3]string{"aliso", "viejo", "ca"}

	c := 2.5

	fmt.Printf("%+#v\n", c)
	fmt.Println("##################")

	fmt.Printf("%v\n", a)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%+#v\n", a)

	fmt.Println("##################")
	fmt.Printf("%v\n", b)
	fmt.Printf("%+q\n", b)
	fmt.Printf("%#v == %[1]T\n", b)
	fmt.Printf("%+#v\n", b)
}

func run9() {

	argN := len(os.Args)
	if argN < 2 {
		fmt.Println("Give me the size of table")
		return
	}
	if argV, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println("Give me the size of table")
		return
	} else {
		fmt.Println(argV)
		fmt.Printf("%5s", "X")
		for i := 0; i <= int(argV); i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()
		for i := 0; i <= int(argV); i++ {
			fmt.Printf("%5d", i)
			for j := 0; j <= int(argV); j++ {
				fmt.Printf("%5d", i*j)
			}
			fmt.Println()
		}

	}

}

func run8() {
	var (
		a int
		b string
	)

	fmt.Scanf("%d,%s", &a, &b)
	fmt.Printf("[%d]-[%s]", a, b)
}

func run7() {
	for a := range [5]int{} {
		fmt.Println(a)
	}
	fmt.Println("###########")
	for b, _ := range [5]int{} {
		fmt.Println(b)
	}
	fmt.Println("###########")
	for c, d := range [5]int{} {
		fmt.Println(c, d)
	}
	fmt.Println("###########")
}

func run6() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	year := time.Now().Year()

	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	days, m := 28, strings.ToLower(os.Args[1])

	switch m {
	case "june", "september", "november":
		days = 30
	case "janunary", "march", "july", "august", "october", "december":
		days = 31
	case "feburary":
		if leap {
			days = 29
		}
	default:
		fmt.Printf("%q is not a month.\n", m)
		return
	}

	fmt.Printf("%q has %d days.\n", m, days)
}

func run5() {
	const (
		usage       = "Usage: [username] [password]"
		errUser     = "Access denied for %q.\n"
		errPwd      = "Invalid password for %q.\n"
		accessOK    = "Access granted to %q.\n"
		user, user2 = "jack", "inanc"
		pass, pass2 = "1888", "1879"
	)

	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	//
	// REFACTOR THIS TO A SWITCH
	//
	switch {
	case u != user && u != user2:
		fmt.Printf(errUser, u)
	case u == user && p == pass:
		fallthrough
	case u == user2 && p == pass2:
		{
			fmt.Printf(accessOK, u)
		}
	default:
		fmt.Printf(errPwd, u)
	}
}

func run4() {

	switch h := 14; {
	case h >= 6:
		fmt.Println("good morning")
	case h >= 12:
		fmt.Println("good afternoon")
	case h >= 18:
		fmt.Println("good evening ")
	default:
		fmt.Println("good night")
	}
	a := 911
	s := fmt.Sprintf("hex [%#x] \t binary [%[1]b] \t %[1]d \t oct [%#[1]o]", a)

	fmt.Println(s)
}

func run3() {
	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
		return
	}

	// Notice that:
	// I've intentionally created this solution as verbose
	// as I can.
	//
	// See the next exercise.

	var (
		msgNot string
	)
	if year%400 == 0 {
		// leap = true
		msgNot = ""
	} else if year%100 == 0 {
		// leap = false
		msgNot = "not"
	} else if year%4 == 0 {
		// leap = true
		msgNot = ""
	}

	fmt.Printf("%d is %s a leap year.\n", year, msgNot)

	// 	if leap {
	// 		fmt.Printf("%d is a leap year.\n", year)
	// 	} else {
	// 		fmt.Printf("%d is not a leap year.\n", year)
	// 	}
}

func run2() {
	if len(os.Args) < 2 {

		fmt.Println(usage2)
		return

	} else if age, err := strconv.Atoi(os.Args[1]); err != nil || age < 0 {
		fmt.Printf("age is [%d]\n", age)
		if age < 0 {
			fmt.Println("Age must be positive")
		} else {
			fmt.Printf("%s is invalid, return error %v", os.Args[1], err)
		}
	} else if age > 17 {
		fmt.Println("R-Rated")
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	} else if age < 13 {
		fmt.Println("PG-Rated")
	}
}

func run1() {
	//meters := feet * 0.3048
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}
	var (
		f   float64
		err error
	)

	if f, err = strconv.ParseFloat(os.Args[1], 64); err == nil {
		m := f * 0.3048
		fmt.Printf("%g feet is %g meters", f, m)
	} else {
		fmt.Printf("error: %q is not a number, - detail error %s", os.Args[1], err)
	}

	fmt.Println(f)
}
