package slices_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vulpes-ferrilata/slices"
)

var _ = Describe("Find", func() {
	var Err = errors.New("error")

	When("slice of integer from 1 to 5", func() {
		slice := []int{1, 2, 3, 4, 5}

		Context("with predicate function matches anything", func() {
			result, err := slices.Find(func(object int) (bool, error) {
				return true, nil
			}, slice...)

			It("must return 1", func() {
				Expect(result).Should(Equal(1))
			})

			It("has no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with predicate function matchs 1", func() {
			result, err := slices.Find(func(object int) (bool, error) {
				return object == 1, nil
			}, slice...)

			It("must return 1", func() {
				Expect(result).Should(Equal(1))
			})

			It("has no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with predicate function matchs 5", func() {
			result, err := slices.Find(func(object int) (bool, error) {
				return object == 5, nil
			}, slice...)

			It("must return 5", func() {
				Expect(result).Should(Equal(5))
			})

			It("has no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with predicate function matchs nothing", func() {
			_, err := slices.Find(func(object int) (bool, error) {
				return false, nil
			}, slice...)

			It("must return error 'no match found'", func() {
				Expect(err).Should(MatchError(slices.ErrNoMatchFound))
			})
		})

		Context("with predicate function matchs 6", func() {
			_, err := slices.Find(func(object int) (bool, error) {
				return object == 6, nil
			}, slice...)

			It("must return error 'no match found'", func() {
				Expect(err).Should(MatchError(slices.ErrNoMatchFound))
			})
		})

		Context("with predicate function matches 0", func() {
			_, err := slices.Find(func(object int) (bool, error) {
				return object == 0, nil
			}, slice...)

			It("must return error 'no match found'", func() {
				Expect(err).Should(MatchError(slices.ErrNoMatchFound))
			})
		})

		Context("with predicate function return error", func() {
			_, err := slices.Find(func(object int) (bool, error) {
				return true, Err
			}, slice...)

			It("must return error", func() {
				Expect(err).Should(MatchError(Err))
			})
		})
	})
})
