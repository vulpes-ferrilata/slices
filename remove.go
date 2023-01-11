package slices

func Remove[T comparable](slice []T, elements ...T) []T {
	newSlice := make([]T, 0)

	for _, sliceElement := range slice {
		isExists := false
		for _, element := range elements {
			if sliceElement == element {
				isExists = true
				break
			}
		}

		if !isExists {
			newSlice = append(newSlice, sliceElement)
		}
	}

	return newSlice
}
