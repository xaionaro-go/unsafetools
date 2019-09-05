# Description

This package provides function `FieldByName` to access to any field (including private/unexported) of a structure.

# Use case

This package is supposed to be used for unit-tests only. If you think about using it in a real production then it seems it is something wrong in your program. However yes, you can use is if you want :)

[An use case example in github.com/xaionaro-go/picapi](https://github.com/xaionaro-go/picapi/blob/2ac776187b13158bca34bafe7cbff5487f478b9b/httpserver/http_server_handle_resize_test.go#L22).

# Example

`github.com/xaionaro-go/unsafetools/test/types.go`
```go
package test

type privateStruct struct {
	enableBonus bool
}

type StructWithPrivate struct {
	privateStruct

	initialized bool
}

func (s *StructWithPrivate) HelloWorld() (result string) {
	defer func() {
		if s.enableBonus {
			result += ` (bonus!)`
		}
	}()

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

	*FieldByName(s, `initialized`).(*bool) = true
	assert.Equal(t, `hello world!`, s.HelloWorld())

	*FieldByName(s, `initialized`).(*bool) = false
	assert.Equal(t, ``, s.HelloWorld())
}
```
