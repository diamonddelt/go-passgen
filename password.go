package main

import (
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
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

// generateAlphanumericPassword generates a single cryptographically secure, base64-encoded alphanumeric password
// l (*int) = the length of the password to generate
// Returns a password string (string)
func generateAlphanumericPassword(l *int) string {
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

// generateNumericPassword generates a single non-cryptographically secure numeric password
// l (*int) = the length of the password to generate
// Returns a password string (string)
func generateNumericPassword(l *int) string {
	// Since numeric only passwords can be cracked very easily unless they are extremely long,
	// this method foregoes the crypto/rand package for math/rand for ease of use
	var r string
	for i := 0; i < *l; i++ {
		r += strconv.Itoa(mrand.Intn(9))
	}
	return r
}

// GeneratePasswords generates multiple random passwords based on the type passed in
// t (*string) = the type of password to generate
// n (*int) = the number of passwords to generate
// l (*int) = the length of each password
// Returns a slice of generated passwords and an error if encountered ([]string, error)
func GeneratePasswords(t *string, n *int, l *int) ([]string, error) {
	var r []string
	var p string
	var err error
	mrand.Seed(time.Now().UTC().UnixNano()) // this needs to move anywhere it will only be called once; moved after refactor 6/18/17

// GenerateNumericPasswords generates multiple random numeric passwords
// n (*int) = a pointer to a value specifying the number of passwords to generate (int)
// l (*int) = a pointer to a value specifying the length of each password (int)
// Returns an array of generated passwords ([]string)
func GenerateNumericPasswords(n *int, l *int) []string {
	mrand.Seed(time.Now().UTC().UnixNano()) // consider moving this to main goroutine for highly concurrent usage
	var r []string
	for i := 0; i < *n; i++ {
		switch {
		case *t == "a", *t == "alphanumeric":
			p = generateAlphanumericPassword(l)
		case *t == "n", *t == "numeric":
			p = generateNumericPassword(l)
		default:
			err = fmt.Errorf(`"%s" is not a supported flag value for "type". Supported values are "alphanumeric"/"a" and "numeric"/"n".\n`, *t)
			break
		}
		r = append(r, p)
	}
	return r, err
}
