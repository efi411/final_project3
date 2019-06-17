package channel_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestChannels(t *testing.T) {

	RegisterFailHandler(Fail)
	RunSpecs(t, "Channel Suite")
}

var _ = Describe("Channel", func() {
	Describe("Check channel's functionality", func() {
		Context("Create channel", func() {
			It("Create successed", func() {
			})
		})
	})
})
