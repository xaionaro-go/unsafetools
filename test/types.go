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
