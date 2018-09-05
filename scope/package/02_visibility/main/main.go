package main

import (
	"fmt"

	"github.com/oleg-koval/goTraining/scope/package/02_visibility/vis"
)

func main() {
	fmt.Println("call main.go")
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
