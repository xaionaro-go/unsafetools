# Description

This package provides function `FieldByName` to access to any field (including private/unexported) of a structure.

# Example

`github.com/xaionaro-go/unsafetools/test/types.go`
```go
package test

type StructWithPrivate struct {
	initialized bool
}

func (s *StructWithPrivate) HelloWorld() string {
	if !s.initialized {
		return ``
	}

	return `hello world!`
}
```

`github.com/xaionaro-go/unsafetools/unsafetools_test.go`
```go
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
```
