package main

import (
	"fmt"
	"time"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/en_CA"
	"github.com/go-playground/locales/fr"
	"github.com/go-playground/locales/nl"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
)

var universalTraslator *ut.UniversalTranslator

func main() {
	universalTraslator = ut.New(zh.New(), zh.New(), en.New(), en_CA.New(), nl.New(), fr.New())

	zh, _ := universalTraslator.GetTranslator("zh")
	en, _ := universalTraslator.GetTranslator("en")

	fmt.Println("zh Cardinal Plural Rules:", zh.PluralsCardinal())
	fmt.Println("zh Ordinal Plural Rules:", zh.PluralsOrdinal())
	fmt.Println("zh Range Plural Rules:", zh.PluralsRange())
	fmt.Println("en Cardinal Plural Rules:", en.PluralsCardinal())
	fmt.Println("en Ordinal Plural Rules:", en.PluralsOrdinal())
	fmt.Println("en Range Plural Rules:", en.PluralsRange())
	fmt.Println()

	if err := universalTraslator.Import(ut.FormatJSON, "translations"); err != nil {
		panic(err)
	}

	fmt.Println(zh.T("normal_test", "月球"))
	fmt.Println(en.T("normal_test", "moon"))
	fmt.Println()
	fmt.Println(zh.C("cardinal_test", 1, 0, zh.FmtNumber(1, 0)))
	fmt.Println(zh.C("cardinal_test", 2, 0, zh.FmtNumber(2, 0)))
	fmt.Println(zh.C("cardinal_test", 10456.25, 2, zh.FmtNumber(10456.25, 2)))
	fmt.Println(en.C("cardinal_test", 1, 0, en.FmtNumber(1, 0)))
	fmt.Println(en.C("cardinal_test", 2, 0, en.FmtNumber(2, 0)))
	fmt.Println(en.C("cardinal_test", 10456.25, 2, en.FmtNumber(10456.25, 2)))
	fmt.Println()
	fmt.Println(zh.O("ordinal_test", 1, 0, zh.FmtNumber(1, 0)))
	fmt.Println(zh.O("ordinal_test", 2, 0, zh.FmtNumber(2, 0)))
	fmt.Println(zh.O("ordinal_test", 3, 0, zh.FmtNumber(3, 0)))
	fmt.Println(zh.O("ordinal_test", 4, 0, zh.FmtNumber(4, 0)))
	fmt.Println(zh.O("ordinal_test", 10456.25, 0, zh.FmtNumber(10456.25, 0)))
	fmt.Println(en.O("ordinal_test", 1, 0, en.FmtNumber(1, 0)))
	fmt.Println(en.O("ordinal_test", 2, 0, en.FmtNumber(2, 0)))
	fmt.Println(en.O("ordinal_test", 3, 0, en.FmtNumber(3, 0)))
	fmt.Println(en.O("ordinal_test", 4, 0, en.FmtNumber(4, 0)))
	fmt.Println(en.O("ordinal_test", 10456.25, 0, en.FmtNumber(10456.25, 0)))
	fmt.Println()
	fmt.Println(zh.R("range_test", 0, 0, 1, 0, zh.FmtNumber(0, 0), zh.FmtNumber(1, 0)))
	fmt.Println(zh.R("range_test", 1, 0, 2, 0, zh.FmtNumber(1, 0), zh.FmtNumber(2, 0)))
	fmt.Println(zh.R("range_test", 1, 0, 100, 0, zh.FmtNumber(1, 0), zh.FmtNumber(100, 0)))
	fmt.Println(en.R("range_test", 0, 0, 1, 0, en.FmtNumber(0, 0), en.FmtNumber(1, 0)))
	fmt.Println(en.R("range_test", 1, 0, 2, 0, en.FmtNumber(1, 0), en.FmtNumber(2, 0)))
	fmt.Println(en.R("range_test", 1, 0, 100, 0, en.FmtNumber(1, 0), en.FmtNumber(100, 0)))
	fmt.Println()
	fmt.Println(zh.FmtDateFull(time.Now()))
	fmt.Println(en.FmtDateFull(time.Now()))
}

// FmtNumber 格式化数字. num 为待格式化的数字, v 为保留位数.
// 不同语言的 translator 根据 num 和 digits 选择单/复数格式.
