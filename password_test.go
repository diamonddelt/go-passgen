package main

import (
	"fmt"
	"log"
	"testing"
)

func TestGenerateAlphanumericPasswords(t *testing.T) {
	n, l := 10, 20
	passwords := GenerateAlphanumericPasswords(&n, &l)

	log.Printf("Verifying %d passwords were generated...", n)
	if len(passwords) != n {
		t.Errorf("The number of passwords generated is incorrect, got: %d, want: %d.", len(passwords), n)
	}

	log.Printf("Verifying %d passwords generated have a length of %d...", n, l)
	for _, v := range passwords {
		if len(v) != 20 {
			t.Errorf("The passwords generated are the incorrect length, got: %d, want: %d.", len(v), l)
			break
		}
	}
}

func TestGenerateNumericPasswords(t *testing.T) {
	n, l := 10, 20
	passwords := GenerateNumericPasswords(&n, &l)

	log.Printf("Verifying %d passwords were generated...", n)
	if len(passwords) != n {
		t.Errorf("The number of passwords generated is incorrect, got: %d, want: %d.", len(passwords), n)
	}

	log.Printf("Verifying %d passwords generated have a length of %d...", n, l)
	for _, v := range passwords {
		fmt.Println(v)
		if len(v) != 20 {
			t.Errorf("The password generated is the incorrect length, got: %d, want: %d.", len(v), l)
			break
		}
	}
}
