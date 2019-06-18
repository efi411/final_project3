package main_test

import (
	"testing"
	"time"

	m "final_project3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const timeLimitForAlgo = 7

func TestPlayers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Main Suite")
}

var _ = Describe("Leader election algorithm research", func() {
	Describe("Check leader election algorithm", func() {
		Context("Run algorithm with no crashes - Time, leader and messages string size", func() {
			It("5 players", func() {
				execTime, leader, results := m.RunAlgoAndCalcExecTime(25, 0)
				Expect(leader).To(Equal("24"))
				Ω(len(results)).Should(BeNumerically(">", 1000), "There are too few messages.")
				Ω(time.Minute * timeLimitForAlgo).Should(BeNumerically(">", execTime), "The algorithm shouldn't take too long.")
			})
		})
	})
})
