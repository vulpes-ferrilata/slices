package slices

func Filter[T comparable](predicate predicateFunc[T], elements ...T) ([]T, error) {
	newSlice := make([]T, 0)

	for _, element := range elements {
		isMatch, err := predicate(element)
		if err != nil {
			return newSlice, err
		}
		if isMatch {
			newSlice = append(newSlice, element)
		}
	}

	return newSlice, nil
}
