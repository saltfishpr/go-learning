// @file: read_csv.go
// @description:
// @author: SaltFish
// @date: 2020/08/26

// Package ch12 is chapter 12
package ch12

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

type Book struct {
	title    string
	price    float64
	quantity int
}

func ReadCSV() {
	bks := make([]Book, 1)
	_, filename, _, _ := runtime.Caller(1)
	datapath := path.Join(path.Dir(filename), "study/ch12/input.csv")
	file, err := os.Open(datapath)
	if err != nil {
		log.Fatalf("Error %s opening file products.txt: ", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		// read one line from the file:
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// remove \r and \n so 2(in Windows, in Linux only \n, so 1):
		line = string(line[:len(line)-2])
		// fmt.Printf("The input was: -%s-", line)

		strSl := strings.Split(line, ";")
		book := new(Book)
		book.title = strSl[0]
		book.price, err = strconv.ParseFloat(strSl[1], 32)
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		// fmt.Printf("The quan was:-%s-", strSl[2])
		book.quantity, err = strconv.Atoi(strSl[2])
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		if bks[0].title == "" {
			bks[0] = *book
		} else {
			bks = append(bks, *book)
		}
	}
	fmt.Println("We have read the following books from the file: ")
	for _, bk := range bks {
		fmt.Println(bk)
	}
}
