package slices_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vulpes-ferrilata/slices"
)

var _ = Describe("Filter", func() {
	var Err = errors.New("error")

	When("slice of integer from 1 to 5", func() {
		slice := []int{1, 2, 3, 4, 5}

		Context("with predicate function matches everything", func() {
			results, err := slices.Filter(func(object int) (bool, error) {
				return true, nil
			}, slice...)

			It("must return a slice contains 1, 2, 3, 4, 5", func() {
				Expect(results).Should(ContainElements(1, 2, 3, 4, 5))
			})

			It("has no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with predicate function matchs every elements except 1", func() {
			results, err := slices.Filter(func(object int) (bool, error) {
				return object != 1, nil
			}, slice...)

			It("must return a slice contains 2, 3, 4, 5", func() {
				Expect(results).Should(ContainElements(2, 3, 4, 5))
			})

			It("has no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with predicate function matchs every elements except 5", func() {
			results, err := slices.Filter(func(object int) (bool, error) {
				return object != 5, nil
			}, slice...)

			It("must return a slice containes 1, 2, 3, 4", func() {
				Expect(results).Should(ContainElements(1, 2, 3, 4))
			})

			It("has no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with predicate function matchs nothing", func() {
			results, err := slices.Filter(func(object int) (bool, error) {
				return false, nil
			}, slice...)

			It("must return empty result", func() {
				Expect(results).Should(BeEmpty())
			})

			It("has no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with predicate function matchs 6", func() {
			results, err := slices.Filter(func(object int) (bool, error) {
				return object == 6, nil
			}, slice...)

			It("must return empty result", func() {
				Expect(results).Should(BeEmpty())
			})

			It("has no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with predicate function matches 0", func() {
			results, err := slices.Filter(func(object int) (bool, error) {
				return object == 0, nil
			}, slice...)

			It("must return empty result", func() {
				Expect(results).Should(BeEmpty())
			})

			It("no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with predicate function return error", func() {
			_, err := slices.Filter(func(object int) (bool, error) {
				return true, Err
			}, slice...)

			It("must return error", func() {
				Expect(err).Should(MatchError(Err))
			})
		})
	})
})
