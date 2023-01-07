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
	str := ""

	if len(flag) == 0 && len(selectfile) == 1 {
		output = append(output, ZeroflagSinglefile(pattern, str, selectfile)...)
	} else if len(flag) == 0 && len(selectfile) != 1 {

		output = append(output, ZeroflagMultiplefile(pattern, str, selectfile)...)
	} else {
		for _, flagno := range flag {

			//
			if len(flagno) == 2 {
				output = append(output, singleflagmultiplefiles(flagno, pattern, selectfile)...)
			} else {
				output = append(output, multipleflagmultiplefiles(pattern, flagno, selectfile)...)
			}
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

func ZeroflagSinglefile(pattern, flag string, files []string) []string {
	result := []string{}
	for _, fileno := range files {
		filecontent := filelines(fileno)
		for _, line := range filecontent {
			if strings.Contains(line, pattern) {
				result = append(result, fileno+":"+pattern)
			}
		}
	}
	return result

}

func ZeroflagMultiplefile(pattern, flag string, files []string) []string {
	result := []string{}
	for _, fileno := range files {
		filecontent := filelines(fileno)
		for _, line := range filecontent {
			if strings.Contains(line, pattern) {
				result = append(result, fileno+":"+pattern)
			}
		}
	}
	return result

}

func singleflagmultiplefiles(flag, pattern string, files []string) []string {
	result := []string{}
	if flag == "-n" {
		for _, fileno := range files {
			filecontent := filelines(fileno)
			for index, line := range filecontent {
				if strings.Contains(line, pattern) {
					linenumber := strconv.Itoa(index + 1)
					result = append(result, fileno+":"+linenumber+"-"+line)
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

func multipleflagmultiplefiles(pattern, flag string, files []string) []string {
	result := []string{}
	mapflags := map[string]int{"n": 0, "l": 0, "i": 0, "v": 0, "x": 0}

	res1 := strings.Split(flag, "-")
	for _, choosenflag := range res1 {
		mapflags[choosenflag] = 1
	}

	if mapflags["n"] == 1 && mapflags["i"] == 1 {
		for _, fileno := range files {
			filecontent := filelines(fileno)
			for index, line := range filecontent {
				if strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
					linenumber := strconv.Itoa(index + 1)
					result = append(result, fileno+":"+linenumber+"-"+line)
				}

			}
		}
	} else if mapflags["l"] == 1 && mapflags["n"] == 1 {
		for _, fileno := range files {
			var status int
			filecontent := filelines(fileno)
			for _, line := range filecontent {
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
	} else if mapflags["x"] == 1 && mapflags["n"] == 1 {
		for _, fileno := range files {
			filecontent := filelines(fileno)
			for index, line := range filecontent {
				if pattern == line {
					linenumber := strconv.Itoa(index + 1)
					result = append(result, fileno+":"+linenumber+"-"+line)
				}

			}
		}
	} else if mapflags["l"] == 1 && mapflags["i"] == 1 {
		for _, fileno := range files {
			var status int
			filecontent := filelines(fileno)
			for _, line := range filecontent {
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
	} else if mapflags["l"] == 1 && mapflags["x"] == 1 {
		for _, fileno := range files {
			var status int
			filecontent := filelines(fileno)
			for _, line := range filecontent {
				status = 0
				if pattern == line {
					status = 1
					break
				}
			}
			if status == 1 {
				result = append(result, fileno)
			}
		}
	} else if mapflags["i"] == 1 && mapflags["x"] == 1 {
		for _, fileno := range files {
			filecontent := filelines(fileno)
			for _, line := range filecontent {
				if pattern == line {
					if strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
						result = append(result, line)
					}
				}
			}

		}
	} else {
		result = []string{}
	}
	return result
}
