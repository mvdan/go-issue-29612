package main

import (
	ssa1 "test.tld/foo/p1/ssa"
	ssa2 "test.tld/foo/p2/ssa"
)

func main() {
	// Remove the ssa1 import and these two lines to magically fix the program.
	v1 := &ssa1.T{}
	_ = v1

	v2 := &ssa2.T{}
	ssa2.Works(v2)
	println("Works succeeded")
	ssa2.Panics(v2)
	println("Panics succeeded")
}
