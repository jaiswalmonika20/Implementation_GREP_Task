package main

import (
	"grepimplementation/search/Grep"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Search(t *testing.T) {
	t.Run("Not selecting flag with one file testing", Test_NoFlagSingleFile)
	t.Run("Not selecting flag with multiplr file", Test_NoFlagMultipleFile)
	t.Run("selecting -n flag for testing", Test_CaseOne)
	t.Run("selecting -l flag for testing", Test_CaseTwo)
	t.Run("selecting -i flag for testing", Test_CaseThree)
	t.Run("selecting -v flag for testing", Test_CaseFour)
	t.Run("selecting -x flag for testing", Test_CaseFive)
}

// not selecting a flag in single file
func Test_NoFlagSingleFile(t *testing.T) {
	actual := Grep.Search("hello", []string{}, []string{"input1.txt"})
	expected := []string{"input1.txt:hello", "input1.txt:hello", "input1.txt:hello"}
	assert.Equal(t, actual, expected, "the actual and expected result not match")
}

// not selecting a flag in single file
func Test_NoFlagMultipleFile(t *testing.T) {
	actual := Grep.Search("hello", []string{}, []string{"input1.txt", "input2.txt"})
	expected := []string{"input1.txt:hello", "input1.txt:hello", "input1.txt:hello", "input2.txt:hello"}
	assert.Equal(t, actual, expected, "the actual and expected result not match")
}

// for flag -n
func Test_CaseOne(t *testing.T) {
	actual := Grep.Search("hello", []string{"-n"}, []string{"input1.txt", "input2.txt"})
	expected := []string{"input1.txt:1-hello", "input1.txt:2-hello world", "input1.txt:3-hello world again", "input2.txt:1-hello"}
	assert.Equal(t, actual, expected, "the actual and expected result not match")
}

// for flag -l
func Test_CaseTwo(t *testing.T) {

	actual := Grep.Search("hello", []string{"-l"}, []string{"input1.txt", "input2.txt"})
	expected := []string{"input1.txt", "input2.txt"}
	assert.Equal(t, actual, expected, "the actual and expected result not match")
}

// for flag -i
func Test_CaseThree(t *testing.T) {
	actual := Grep.Search("heLlO", []string{"-i"}, []string{"input1.txt", "input2.txt"})
	expected := []string{"hello", "hello world", "hello world again", "HelloAGAin", "hello"}
	assert.Equal(t, actual, expected, "the actual and expected result not match")
}

// for flag -v
func Test_CaseFour(t *testing.T) {
	actual := Grep.Search("hello", []string{"-v"}, []string{"input1.txt", "input2.txt"})
	expected := []string{"HelloAGAin", "again", "hi my name", "my name is monika"}
	assert.Equal(t, actual, expected, "the actual and expected result not match")
}

// for flag -x
func Test_CaseFive(t *testing.T) {
	actual := Grep.Search("hello", []string{"-x"}, []string{"input1.txt", "input2.txt"})
	expected := []string{"hello", "hello"}
	assert.Equal(t, actual, expected, "the actual and expected result not match")
}
