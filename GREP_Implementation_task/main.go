package main

import (
	"fmt"
	"grepimplementation/search/Grep"
)

func main() {
	// fmt.Println("Grep implementation task")
	// fmt.Println("enter the pattern:")
	// reader := bufio.NewReader(os.Stdin)
	// pattern, _ := reader.ReadString('\n')

	// var flag []string
	// fmt.Println("flags choosen to check the pattern:")
	// flaginput, _ := reader.ReadString('\n')
	// for _, token := range strings.Split(flaginput, " ") {
	// 	flag = append(flag, token)
	// 	fmt.Println(flag)
	// }
	// var selectfile []string
	// fmt.Println("files choosen to check the pattern:")
	// fileinput, _ := reader.ReadString('\n')
	// for _, token := range strings.Split(fileinput, " ") {
	// 	selectfile = append(selectfile, token)
	// 	fmt.Println(selectfile)
	// }
	//var output []string
	pattern := "hello"
	flag := []string{"-n", "-i"}
	selectfile := []string{"input1.txt", "input2.txt"}
	output := Grep.Search(pattern, flag, selectfile)
	fmt.Println(output)

}
