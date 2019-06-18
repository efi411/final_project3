package player_test

import (
"testing"

. "github.com/onsi/ginkgo"
. "github.com/onsi/gomega"
)

func TestPlayers(t *testing.T) {

	RegisterFailHandler(Fail)
	RunSpecs(t, "Player Suite")
}

var _ = Describe("Player", func() {
	Describe("Check player's functionality", func() {
		Context("Create player", func() {
			It("Create successed", func() {
			})
		})
	})
})