package vis

import "fmt"

// PrintVar func to print some cool vars
func PrintVar() {
	fmt.Println("call printer.go")
	fmt.Println(MyName)
	fmt.Println(willNotExport)
}
