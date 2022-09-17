// @file: image.go
// @date: 2021/11/9

package main

import "fmt"

type Image interface {
	Display()
}

type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	realImage := &RealImage{filename: filename}
	realImage.loadFromDisk(filename)
	return realImage
}

func (i *RealImage) Display() {
	fmt.Println("Displaying ", i.filename)
}

func (RealImage) loadFromDisk(filename string) {
	fmt.Println("Loading ", filename)
}

type ProxyImage struct {
	realImage *RealImage
	filename  string
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename: filename}
}

func (i *ProxyImage) Display() {
	if i.realImage == nil {
		i.realImage = NewRealImage(i.filename)
	}
	i.realImage.Display()
}
