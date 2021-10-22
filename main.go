package main

import (
	"fmt"
)

//go:generate msgp

type QryD struct {
	T string
	K string
	V string
	A int8
	D int8
}

func main() {
	v1 := QryD{T: "table", K: "key", V: "value", A: 1, D: 9}
	v2 := QryD{T: "table", K: "key", V: "value", A: 2, D: 8}

	var bts []byte
	// 将 v1 与 v2 依次放入 bts 中
	bts, _ = v1.MarshalMsg(nil)
	bts, _ = v2.MarshalMsg(bts)
	fmt.Println(len(bts)) // 58

	var v3, v4 QryD
	// 从 bts 中取出一个 QryD
	bts, _ = v3.UnmarshalMsg(bts)
	fmt.Println(v3) // {1 9 table key value}
	// 再从 bts 中取出一个 QryD
	bts, _ = v4.UnmarshalMsg(bts)
	fmt.Println(v4) // {2 8 table key value}

	fmt.Println(len(bts)) // 0
}
