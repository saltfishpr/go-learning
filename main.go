// @file: main.go
// @date: 2020/9/11

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func inputExample() {
	inputReader := bufio.NewReader(os.Stdin)
	var input string
	var err error
	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			break
		}
		input = input[:len(input)-1]
		if len(input) == 0 {
			break
		}
		ss := strings.Split(input, " ")
		sort.Strings(ss)
		fmt.Println(strings.Join(ss, " "))
	}
}
