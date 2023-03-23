package slices

type MapperFunc[T comparable, R comparable] func(t T) (R, error)
