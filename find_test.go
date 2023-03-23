package slices_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vulpes-ferrilata/slices"
)

var _ = Describe("Find", func() {
	Err := errors.New("error")

	DescribeTable("when slice of integer",
		func(slice []int, f slices.PredicateFunc[int], expectedResult int, expectedErr error) {
			result, err := slices.Find(f, slice...)
			if expectedErr == nil {
				Expect(err).ShouldNot(HaveOccurred())
			} else {
				Expect(err).Should(MatchError(err))
			}
			Expect(result).Should(BeEquivalentTo(expectedResult))

		},
		Entry(
			"when predicate function return true",
			[]int{2, 3, 4, 5, 6},
			func(object int) (bool, error) {
				return true, nil
			},
			2,
			nil,
		),
		Entry(
			"when predicate function return false",
			[]int{2, 3, 4, 5, 6},
			func(object int) (bool, error) {
				return false, nil
			},
			0,
			slices.ErrNoMatchFound,
		),
		Entry(
			"when predicate function match value in slice",
			[]int{2, 3, 4, 5, 6},
			func(object int) (bool, error) {
				return object == 3, nil
			},
			3,
			nil,
		),
		Entry(
			"when predicate function match some values in slice",
			[]int{2, 3, 4, 5, 6},
			func(object int) (bool, error) {
				return object > 3, nil
			},
			4,
			nil,
		),
		Entry(
			"when predicate function return error",
			[]int{2, 3, 4, 5, 6},
			func(object int) (bool, error) {
				return true, Err
			},
			0,
			Err,
		),
	)
})
