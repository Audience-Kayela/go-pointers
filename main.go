package main

import (
	"fmt"
)

func double(number *int) {
	*number *= 2

}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func main() {
	amount := 6
	double(&amount)
	fmt.Println(amount)

	truth := true
	negate(&truth)
	fmt.Println(truth)

}
