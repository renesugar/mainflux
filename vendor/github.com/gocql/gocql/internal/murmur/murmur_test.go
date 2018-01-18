package murmur

import (
	"strconv"
	"testing"
)

// Test the implementation of murmur3
func TestMurmur3H1(t *testing.T) {
	// these examples are based on adding a index number to a sample string in
	// a loop. The expected values were generated by the java datastax murmur3
	// implementation. The number of examples here of increasing lengths ensure
	// test coverage of all tail-length branches in the murmur3 algorithm
	seriesExpected := [...]uint64{
		0x0000000000000000, // ""
		0x2ac9debed546a380, // "0"
		0x649e4eaa7fc1708e, // "01"
		0xce68f60d7c353bdb, // "012"
		0x0f95757ce7f38254, // "0123"
		0x0f04e459497f3fc1, // "01234"
		0x88c0a92586be0a27, // "012345"
		0x13eb9fb82606f7a6, // "0123456"
		0x8236039b7387354d, // "01234567"
		0x4c1e87519fe738ba, // "012345678"
		0x3f9652ac3effeb24, // "0123456789"
		0x3f33760ded9006c6, // "01234567890"
		0xaed70a6631854cb1, // "012345678901"
		0x8a299a8f8e0e2da7, // "0123456789012"
		0x624b675c779249a6, // "01234567890123"
		0xa4b203bb1d90b9a3, // "012345678901234"
		0xa3293ad698ecb99a, // "0123456789012345"
		0xbc740023dbd50048, // "01234567890123456"
		0x3fe5ab9837d25cdd, // "012345678901234567"
		0x2d0338c1ca87d132, // "0123456789012345678"
	}
	sample := ""
	for i, expected := range seriesExpected {
		assertMurmur3H1(t, []byte(sample), expected)

		sample = sample + strconv.Itoa(i%10)
	}

	// Here are some test examples from other driver implementations
	assertMurmur3H1(t, []byte("hello"), 0xcbd8a7b341bd9b02)
	assertMurmur3H1(t, []byte("hello, world"), 0x342fac623a5ebc8e)
	assertMurmur3H1(t, []byte("19 Jan 2038 at 3:14:07 AM"), 0xb89e5988b737affc)
	assertMurmur3H1(t, []byte("The quick brown fox jumps over the lazy dog."), 0xcd99481f9ee902c9)
}

// helper function for testing the murmur3 implementation
func assertMurmur3H1(t *testing.T, data []byte, expected uint64) {
	actual := Murmur3H1(data)
	if actual != expected {
		t.Errorf("Expected h1 = %x for data = %x, but was %x", expected, data, actual)
	}
}

// Benchmark of the performance of the murmur3 implementation
func BenchmarkMurmur3H1(b *testing.B) {
	data := make([]byte, 1024)
	for i := 0; i < 1024; i++ {
		data[i] = byte(i)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h1 := Murmur3H1(data)
			if h1 != uint64(7627370222079200297) {
				b.Fatalf("expected %d got %d", uint64(7627370222079200297), h1)
			}
		}
	})
}