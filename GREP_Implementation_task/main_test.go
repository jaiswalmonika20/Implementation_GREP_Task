package main

import (
	"grepimplementation/search/Grep"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Search(t *testing.T) {
	t.Run("Zero case testing", Test_CaseZero)
	t.Run("first case testing", Test_CaseOne)
	t.Run("second case testing", Test_CaseTwo)
	t.Run("third case testing", Test_CaseThree)
	t.Run("fourth case testing", Test_CaseFour)
}

// not selecting a flag
func Test_CaseZero(t *testing.T) {
	actual := Grep.Search("hello", []string{}, []string{"input1.txt"})
	expected := []string{"hello", "input1.txt"}
	assert.Equal(t, reflect.DeepEqual(actual, expected), "the actual and expected result not match")
}

// for flag -n
func Test_CaseOne(t *testing.T) {
	actual := Grep.Search("hello", []string{"-n"}, []string{"input1.txt"})
	expected := []string{"input1.txt:1-hello", "input1.txt:2-hello world", "input1.txt:3-hello world again", "input2.txt:1-hello"}
	assert.Equal(t, reflect.DeepEqual(actual, expected), "the actual and expected result not match")
}

// for flag -l
func Test_CaseTwo(t *testing.T) {

	actual := Grep.Search("hello", []string{"-l"}, []string{"input1.txt"})
	expected := []string{"input1.txt", "input2.txt"}
	assert.Equal(t, reflect.DeepEqual(actual, expected), "the actual and expected result not match")
}

// for flag -i
func Test_CaseThree(t *testing.T) {
	actual := Grep.Search("heLlO", []string{"-i"}, []string{"input1.txt"})
	expected := []string{"hello", "hello world", "hello world again", "HelloAGAin", "hello"}
	assert.Equal(t, reflect.DeepEqual(actual, expected), "the actual and expected result not match")
}

// for flag -v
func Test_CaseFour(t *testing.T) {
	actual := Grep.Search("hello", []string{"-v"}, []string{"input1.txt"})
	expected := []string{"HelloAGAin", "again", "hi my name", "hi my name is monika"}
	assert.Equal(t, reflect.DeepEqual(actual, expected), "the actual and expected result not match")
}

// for flag -x
func Test_CaseFive(t *testing.T) {
	actual := Grep.Search("hello", []string{"-x"}, []string{"input1.txt"})
	expected := []string{"hello", "hello"}
	assert.Equal(t, reflect.DeepEqual(actual, expected), "the actual and expected result not match")
}
