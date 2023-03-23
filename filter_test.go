package slices_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vulpes-ferrilata/slices"
)

var _ = Describe("Filter", func() {
	Err := errors.New("error")

	DescribeTable("when slice of integer",
		func(slice []int, f slices.PredicateFunc[int], expectedResults []int, expectedErr error) {
			results, err := slices.Filter(f, slice...)
			if expectedErr == nil {
				Expect(err).ShouldNot(HaveOccurred())
			} else {
				Expect(err).Should(MatchError(err))
			}
			Expect(results).Should(HaveExactElements(expectedResults))

		},
		Entry(
			"when predicate function return true",
			[]int{2, 3, 4, 5, 6},
			func(object int) (bool, error) {
				return true, nil
			},
			[]int{2, 3, 4, 5, 6},
			nil,
		),
		Entry(
			"when predicate function return false",
			[]int{2, 3, 4, 5, 6},
			func(object int) (bool, error) {
				return false, nil
			},
			[]int{},
			nil,
		),
		Entry(
			"when predicate function match value in slice",
			[]int{2, 3, 4, 5, 6},
			func(object int) (bool, error) {
				return object == 2, nil
			},
			[]int{2},
			nil,
		),
		Entry(
			"when predicate function match some values in slice",
			[]int{2, 3, 4, 5, 6},
			func(object int) (bool, error) {
				return object > 3, nil
			},
			[]int{4, 5, 6},
			nil,
		),
		Entry(
			"when predicate function return error",
			[]int{2, 3, 4, 5, 6},
			func(object int) (bool, error) {
				return true, Err
			},
			[]int{},
			Err,
		),
	)
})
