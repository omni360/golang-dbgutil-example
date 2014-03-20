package main

import "github.com/realint/dbgutil"

type mytype struct {
	next *mytype
	prev *mytype
}

func main() {
	var v1 = new(mytype)
	var v2 = new(mytype)
	var v3 = new(mytype)

	v1.prev = v3
	v1.next = v2

	v2.prev = v1
	v2.next = v3

	v3.prev = v2
	v3.next = nil

	dbgutil.FormatDisplay("v1", v1, "v2", v2, "v3", v3).Break(true)
}

