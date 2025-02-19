package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	b := 0
	x, y := 15, 30

	fmt.Println(test(b))
	fmt.Println(test1(b))
	fmt.Printf("%T ", intToString(b))
	fmt.Println(intToString(b))
	fmt.Println(minInt(x, y))
}

func test(sum int) string {
	sum = 0
	for sum < 10 { //аналог while
		sum += 1
	}
	fmt.Println(sum)
	return ""
}

func test1(a int) string {
	a = 0
	for a < 1000 {
		if a == 100 {

		} else {
			fmt.Println("a is not 100")
		}
		a++
	}
	return ""
}

func intToString(a int) (s string) {
	s = strconv.Itoa(a)
	return
}

func minInt(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}
