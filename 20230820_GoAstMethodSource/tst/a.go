package tst

import "methodSourceExample/tst2"

type A struct {
	Age  int
	Name string
}

func (a A) Hello() string {
	return "Hello" + "Damian"
}

func NewA(name string) A {
	tst2.Crap()
	return A{
		Name: name,
		Age:  0,
	}
}
