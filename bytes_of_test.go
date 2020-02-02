package unsafetools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/xaionaro-go/unsafetools"
)

func TestBytesOf(t *testing.T) {
	s := &someTestStruct{1, 2}
	b := BytesOf(s)
	assert.Equal(t, uint8(1), b[0])
	assert.Equal(t, uint8(2), b[1])
	assert.Equal(t, 2, len(b))
}
