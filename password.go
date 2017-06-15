package main

import (
	crand "crypto/rand"
	"encoding/base64"
	mrand "math/rand"
	"strconv"
	"time"
)

// Password which adheres to NIST recommendations
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
	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}

	// encode the byte slice to base64 for general usage
	// base64 is safe for URI headers, HTTP forms, JSON, and plaintext usage
	p := base64.URLEncoding.EncodeToString(b)

	// remove extra characters from encoding to match user length
	p = p[:*l]
	return p
}

// generateNumericPassword generates a single non-cryptographically secure numeric password.
// Uses the current datetime as a seed value for the RNG.
// l (*int) = a pointer to an integer representing the length of the password to generate.
// Returns the random number as an integer.
func generateNumericPassword(l *int) int {
	// Since numeric only passwords can be cracked very easily unless they are extremely long,
	// this method foregoes the crypto/rand package for math/rand for ease of use
	mrand.Seed(time.Now().UTC().UnixNano()) // consider moving this to main goroutine for highly concurrent usage

	var r int
	var tmps string
	for i := 0; i < *l; i++ {
		tmps += strconv.Itoa(mrand.Intn(9))
	}
	r, _ = strconv.Atoi(tmps) // don't care about the error right now
	return r
}

// GeneratePasswords generates multiple random passwords
// n (*int) = a pointer to a value specifying the number of passwords to generate (int)
// l (*int) = a pointer to a value specifying the length of each password (int)
// Returns an array of generated passwords ([]string)
func GeneratePasswords(n *int, l *int) []string {
	var r []string
	for i := 0; i < *n; i++ {
		p := generatePassword(l)
		r = append(r, p)
	}
	return r
}
