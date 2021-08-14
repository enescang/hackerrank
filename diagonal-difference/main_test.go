package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Value    [][]int32
	Expected int32
}

func TestDiagonalDifference(t *testing.T) {
	var tests []TestStruct
	tests = append(tests, TestStruct{[][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}, 15})
	tests = append(tests, TestStruct{[][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}, 2})
	for _, v := range tests {
		actual := DiagonalDifference(v.Value)
		assert.Equal(t, v.Expected, actual)
	}
}
