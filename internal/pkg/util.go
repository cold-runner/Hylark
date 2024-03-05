package pkg

func Convert[T string | int | bool](v T) *T {
	return &v
}
