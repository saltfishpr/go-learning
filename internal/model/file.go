// @file: file.go
// @date: 2021/10/29

package model

type File struct {
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Content string `json:"content"`
}
