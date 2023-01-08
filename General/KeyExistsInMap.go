package main

import "fmt"

func main() {

	myMap := make(map[string]int)
	myMap["a"] = 100
	myMap["b"] = 0
	myMap["c"] = 300

	fmt.Println(myMap)
	fmt.Println(myMap["b"])
	fmt.Println(myMap["d"])

	val, ok := myMap["a"]
	if ok {
		fmt.Println("It is there!")
		fmt.Println(val)
	} else {
		fmt.Println("Not there!")
	}
}
