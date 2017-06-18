package main

import (
	"log"
	"testing"
)

func TestGeneratePasswords(t *testing.T) {
	// Test Cases:
	// Generated passwords are the correct length
	// The correct number of passwords are generated
	// A bad type input breaks the generation of passwords
	// TODO: The type of password generated (alpha/numeric) matches the input
	// TODO: Each consecutive password is different from the previous if multiple are specified
	// TODO: A file is written to the file system if specified

	n, l := 10, 20
	var passwords []string
	var err error
	ta1, ta2 := "a", "alphanumeric"
	tn1, tn2 := "n", "numeric"
	tb1 := "3x_38d??.31@" // bad input type

	// Generated passwords are the correct length
	passwords, err = GeneratePasswords(&ta1, &n, &l)
	log.Printf("Verifying %d passwords generated have a length of %d...", n, l)
	for _, v := range passwords {
		if len(v) != 20 {
			t.Errorf("The passwords generated are the incorrect length, got: %d, want: %d.", len(v), l)
			break
		}
	}

	// The correct number of passwords are generated
	log.Printf("Verifying %d passwords were generated...", n)
	if len(passwords) != n {
		t.Errorf("The number of passwords generated is incorrect, got: %d, want: %d.", len(passwords), n)
	}

	// A bad type input breaks generation of passwords
	log.Printf("Verifying bad input type breaks password generation...")
	_, err = GeneratePasswords(&tb1, &n, &l)
	if err == nil {
		t.Errorf("A bad type input was not handled, got: %s, want %s/%s/%s/%s.", tb1, ta1, ta2, tn1, tn2)
	}
}
