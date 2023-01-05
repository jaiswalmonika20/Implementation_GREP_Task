package main

import (
	"fmt"
	"grepimplementation/search/Grep"
)

func main() {
	//fmt.Println("Grep implementation task")
	//fmt.Println("enter the pattern:")
	//reader := bufio.NewReader(os.Stdin)
	//pattern, _ := reader.ReadString('\n')

	//var flag []string
	//fmt.Println("flags choosen to check the pattern:")
	//flaginput, _ := reader.ReadString('\n')
	//flag = append(flag, flaginput)

	//var selectfile []string
	// fmt.Println("flies choosen to check the pattern:")
	// fileinput, _ := reader.ReadString('\n')
	// selectfile = append(selectfile, fileinput)
	// var output []string
	pattern := "hello"
	flag := []string{"-n", "-i"}
	selectfile := []string{"input1.txt", "input2.txt"}
	output := Grep.Search(pattern, flag, selectfile)
	fmt.Println(output)

}
