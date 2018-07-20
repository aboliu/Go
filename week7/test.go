package main

import (
	"phone"
)

type phoneA struct {
	owner         string
	contactPerson string
	black         string
}

func (a phoneA) Owner() string {
	return a.owner
}

func (a phoneA) ContactPerson() string {
	return a.contactPerson
}

func (a phoneA) BlackList() string {
	return a.black
}

type phoneB struct {
	owner         string
	contactPerson string
	black         string
}

func (b phoneB) Owner() string {
	return b.owner
}

func (b phoneB) ContactPerson() string {
	return b.contactPerson
}

func (b phoneB) BlackList() string {
	return b.black
}

func main() {
	a := phoneA{
		owner:         "PhoneA",
		contactPerson: "聯絡人:PhoneB",
		black:         "黑名單:PhoneB",
	}

	b := phoneB{
		owner:         "PhoneB",
		contactPerson: "聯絡人:PhoneA",
		black:         "黑名單:PhoneA",
	}

	phone.CallOut(a)
	phone.CallIn(b)

	phone.CallOut(b)
	phone.CallIn(a)
}
