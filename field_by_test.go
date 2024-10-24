package unsafetools

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xaionaro-go/unsafetools/test"
)

func TestFindByName_positive(t *testing.T) {
	s := &test.StructWithPrivate{}

	assert.Equal(t, ``, s.HelloWorld())

	*FieldByName[bool](s, `initialized`) = true
	assert.Equal(t, `hello world!`, s.HelloWorld())

	*FieldByName[bool](s, `initialized`) = false
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
		FieldByName[struct{}](s, `wrongField`)
	})
	countIfPanic(func() {
		_ = FieldByName[struct{}](s, `initialized`) /* wrong type */
	})

	assert.Equal(t, 2, panicCount)
}

func TestFindByPath_positive(t *testing.T) {
	s := &test.StructWithPrivate{}

	assert.Equal(t, ``, s.HelloWorld())

	*FieldByPath[bool](s, []string{`initialized`}) = true
	assert.Equal(t, `hello world!`, s.HelloWorld())

	*FieldByPath[bool](s, []string{`privateStruct`, `enableBonus`}) = true
	assert.Equal(t, `hello world! (bonus!)`, s.HelloWorld())
}

func TestFindByPath_negative(t *testing.T) {
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
		FieldByPath[struct{}](s, []string{ /* empty path */ })
	})
	countIfPanic(func() {
		FieldByPath[struct{}](s, []string{`wrongField`})
	})
	countIfPanic(func() {
		FieldByPath[struct{}](s, []string{`privateStruct`, `wrongField`})
	})
	countIfPanic(func() {
		FieldByPath[struct{}](s, []string{`privateStruct`, `enableBonus`, `extraField`})
	})
	countIfPanic(func() {
		_ = FieldByPath[struct{}](s, []string{`initialized`}) /* wrong type */
	})

	assert.Equal(t, 5, panicCount)
}
