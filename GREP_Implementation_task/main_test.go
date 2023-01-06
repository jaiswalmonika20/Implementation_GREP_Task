package main

import (
	"fmt"
	"grepimplementation/search/Grep"
	"reflect"
	"testing"
)

func Test_Search(t *testing.T) {
	file := []string{"input1.txt", "input2.txt"}
	flags := []string{"-n", "-l", "-i", "-v", "-x"}

	//not selecting a flag
	fmt.Println("test case 1")
	if len(flags) == 0 {
		actual1 := Grep.Search("hello", []string{}, file)
		expected1 := []string{"hello", "input1.txt"}
		t.Errorf("actual result(%s) is not equal to expected result(%s)", actual1, expected1)
	}

	//for flag -n
	fmt.Println("test case 2")
	actual2 := Grep.Search("hello", []string{"-n"}, file)
	expected2 := []string{"input1.txt:1-hello", "input1.txt:2-hello world", "input1.txt:3-hello world again", "input2.txt:1-hello"}
	if reflect.DeepEqual(actual2, expected2) {
		t.Logf("actual result(%s) is equal to expected result(%s)", actual2, expected2)
	} else {
		t.Errorf("actual result(%s) is not equal to expected result(%s)", actual2, expected2)
	}

	//for flag -l
	fmt.Println("test case 3")

	actual3 := Grep.Search("hello", []string{"-l"}, file)
	expected3 := []string{"input1.txt", "input2.txt"}
	if reflect.DeepEqual(actual3, expected3) {
		t.Logf("actual result(%s) is equal to expected result(%s)", actual3, expected3)
	} else {
		t.Errorf("actual result(%s) is not equal to expected result(%s)", actual3, expected3)
	}

	//for flag -i
	fmt.Println("test case 4")
	actual4 := Grep.Search("heLlO", []string{"-i"}, file)
	expected4 := []string{"hello", "hello world", "hello world again", "HelloAGAin", "hello"}
	if reflect.DeepEqual(actual3, expected3) {
		t.Logf("actual result(%s) is equal to expected result(%s)", actual4, expected4)
	} else {
		t.Errorf("actual result(%s) is not equal to expected result(%s)", actual4, expected4)
	}

	//for flag -v
	fmt.Println("test case 5")
	actual5 := Grep.Search("hello", []string{"-v"}, file)
	expected5 := []string{"HelloAGAin", "again", "hi my name", "hi my name is monika"}
	if reflect.DeepEqual(actual3, expected3) {
		t.Logf("actual result(%s) is equal to expected result(%s)", actual5, expected5)
	} else {
		t.Errorf("actual result(%s) is not equal to expected result(%s)", actual5, expected5)
	}

	//for flag -x
	fmt.Println("test case 6")
	actual6 := Grep.Search("hello", []string{"-x"}, file)
	expected6 := []string{"hello", "hello"}
	if reflect.DeepEqual(actual3, expected3) {
		t.Logf("actual result(%s) is equal to expected result(%s)", actual6, expected6)
	} else {
		t.Errorf("actual result(%s) is not equal to expected result(%s)", actual6, expected6)
	}
}
