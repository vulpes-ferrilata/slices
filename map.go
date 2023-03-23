package slices

func Map[T comparable, R comparable](f MapperFunc[T, R], elements ...T) ([]R, error) {
	results := make([]R, 0)

	for _, element := range elements {
		result, err := f(element)
		if err != nil {
			return results, err
		}

		results = append(results, result)
	}

	return results, nil
}
