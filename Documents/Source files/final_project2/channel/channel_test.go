package channel_test

import (
	cha "final_project2/channel"
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
				channel := make(chan int, 1)
				testChannel, err := cha.New(1, channel)
				Expect(err).To(BeNil())
				Expect(testChannel.GetChannel()).ToNot(BeNil())
			})
			It("Create with high probability", func() {
				channel := make(chan int, 1)
				_, err := cha.New(1.5, channel)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal(cha.ChannelsProbNotGoodErrMsg))
			})

			It("Create with low probability", func() {
				channel := make(chan int, 1)
				_, err := cha.New(-3.4, channel)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal(cha.ChannelsProbNotGoodErrMsg))
			})

			It("Create with zero probability", func() {
				channel := make(chan int, 1)
				_, err := cha.New(0, channel)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal(cha.ChannelsProbNotGoodErrMsg))
			})
		})
		Context("Insert number to channel", func() {
			It("Insert number with probability 1", func() {
				channel := make(chan int, 1)
				testChannel, _ := cha.New(1, channel)
				err := testChannel.InsertNumber(5)
				Expect(err).To(BeNil())
				Expect(<-testChannel.GetChannel()).To(Equal(5))
			})
			It("Insert number with probability 0", func() {
				channel := make(chan int, 1)
				testChannel, _ := cha.New(0, channel)
				err := testChannel.InsertNumber(5)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal(cha.MessageLostErrMsg))
			})
		})
		Context("Get sum from channel", func() {
			It("Get sum", func() {
				channel := make(chan int, 3)
				testChannel, _ := cha.New(1, channel)
				testChannel.InsertNumber(1)
				testChannel.InsertNumber(2)
				testChannel.InsertNumber(3)
				Expect(testChannel.GetSum()).To(Equal(6))
			})
		})
	})
})
