//Enes Can Güneş
//Problem: https://www.hackerrank.com/challenges/staircase/problem
package main

import (
	"fmt"
	"strings"
)

//Example n = 4
//Expected:
// 	  #
// 	 ##
//  ###
// ####
func staircase(n int32) {
	intN := int(n)
	result := ""
	for i := 0; i < intN; i++ {
		line := ""
		line += strings.Repeat(" ", intN-1-i)
		line += strings.Repeat("#", i+1)
		result += line + "\n"
	}
	fmt.Println(result)
}
