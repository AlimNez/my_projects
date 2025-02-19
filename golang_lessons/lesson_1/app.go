package main

import "fmt"

func main() {
	var name, age = "Alim", 24
	c := fmt.Sprintf("my name is %s and i'm %d years old", name, age)
	fmt.Println(c)
}
