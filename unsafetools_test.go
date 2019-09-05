package unsafetools

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xaionaro-go/unsafetools/test"
)

func TestFindByName_positive(t *testing.T) {
	s := &test.StructWithPrivate{}

	assert.Equal(t, ``, s.HelloWorld())

	*FieldByName(s, `initialized`).(*bool) = true
	assert.Equal(t, `hello world!`, s.HelloWorld())

	*FieldByName(s, `initialized`).(*bool) = false
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
		FieldByName(s, `wrongField`)
	})
	countIfPanic(func() {
		FieldByName(*s /* not pointer */, `initialized`)
	})
	countIfPanic(func() {
		_ = FieldByName(s, `initialized`).(bool) /* wrong type */
	})

	assert.Equal(t, 3, panicCount)
}

func TestFindByPath_positive(t *testing.T) {
	s := &test.StructWithPrivate{}

	assert.Equal(t, ``, s.HelloWorld())

	*FieldByPath(s, []string{`initialized`}).(*bool) = true
	assert.Equal(t, `hello world!`, s.HelloWorld())

	*FieldByPath(s, []string{`privateStruct`, `enableBonus`}).(*bool) = true
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
		FieldByPath(s, []string{ /* empty path */ })
	})
	countIfPanic(func() {
		FieldByPath(s, []string{`wrongField`})
	})
	countIfPanic(func() {
		FieldByPath(s, []string{`privateStruct`, `wrongField`})
	})
	countIfPanic(func() {
		FieldByPath(s, []string{`privateStruct`, `enableBonus`, `extraField`})
	})
	countIfPanic(func() {
		FieldByPath(*s /* not pointer */, []string{`initialized`})
	})
	countIfPanic(func() {
		_ = FieldByPath(s, []string{`initialized`}).(bool) /* wrong type */
	})

	assert.Equal(t, 6, panicCount)
}
