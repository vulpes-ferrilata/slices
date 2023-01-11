package slices

func Map[T comparable, R comparable](mapper mapperFunc[T, R], elements ...T) ([]R, error) {
	results := make([]R, 0)

	for _, element := range elements {
		result, err := mapper(element)
		if err != nil {
			return results, err
		}

		results = append(results, result)
	}

	return results, nil
}
