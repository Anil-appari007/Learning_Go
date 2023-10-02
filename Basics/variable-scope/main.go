package main

import "fmt"

var o_var string = "outside variable"

var o_var2 string = ""

func main() {
	fmt.Println(o_var)
	for each, value := range o_var {
		fmt.Printf("%v->%v->%s->%c\n", each, value, string(value), value)
		o_var2 += string(value)
	}

	// fmt.Printf(each)	## output: Basics\variable-scope\main.go:13:13: undefined: each
	fmt.Println("value of o_var is added to o_var2 inside loop and is accessible outside loop ", o_var2)
}
