package slices_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vulpes-ferrilata/slices"
)

var _ = Describe("Map", func() {
	Err := errors.New("error")

	DescribeTable("when slice of integer and mapper return string",
		func(slice []int, f slices.MapperFunc[int, string], expectedResults []string, expectedErr error) {
			results, err := slices.Map(f, slice...)
			if expectedErr == nil {
				Expect(err).ShouldNot(HaveOccurred())
			} else {
				Expect(err).Should(MatchError(err))
			}
			Expect(results).Should(HaveExactElements(expectedResults))

		},
		Entry(
			"when mapper function return static result",
			[]int{2, 3, 4, 5, 6},
			func(object int) (string, error) {
				return "a", nil
			},
			[]string{"a", "a", "a", "a", "a"},
			nil,
		),
		Entry(
			"when mapper function return dynamic result which convert from argument",
			[]int{2, 3, 4, 5, 6},
			func(object int) (string, error) {
				return fmt.Sprint(object), nil
			},
			[]string{"2", "3", "4", "5", "6"},
			nil,
		),
		Entry(
			"when predicate function return error",
			[]int{2, 3, 4, 5, 6},
			func(object int) (string, error) {
				return "", Err
			},
			[]string{},
			Err,
		),
	)
})
