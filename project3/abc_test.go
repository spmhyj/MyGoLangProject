package integers

import (
	"fmt"
	"testing"
)

func Test_adder(t *testing.T) {
	sum := Add(5, 3)

	expect := 8
	if sum != expect {
		t.Errorf("expect %d, but got %d", expect, sum)
	}
}

func ExampleAdd() {
	sum := Add(6, 9)
	fmt.Println(sum)
	//Output: 15
}

func ExampleSubstractabc() {
	sum := Substract(2, 1)
	fmt.Println(sum)
	//Output: 2
}
