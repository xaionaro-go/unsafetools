package unsafetools

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xaionaro-go/unsafetools/test"
)

func TestFindByName_positive(t *testing.T) {
	s := &test.StructWithPrivate{}

	assert.Equal(t, ``, s.HelloWorld())

	*(*bool)(FieldByName(s, `initialized`, (*bool)(nil))) = true
	assert.Equal(t, `hello world!`, s.HelloWorld())

	*(*bool)(FieldByName(s, `initialized`, nil)) = false
	assert.Equal(t, ``, s.HelloWorld())
}

func TestFindByName_negative(t *testing.T) {
	s := &test.StructWithPrivate{}

	panicCount := 0
	countIfPanic := func(fn func()) {
		defer func() {
			r := recover()
			if r == nil {
				return
			}
			panicCount++
		}()

		fn()
	}

	countIfPanic(func() {
		FieldByName(s, `wrongField`, (*bool)(nil))
	})
	countIfPanic(func() {
		FieldByName(*s /* not pointer */, `initialized`, (*bool)(nil))
	})
	countIfPanic(func() {
		FieldByName(s, `initialized`, true /* not pointer */)
	})
	countIfPanic(func() {
		FieldByName(s, `initialized`, (**bool)(nil) /* wrong sample */)
	})

	assert.Equal(t, 4, panicCount)
}
