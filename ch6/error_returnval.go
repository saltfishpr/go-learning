// @file: error_returnval.go
// @description: 有命名的函数返回值
// @author: SaltFish
// @date: 2020/08/04

// Package ch6 is chapter 6
package ch6

import (
	"errors"
	"fmt"
	"math"
)

// mySqrt 计算一个 float64 类型浮点数的平方根
func mySqrt(f float64) (float64, error) {
	// return an error as second parameter if invalid input
	if f < 0 {
		return float64(math.NaN()), errors.New("i won't be able to do a sqrt of negative number")
	}
	// otherwise use default square root function
	return math.Sqrt(f), nil
}

// mySqrt2 计算一个 float64 类型浮点数的平方根，有默认返回值
// name the return variables - by default it will have 'zero-ed' values i.e. numbers are 0, string is empty, etc.
func mySqrt2(f float64) (ret float64, err error) {
	if f < 0 {
		// then you can use those variables in code
		ret = float64(math.NaN())
		err = errors.New("i won't be able to do a sqrt of negative number")
	} else {
		ret = math.Sqrt(f)
		// err is not assigned, so it gets default value nil
	}
	return // automatically return the named return variables ret and err
}

// ErrReturnVal is fun
func ErrReturnVal() {
	fmt.Print("First example with -1: ")
	ret1, err1 := mySqrt(-1)
	if err1 != nil {
		fmt.Println("Error! Return values are: ", ret1, err1)
	} else {
		fmt.Println("It's ok! Return values are: ", ret1, err1)
	}

	fmt.Print("Second example with 5: ")
	if ret2, err2 := mySqrt(5); err2 != nil {
		fmt.Println("Error! Return values are: ", ret2, err2)
	} else {
		fmt.Println("It's ok! Return values are: ", ret2, err2)
	}
	// named return variables:
	fmt.Println(mySqrt2(5))
}
