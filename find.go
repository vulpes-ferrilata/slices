package slices

func Find[T comparable](f PredicateFunc[T], elements ...T) (T, error) {
	var empty T

	for _, element := range elements {
		isMatch, err := f(element)
		if err != nil {
			return empty, err
		}
		if isMatch {
			return element, nil
		}
	}

	return empty, ErrNoMatchFound
}
