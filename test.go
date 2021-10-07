package main

import "fmt"

/**
@author Shitanford
@create 2021-09-30 11:49
*/

func main() {
	x := make(map[string]int, 2)
	fmt.Println(x["a"])
	fmt.Printf("%f", 10e9)
}

func test(x map[string]int) {
	for i := 0; i < 10000; i++ {
		x[string(i)+"sxy"] = i
	}
}