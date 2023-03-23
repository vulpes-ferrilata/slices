package slices

func Any[T comparable](f PredicateFunc[T], elements ...T) (bool, error) {
	for _, element := range elements {
		isMatch, err := f(element)
		if err != nil {
			return false, err
		}
		if isMatch {
			return true, nil
		}
	}

	return false, nil
}
