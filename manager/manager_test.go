package manager_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestChannels(t *testing.T) {

	RegisterFailHandler(Fail)
	RunSpecs(t, "Manager Suite")
}

var _ = Describe("Manager", func() {
	Describe("Check manager's functionality", func() {
		Context("Create manager", func() {
			It("Create two managers and check that we are getting the same one", func() {
			})
		})
	})
})
