//Enes Can Güneş
//Problem: https://www.hackerrank.com/challenges/diagonal-difference/problem
package main

func DiagonalDifference(arr [][]int32) int32 {
	var leftToRight int32
	var rightToLeft int32
	var point int = 0

	for i, _ := range arr {
		leftToRight += arr[i][point]
		rightToLeft += arr[i][len(arr)-point-1]
		point += 1
	}
	if leftToRight-rightToLeft < 0 {
		return rightToLeft - leftToRight
	}
	return leftToRight - rightToLeft
}
