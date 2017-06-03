package main

import (
	"flag"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Enter the length of the password:")
	lengthFlagPtr := flag.Int("length", 8, "a number representing the length of the password to generate") // returns an &int

	flag.Parse() // parse commandline args

	// use Go's reflect package to inspect variable type at runtime ; NOTE: reflection is SLOW
	fmt.Println(reflect.TypeOf(lengthFlagPtr))
	fmt.Printf("Length: %d", *lengthFlagPtr)
}

// Password which adheres to the specifications in the README
type Password struct {
	length     int    // could also be int16, int32, or int64
	complexity string // how complex is it? according to NIST

	// TODO: research https://golang.org/pkg/flag/
}
