package slices_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vulpes-ferrilata/slices"
)

var _ = Describe("All", func() {
	var Err = errors.New("error")

	When("slice of integer contains 5 elements of repeated 2", func() {
		slice := []int{2, 2, 2, 2, 2}

		Context("with predicate function matchs anything", func() {
			isExists, err := slices.All(func(object int) (bool, error) {
				return true, nil
			}, slice...)

			It("must return true", func() {
				Expect(isExists).Should(BeTrue())
			})

			It("has no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with predicate function matchs 2", func() {
			isExists, err := slices.All(func(object int) (bool, error) {
				return object == 2, nil
			}, slice...)

			It("must return true", func() {
				Expect(isExists).Should(BeTrue())
			})

			It("has no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with predicate function matchs 3", func() {
			isExists, err := slices.All(func(object int) (bool, error) {
				return object == 3, nil
			}, slice...)

			It("must return false", func() {
				Expect(isExists).Should(BeFalse())
			})

			It("has no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with predicate function matchs 1", func() {
			isExists, err := slices.All(func(object int) (bool, error) {
				return object == 1, nil
			}, slice...)

			It("must return false", func() {
				Expect(isExists).Should(BeFalse())
			})

			It("has no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with predicate function matches nothing", func() {
			isExists, err := slices.All(func(object int) (bool, error) {
				return false, nil
			}, slice...)

			It("must return false", func() {
				Expect(isExists).Should(BeFalse())
			})

			It("has no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with predicate function return error", func() {
			_, err := slices.All(func(object int) (bool, error) {
				return true, Err
			}, slice...)

			It("must return error", func() {
				Expect(err).Should(MatchError(Err))
			})
		})
	})
})
