package main

import "testing"

// Using Go 1.7+ ability to express multiple benchmarks in table-driven format
func BenchmarkGenerateNumericPasswords(b *testing.B) {
	benchmarks := []struct {
		name   string
		number int
		length int
	}{
		{`Number: 1 | Length: 10`, 1, 10},
		{`Number: 300  Length: 300`, 300, 300},
	}

	for _, v := range benchmarks {
		b.Run(v.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GenerateNumericPasswords(&v.number, &v.length)
			}
		})
	}

}

func BenchmarkGenerateAlphanumericPasswords(b *testing.B) {
	benchmarks := []struct {
		name   string
		number int
		length int
	}{
		{`Number: 1 | Length: 10`, 1, 10},
		{`Number: 300  Length: 300`, 300, 300},
	}

	for _, v := range benchmarks {
		b.Run(v.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GenerateAlphanumericPasswords(&v.number, &v.length)
			}
		})
	}

}

// TODO: Go 1.7+ also allows sub-tests and sub-benchmarks
