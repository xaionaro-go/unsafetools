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
