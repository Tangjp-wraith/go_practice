package main

import (
	"strings"
	"testing"
)

func strPlus1(a []string) string {
	var s, sep string
	for i := 0; i < len(a); i++ {
		s += sep + a[i]
		sep = " "
	}
	return s
}

func strPlus2(a []string) string {
	return strings.Join(a, " ")
}

func BenchmarkStrPlus1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strPlus1([]string{"BGM", "is", "learning", "Golang"})
	}
}

func BenchmarkStrPlus2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strPlus2([]string{"BGM", "is", "learning", "Golang"})
	}
}
