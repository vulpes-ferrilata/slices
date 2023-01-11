package slices_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vulpes-ferrilata/slices"
)

var _ = Describe("Map", func() {
	var Err = errors.New("error")

	When("slice of integer from 1 to 5", func() {
		slice := []int{1, 2, 3, 4, 5}

		Context("with mapper function return string", func() {
			results, err := slices.Map(func(object int) (string, error) {
				return fmt.Sprint(object), nil
			}, slice...)

			It(`must return a slice contains "1", "2", "3", "4", "5"`, func() {
				Expect(results).Should(ContainElements("1", "2", "3", "4", "5"))
			})

			It("has no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with mapper function return error", func() {
			_, err := slices.Map(func(object int) (string, error) {
				return "", Err
			}, slice...)

			It("must return error", func() {
				Expect(err).Should(MatchError(Err))
			})
		})
	})
})
