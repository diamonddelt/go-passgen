package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// WritePasswordsToFile will write an incoming slice of bytes to a file on the local filesystem.
// fp (*string) = the file path to write the passwords to.
// fe (*string) = the file extension the file should be.
// dat (*[]string) = the raw data being written.
func WritePasswordsToFile(fp *string, fe *string, dat *[]string) {
	// TODO: append a datetime to this string
	n := "passgen-output" // the name of the file itself

	// TODO: use path/filepath for a better filepath
	fn := *fp + "/" + n + *fe // the fully qualified filepath with extension

	// create the file
	f, err := os.Create(fn)
	check(err)
	defer f.Close()

	if *fe == `.csv` {
		writeToCSV(f, dat)
	} else {
		// write the file
		for _, s := range *dat {
			res := s + "\n"
			_, err := f.WriteString(res)
			check(err)
			// fmt.Printf("Wrote %d bytes\n", i)
		}
	}
	f.Sync()
	fmt.Printf("Password file generated at: %s\n", fn)
}

// writeToCSV takes a pointer to a slice of strings, and writes each string iteratively to a csv file object
// f (*os.File) = the file object which is being written to
// dat (*[]string) = the data to write to the file object
func writeToCSV(f *os.File, dat *[]string) {
	var err error
	writer := csv.NewWriter(f) // create the csv writer
	defer writer.Flush()       // defer flushing the write buffer until the function returns

	for _, v := range *dat {
		var row []string
		row = append(row, v)
		err = writer.Write(row)
		check(err)
	}
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
