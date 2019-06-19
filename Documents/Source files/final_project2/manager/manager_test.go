package manager_test

import (
	m "final_project2/manager"
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
				manager1 := m.GetInstance()
				manager2 := m.GetInstance()
				Expect(manager1).To(Equal(manager2))
			})
		})
		Context("Start game", func() {
			It("Start game successed", func() {
				manager := m.GetInstance()
				err := manager.StartGame(3, 1)
				Expect(err).To(BeNil())
			})
		})
	})
})
