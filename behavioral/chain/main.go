// @file: main.go
// @date: 2021/11/9

package main

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

type ChainNode interface {
	Execute(*Patient)
	SetNext(ChainNode)
}

func main() {
	cashier := &Cashier{}

	medical := &Medical{}
	medical.SetNext(cashier)

	doctor := &Doctor{}
	doctor.SetNext(medical)

	reception := &Reception{}
	reception.SetNext(doctor)

	patient := &Patient{name: "abc"}
	reception.Execute(patient)
}
