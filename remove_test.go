package slices_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vulpes-ferrilata/slices"
)

var _ = Describe("Remove", func() {
	When("slice of integer from 1 to 5", func() {
		slice := []int{1, 2, 3, 4, 5}

		Context("remove element 1", func() {
			newSlice := slices.Remove(slice, 1)

			It("must return a slice contains 2, 3, 4, 5", func() {
				Expect(newSlice).Should(ContainElements(2, 3, 4, 5))
			})
		})

		Context("remove element 5", func() {
			newSlice := slices.Remove(slice, 5)

			It("must return a slice contains 1, 2, 3, 4", func() {
				Expect(newSlice).Should(ContainElements(1, 2, 3, 4))
			})
		})

		Context("remove element 3", func() {
			newSlice := slices.Remove(slice, 3)

			It("must contains 1, 2, 4, 5", func() {
				Expect(newSlice).Should(ContainElements(1, 2, 4, 5))
			})
		})

		Context("remove element 6", func() {
			newSlice := slices.Remove(slice, 6)

			It("must return a slice contains 1, 2, 3, 4, 5", func() {
				Expect(newSlice).Should(ContainElements(1, 2, 3, 4, 5))
			})
		})

		Context("remove element 0", func() {
			newSlice := slices.Remove(slice, 0)

			It("must return a slice contains 1, 2, 3, 4, 5", func() {
				Expect(newSlice).Should(ContainElements(1, 2, 3, 4, 5))
			})
		})
	})
})
