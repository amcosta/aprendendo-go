package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please inform <source-temperature> <target-temperature>")
	}

	arg1 := os.Args[1]
	arg2 := os.Args[2]

	fmt.Println(string(arg1[len(arg1)-1]), arg2)

	// var source interface{}
	// switch string(arg1[len(arg1)-1]) {
	// case "C":
	// 	source := celsius.Celsius{}
	// default:

	// 	os.Exit(1)
	// }

	// fmt.Println(source)
}
