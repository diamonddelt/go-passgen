package main

import (
	"crypto/rand"
	"encoding/base64"
)

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

// generatePassword generates a single cryptographically secure password.
// l (*int) = a pointer to an integer representing the length of the password to generate.
// Returns the base64-encoded password as a string.
func generatePassword(l *int) string {
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

// GeneratePasswords generates multiple random passwords
// n = a pointer to a value specifying the number of passwords to generate (int)
// l = a pointer to a value specifying the length of each password (int)
// Returns an array of generated passwords ([]string)
func GeneratePasswords(n *int, l *int) []string {
	var r []string
	for i := 0; i < *n; i++ {
		p := generatePassword(l)
		r = append(r, p)
	}
	return r
}
