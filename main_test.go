package main

import (
	"testing"
)

// Used for performance testing
func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
