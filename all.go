package slices

func All[T comparable](f PredicateFunc[T], elements ...T) (bool, error) {
	for _, element := range elements {
		isMatch, err := f(element)
		if err != nil {
			return false, err
		}
		if !isMatch {
			return false, nil
		}
	}

	return true, nil
}
