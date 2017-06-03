package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
)

func main() {
	d := getDir()
	fmt.Println(d) // display the current directory ; needed later for printing to file

	lengthFlagPtr := flag.Int("length", 8, "a number representing the length of the password to generate") // returns an &int

	flag.Parse() // parse commandline args

	// use Go's reflect package to inspect variable type at runtime ; NOTE: reflection is SLOW
	fmt.Println(reflect.TypeOf(lengthFlagPtr))
	fmt.Printf("Length: %d", *lengthFlagPtr)
}

// Password which adheres to NIST recommendations
type Password struct {
	id             string // the actual password string
	length         int    // could also be int16, int32, or int64
	complexity     string // how complex is it? according to NIST
	isNumeric      bool
	isAlphanumeric bool
	isAlphabetical bool
	// TODO: research https://golang.org/pkg/flag/
}

// Unexported method to derive password based on the Password struct it receives
func (p *Password) generatePassword() (string, bool) {
	// TODO: logic to generate password
	// use golang crypto/rand library instead of default rand
	// research a 'byte slice' for this

	return "", true
}

// Validate user is in a useable directory
func getDir() string {
	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return d
}
