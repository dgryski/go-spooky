package spooky

import "testing"

func BenchmarkShort(b *testing.B) {

	buf := make([]byte, 32)

	s1 := uint64(0)
	s2 := uint64(0)

	for i := 0; i < b.N; i++ {
		Short(buf, &s1, &s2)
	}
}

func BenchmarkHash64(b *testing.B) {

	buf := make([]byte, 2048)

	for i := 0; i < b.N; i++ {
		Hash64(buf, 0)
	}
}
