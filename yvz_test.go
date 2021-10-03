package main

import (
	"testing"
)

func Benchmark(b *testing.B) {
	Walkdir("/home/yvz/dev/go_projects/koerber")
}
