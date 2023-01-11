package slices

type mapperFunc[T comparable, R comparable] func(t T) (R, error)
