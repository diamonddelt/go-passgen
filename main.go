package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now() // crude execution time benchmarking

	// d := getDir()
	// fmt.Println(d) // display the current directory ; needed later for printing to file

	// --- BEGIN CLI FLAGS ---
	n := flag.Int("number", 1, "the number of passwords to generate")     // returns an &int
	l := flag.Int("length", 8, "the length of each password to generate") // returns an &int
	// --- END CLI FLAGS ---

	flag.Parse() // parse commandline args - MUST BE CALLED AFTER ALL FLAGS ARE INITIALIZED

	// generate desired number of passwords
	r := generatePasswords(n, l)
	for i, v := range r {
		ct := i + 1 // user friendly counter
		fmt.Printf("Password #%d: %s\n", ct, v)
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s", elapsed) // crude execution time benchmarking

	// use Go's reflect package to inspect variable type at runtime ; NOTE: reflection is SLOW
	// --- DEBUG FLAGS ---
	// fmt.Println(reflect.TypeOf(lengthFlagPtr))
	// fmt.Printf("Length: %d", *lengthFlagPtr)
	// --- END DEBUG FLAGS ---
}

// Password which adheres to NIST recommendations
// Use this later
type Password struct {
	id             string // the actual password string
	length         int    // could also be int16, int32, or int64
	complexity     string // how complex is it? according to NIST
	isNumeric      bool
	isAlphanumeric bool
	isAlphabetical bool
}

// Must be a 'CSPRNG' - "Cryptographically Secure Pseudo-Random Number Generator"
// Requires crypto/rand and encoding/base64 packages
func generatePassword(l *int) string {
	// generate cryptographically secure byte slice
	b := make([]byte, *l)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	// encode the byte slice to base64 for general usage
	// base64 is safe for URI headers, HTTP forms, JSON, and plaintext usage
	p := base64.URLEncoding.EncodeToString(b)
	return p
}

func generatePasswords(n *int, l *int) []string {
	var r []string
	for i := 0; i < *n; i++ {
		p := generatePassword(l)
		r = append(r, p)
	}
	return r
}

// Get working directory of the application
func getDir() string {
	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return d
}
