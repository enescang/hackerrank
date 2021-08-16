//Enes Can Güneş
//Problem: https://www.hackerrank.com/challenges/plus-minus/problem
package main

import "fmt"

func plusMinus(arr []int32) {
	// Write your code here
	var positives float32
	var negatives float32
	var zeros float32
	for _, v := range arr {
		if v == 0 {
			zeros++
		}
		if v < 0 {
			negatives++
		}
		if v > 0 {
			positives++
		}
	}
	length := float32(len(arr))
	fmt.Println(positives / length)
	fmt.Println(negatives / length)
	fmt.Println(zeros / length)

}
