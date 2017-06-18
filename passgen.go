package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// --- BEGIN CLI FLAGS
	n := flag.Int("number", 1, "The number of passwords to generate.")
	l := flag.Int("length", 8, "The length of each password to generate.")
	f := flag.Bool("file", false, "A boolean indicating whether the passwords should print to a file in the current directory (Currently only supports .txt file extensions).")
	t := flag.String("type", "alphanumeric", "The type of password to generate. Valid values include \"alphanumeric\"/\"n\" and \"numeric\"/\"n\".")
	// --- END CLI FLAGS
	flag.Parse()

	r, err := GeneratePasswords(t, n, l)
	if err != nil {
		log.Fatal(err)
	}

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
}
