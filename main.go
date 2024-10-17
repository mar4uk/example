package main

import (
	"fmt"

	app "github.com/mar4uk/hello/pkg/go"
)

func main() {
	println("Ba dum, tss!")
	a := app.App{
		BundleID: "abc",
	}
	fmt.Println(a)
}
