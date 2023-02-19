package slices

func All[T comparable](predicate predicateFunc[T], elements ...T) (bool, error) {
	for _, element := range elements {
		isMatch, err := predicate(element)
		if err != nil {
			return false, err
		}
		if !isMatch {
			return false, nil
		}
	}

	return true, nil
}
