package slices

func Any[T comparable](predicate predicateFunc[T], elements ...T) (bool, error) {
	for _, element := range elements {
		isMatch, err := predicate(element)
		if err != nil {
			return false, err
		}
		if isMatch {
			return true, nil
		}
	}

	return false, nil
}
