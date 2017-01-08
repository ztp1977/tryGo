package http

import (
	"testing"
	"time"
)

//go test -bench . -benchmem
func BenchmarkConcat(b *testing.B) {
	var ss []string
	for n := 0; n < 100; n++ {
		ss = append(ss, "foo")
	}

	for i := 0; i < b.N; i++ {
		Concat(ss...)
	}
}

func BenchmarkConcat2(b *testing.B) {
	var ss []string
	for n := 0; n < 100; n++ {
		ss = append(ss, "foo")
	}

	for i := 0; i < b.N; i++ {
		Concat2(ss...)
	}
}

func BenchmarkDivision(b *testing.B) {
	b.StopTimer()
	// 初期化などの時間はここに含まないようにするため
	time.Sleep(10 * time.Nanosecond)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
