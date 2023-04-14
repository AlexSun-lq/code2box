package code2box

import "testing"

// BenchmarkGetDigitsNumv2_1234-12    	134725719	         8.894 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	demo/code2box	2.873s
func BenchmarkGetDigitsNumv2_1234(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getDigitsNumv2(1234)
	}
}

// BenchmarkGetDigitsNumv2_1234567890-12    	128504234	         8.955 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	demo/code2box	2.231s
func BenchmarkGetDigitsNumv2_1234567890(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getDigitsNumv2(1234)
	}
}

// BenchmarkGetDigitsNum_1234-12    	369869889	         3.436 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	demo/code2box	1.712s
func BenchmarkGetDigitsNum_1234(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDigitsNum(1234)
	}
}

// BenchmarkGetDigitsNum_1234567890-12    	349662285	         3.453 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	demo/code2box	1.675s
func BenchmarkGetDigitsNum_1234567890(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDigitsNum(1234)
	}
}
