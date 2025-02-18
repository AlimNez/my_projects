package main

import "fmt"

func main() {
	b := 0
	fmt.Println(test(b))
	fmt.Println(test1(b))
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

//func ()  {}
