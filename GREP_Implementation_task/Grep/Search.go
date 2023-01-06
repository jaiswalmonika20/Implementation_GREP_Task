package Grep

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Search(pattern string, flag, selectfile []string) []string {
	output := []string{}
	for _, flagno := range flag {
		if len(flagno)==2{
		output = append(output, choosesingleflag(flagno, pattern, selectfile)...)
		} else{
			output=append(output,choosemultipleflag(pattern,flagno,file)...)
		}
	}
	return output
}

func filelines(files string) []string {
	file, _ := os.Open(files)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var content []string
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	file.Close()
	return content
}

func choosesingleflag(flag, pattern string, files []string) []string {
	result := []string{}
	if flag == "-n" {
		for _, fileno := range files {
			filecontent := filelines(fileno)
			for index, line := range filecontent {
				if strings.Contains(line, pattern) {
					linenumber := strconv.Itoa(index + 1)
					result = append(result, fileno+":"+linenumber+"-"+pattern)
				}
			}
		}
	}
	if flag == "-l" {
		for _, fileno := range files {
			var status int
			filecontent := filelines(fileno)
			for _, line := range filecontent {
				status = 0
				if strings.Contains(line, pattern) {
					status = 1
					break
				}
			}
			if status == 1 {
				result = append(result, fileno)
			}
		}
	}
	if flag == "-i" {
		for _, fileno := range files {
			filecontent := filelines(fileno)
			for _, line := range filecontent {
				if strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
					result = append(result, line)
				}
			}
		}
	}
	if flag == "-v" {
		for _, fileno := range files {
			filecontent := filelines(fileno)
			for _, line := range filecontent {
				if !strings.Contains(line, pattern) {
					result = append(result, line)
				}
			}
		}
	}
	if flag == "-x" {
		for _, fileno := range files {
			filecontent := filelines(fileno)
			for _, line := range filecontent {
				if pattern == line {
					result = append(result, line)
				}
			}
		}
	}

	return result
}

func choosemultipleflag(pattern string,flag,files []string,)[]string{
	result:=[]string{}
	mapflags:=map[string]int{"-n":0,"-l":0,"-i":0,"-v":0,"-x":0}
	for _ ,choosenflag := range flag{
		mapflags[choosenflag]=1
	}

	if (mapflags["-n"]==1 && mapflags["-i"]==1){
		for _, fileno := range files {
			filecontent := filelines(fileno)
			for index, line := range filecontent{
				if strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
				linenumber := strconv.Itoa(index + 1)
					result = append(result, fileno+":"+linenumber+"-"+pattern)

			}
		}
	} else if (mapflags["-l"] == 1 ; mapflags["-n"] == 1){
		for _, fileno := range files {
			filecontent := filelines(fileno)
			for index, line := range filecontent{
				status = 0
				if strings.Contains(line, pattern) {
					fmt.Println("bye")
					status = 1
					break
				}
			}
			if status == 1 {
				result = append(result, fileno)
			}
			
		}
	} else if (mapflags["-x"]==1 && mapflags["-n"]==1){
		for _, fileno := range files {
			filecontent := filelines(fileno)
			for index, line := range filecontent{
				if pattern == line {
					linenumber := strconv.Itoa(index + 1)
					result = append(result, fileno+":"+linenumber+"-"+pattern)
				}


			}
		}
	} else if (mapflags["-l"]==1 && mapflags["-i"]==1){
		for _, fileno := range files {
			filecontent := filelines(fileno)
			for index, line := range filecontent{
				status = 0
				if strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
					status = 1
					break
				}
			}
			if status == 1 {
				result = append(result, fileno)
			}
		}
	} else if (mapflags["-l"]==1 && mapflags["-x"]==1){
		for _, fileno := range files {
			filecontent := filelines(fileno)
			for index, line := range filecontent{
				status=0
				if pattern == line {
					status = 1
					break
				}
			}
			if status == 1 {
				result = append(result, fileno)
			}
		}
}else if (mapflags["-i"]==1 && mapflags["-x"]==1){
	for _, fileno := range files {
		filecontent := filelines(fileno)
		for index, line := range filecontent{
			if pattern == line{
				if strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
					result = append(result, line)
			}
		}
		}
	}
} else {
	result=append(result,[]string{})

}
return result
}
