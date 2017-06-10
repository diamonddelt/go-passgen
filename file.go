package main

import (
	"fmt"
	"log"
	"os"
)

// WritePasswordsToFile will write an incoming slice of bytes to a file on the local filesystem.
// fp (*string) = the file path to write the passwords to.
// fe (*string) = the file extension the file should be.
// dat ([]byte) = the raw data being written.
func WritePasswordsToFile(fp *string, fe *string, dat *[]string) {
	// TODO: append a datetime to this string
	n := "passgen-output" // the name of the file itself

	// TODO: use path/filepath for a better filepath
	fn := *fp + "/" + n + *fe // the fully qualified filepath with extension

	// create the file
	f, err := os.Create(fn)
	check(err)
	defer f.Close()

	// write the file
	for _, s := range *dat {
		res := s + "\n"
		_, err := f.WriteString(res)
		check(err)
		// fmt.Printf("Wrote %d bytes\n", i)
	}
	f.Sync()
	fmt.Printf("Password file generated at: %s\n", fn)
}

// check logs and breaks on all errors specific to file IO
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// GetDir returns the current working directory of program execution
func GetDir() string {
	d, err := os.Getwd()
	check(err)
	return d
}
