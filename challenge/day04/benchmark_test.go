package day04

import (
	"testing"

	"github.com/nlowe/aoc2025/challenge"
)

func BenchmarkA(b *testing.B) {
	for b.Loop() {
		_ = partA(challenge.InputFile())
	}
}

func BenchmarkB(b *testing.B) {
	for b.Loop() {
		_ = partB(challenge.InputFile())
	}
}
