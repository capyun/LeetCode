package main

import (
	"bufio"
	"fmt"
)

/**
@author Shitanford
@create 2021-03-11 20:46
*/

func main()  {
	//ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
	data := []byte{1, 2, 3}
	for _ = range data {
		ad, token, _ := bufio.ScanBytes(data, false)
		fmt.Printf("%d, %Val", ad, token)
	}
}


