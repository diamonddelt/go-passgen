package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	start := time.Now() // crude execution time benchmarking

	// --- BEGIN CLI FLAGS
	n := flag.Int("number", 1, "the number of passwords to generate")     // returns an &int
	l := flag.Int("length", 8, "the length of each password to generate") // returns an &int
	f := flag.Bool("file", false, "boolean indicating whether the passwords should print to a file in the current directory")
	// --- END CLI FLAGS
	flag.Parse()

	r := GeneratePasswords(n, l) // get []string of passwords
	if *f {
		dir := GetDir()
		// TODO: Support more than just .txt file extensions
		ext := ".txt"
		WritePasswordsToFile(&dir, &ext, &r)
	} else {
		for i, v := range r {
			ct := i + 1
			fmt.Printf("Password #%d: %s\n", ct, v)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s", elapsed) // crude execution time benchmarking

	// use Go's reflect package to inspect variable type at runtime ; NOTE: reflection is SLOW
	// --- DEBUG FLAGS ---
	// fmt.Println(reflect.TypeOf(lengthFlagPtr))
	// fmt.Printf("Length: %d", *lengthFlagPtr)
	// --- END DEBUG FLAGS ---
}
