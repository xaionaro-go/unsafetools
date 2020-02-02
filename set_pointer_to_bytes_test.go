package unsafetools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/xaionaro-go/unsafetools"
)

func TestSetPointerToBytesPositive(t *testing.T) {
	b := []byte{1, 2}
	var s *someTestStruct
	assert.NoError(t, SetPointerToBytes(&s, b))
	assert.Equal(t, uint8(1), s.a)
	assert.Equal(t, uint8(2), s.b)
}

func TestSetPointerToBytesNegative(t *testing.T) {
	b := []byte{1}
	var s *someTestStruct
	assert.Error(t, SetPointerToBytes(&s, b))
}

// BenchmarkSetPointerToBytes-8   	64212001	        15.7 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSetPointerToBytes(b *testing.B) {
	bytes := []byte{1, 2}
	var s *someTestStruct

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SetPointerToBytes(&s, bytes)
	}
}
