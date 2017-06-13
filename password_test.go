package main

import (
	"fmt"
	"testing"
)

func TestGeneratePasswords(t *testing.T) {

	n, l := 10, 20
	fmt.Printf("Testing generating %d passwords with a length of %d...\n", n, l)
	passwords := GeneratePasswords(&n, &l)
	for _, v := range passwords {
		fmt.Println(v)
		if len(v) != 20 {
			t.Errorf("The password generated is the incorrect length, got: %d, want: %d.", len(v), 20)
		}
	}
}
