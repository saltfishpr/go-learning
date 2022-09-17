// @file: observer.go
// @date: 2021/11/11

package main

import "fmt"

type Observer interface {
	Update()
}

type BinaryObserver struct {
	subject *Subject
}

func NewBinaryObserver(subject *Subject) *BinaryObserver {
	binaryObserver := &BinaryObserver{subject: subject}
	binaryObserver.subject.Attach(binaryObserver)
	return binaryObserver
}

func (o *BinaryObserver) Update() {
	fmt.Printf("Binary string: %b\n", o.subject.GetState())
}

type OctalObserver struct {
	subject *Subject
}

func NewOctalObserver(subject *Subject) *OctalObserver {
	octalObserver := &OctalObserver{subject: subject}
	octalObserver.subject.Attach(octalObserver)
	return octalObserver
}

func (o *OctalObserver) Update() {
	fmt.Printf("Octal string: %o\n", o.subject.GetState())
}

type HexaObserver struct {
	subject *Subject
}

func NewHexaObserver(subject *Subject) *HexaObserver {
	hexaObserver := &HexaObserver{subject: subject}
	hexaObserver.subject.Attach(hexaObserver)
	return hexaObserver
}

func (o *HexaObserver) Update() {
	fmt.Printf("Hex string: %X\n", o.subject.GetState())
}
