package slices_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vulpes-ferrilata/slices"
)

var _ = Describe("Remove", func() {
	DescribeTable("when slice of integer",
		func(slice []int, elements []int, expectedResults []int) {
			results := slices.Remove(slice, elements...)
			Expect(results).Should(BeEquivalentTo(expectedResults))

		},
		Entry(
			"when element match first value of the slice",
			[]int{2, 3, 4, 5, 6},
			[]int{2},
			[]int{3, 4, 5, 6},
		),
		Entry(
			"when element match last value of the slice",
			[]int{2, 3, 4, 5, 6},
			[]int{6},
			[]int{2, 3, 4, 5},
		),
		Entry(
			"when element match some values of the slice",
			[]int{2, 3, 4, 5, 6},
			[]int{2, 4, 6},
			[]int{3, 5},
		),
	)
})
