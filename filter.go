package slices

func Filter[T comparable](f PredicateFunc[T], elements ...T) ([]T, error) {
	newSlice := make([]T, 0)

	for _, element := range elements {
		isMatch, err := f(element)
		if err != nil {
			return newSlice, err
		}
		if isMatch {
			newSlice = append(newSlice, element)
		}
	}

	return newSlice, nil
}
