package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// --- BEGIN CLI FLAGS
	// Number
	n := flag.Int("number", 1, "The number of passwords to generate.")
	// Length
	l := flag.Int("length", 8, "The length of each password to generate.")
	// File
	f := flag.Bool("file", false, "A boolean indicating whether the passwords should print to a file in the current directory")
	// File Extension
	fe := flag.String("ext", "txt", "File extension to write. \n"+
		"USAGE: \n"+
		"-ext={txt|csv} \n")
	// Type
	t := flag.String("type", "alphanumeric", "The type of password to generate. Valid values include \"alphanumeric\"/\"n\" and \"numeric\"/\"n\".")
	// --- END CLI FLAGS
	flag.Parse()

	// handle file extensions
	var ext string
	switch *fe {
	case "txt", "csv":
		ext = "." + *fe
	default:
		fmt.Printf("%s is not a supported file extension.", *fe)
		os.Exit(2)
	}

	// TODO: Move this logic into a function
	switch {
	case *t == "a", *t == "alphanumeric":
		r := GenerateAlphanumericPasswords(n, l)
		if *f {
			dir := GetDir()
			WritePasswordsToFile(&dir, &ext, &r)
		} else {
			for i, v := range r {
				ct := i + 1
				fmt.Printf("Password #%d: %s\n", ct, v)
			}
		}
	case *t == "n", *t == "numeric":
		r := GenerateNumericPasswords(n, l)
		if *f {
			dir := GetDir()
			WritePasswordsToFile(&dir, &ext, &r)
		} else {
			for i, v := range r {
				ct := i + 1
				fmt.Printf("Password #%d: %s\n", ct, v)
			}
		}
	}
}
