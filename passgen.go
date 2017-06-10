package main

import (
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

	r := GeneratePasswords(n, l)
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

// Get working directory of the application
func getDir() string {
	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return d
}
