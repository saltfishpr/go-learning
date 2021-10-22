// @file: fileinput.go
// @description:
// @author: SaltFish
// @date: 2020/08/25

// Package ch12 is chapter 12
package ch12

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
)

func FileInput() {
	_, filename, _, _ := runtime.Caller(1)
	datapath := path.Join(path.Dir(filename), "study/ch12/input.txt")
	inputFile, inputError := os.Open(datapath)
	if inputError != nil {
		fmt.Printf(
			"An error occurred on opening the inputfile\n" +
				"Does the file exist?\n" +
				"Have you got acces to it?\n",
		)
		return // exit the function on error
	}
	defer inputFile.Close() // 确保在程序退出前关闭该文件

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}
