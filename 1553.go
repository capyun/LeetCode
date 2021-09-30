package main

import "fmt"

/**
@author Shitanford
@create 2021-04-18 10:15
*/

var ate = map[int]int{}

func main() {
	var n int
	for {
		fmt.Scanf("%d", &n)
		fmt.Printf("%d\n", eatOrange(n))
	}
}

func eatOrange(n int) int {
	if n<=1 {
		return 1
	}
	if x, ok := ate[n]; ok {
		return x
	}
	n2 := eatOrange(n/2)+1+n%2
	n3 := eatOrange(n/3)+1+n%3
	if n2<n3 {
		ate[n] = n2
		return n2
	} else {
		ate[n] = n3
		return n3
	}
}