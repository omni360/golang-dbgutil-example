package main

import "github.com/realint/dbgutil"

func main() {
	var a = 10

	dbgutil.Display("a", a).Break(a == 10)
}

