package main

import "github.com/realint/dbgutil"

type mytype struct {
	next *mytype
	prev *mytype
}

func main() {
	var v1 = new(mytype)
	var v2 = new(mytype)

	v1.prev = nil
	v1.next = v2

	v2.prev = v1
	v2.next = nil

	dbgutil.Display("v1", v1, "v2", v2).Break(true)
}

