package main

import (
	"scope/packageone"
)

var myVar = "This is the package level variable" // package level variable

func main() {
	var blockVar = "This is the block level variable"

	packageone.PrintMe(myVar, blockVar)

}
