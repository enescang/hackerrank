package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//How to capture stdouts: https://stackoverflow.com/a/29339052/11316913
func TestPlusMinus(t *testing.T) {
	std := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	plusMinus([]int32{-1, -2, -3, 0, 0, 0, 1, 2, 3})
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = std
	expected := "0.33333334\n0.33333334\n0.33333334\n"
	assert.Equal(t, expected, string(out))
}
