package main

import (
	"strings"
)

/**
@author Shitanford
@create 2021-03-11 17:31
*/

func main()  {

}

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}