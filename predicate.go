package slices

type PredicateFunc[T comparable] func(object T) (bool, error)
