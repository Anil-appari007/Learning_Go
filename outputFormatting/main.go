package main

import "fmt"

func main() {
	a := "hello"
	b := "world"
	fmt.Print(a, b)
	fmt.Print("\n", a, " ", b)
	fmt.Print("\n", a, "\t", b)
	fmt.Print("\n", a, "\t\t", b)

	c := 23
	d := 900

	fmt.Print("\n", c, d)

	e := 98.09
	f := 9.34
	fmt.Print("\n", e, f)

	g := true
	h := false
	fmt.Print("\n", g, h)

	fmt.Printf("\nvalue of a is %s", a)
	fmt.Printf("\n value of a is %v", a)
	fmt.Printf("\n A is %T", a)

	fmt.Printf("\n value of c is %s", c)
	fmt.Printf("\n value of c is %v", c)
	fmt.Printf("\n C is %T", c)

	fmt.Printf("\n value of e is %s", e)
	fmt.Printf("\n value of e is %v", e)

	fmt.Printf("\ne is %T", e)

	fmt.Printf("\n g is %s", g)
	fmt.Printf("\n g is %v", g)
	fmt.Printf("\n g is %T", g)

	fmt.Printf("\n g is %#v", g)
	fmt.Printf("\n a is %#v", a)
	fmt.Printf("\n c is %#v", c)
	fmt.Printf("\n e is %#v", e)
}
