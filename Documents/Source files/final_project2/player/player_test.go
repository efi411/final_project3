package player_test

import (
	cha "final_project2/channel"
	p "final_project2/player"
	"testing"

	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	TestUsername     = "testUser"
	TestRandomNumber = 15
)

var (
	channelPlayer cha.Channel
	channels      []cha.Channel
)

func TestPlayers(t *testing.T) {

	RegisterFailHandler(Fail)
	RunSpecs(t, "Player Suite")
}

var _ = Describe("Player", func() {

	BeforeEach(func() {
		channelPlayer, channels = createPlayersChannels(1)
	})

	Describe("Check player's functionality", func() {
		Context("Create player", func() {
			It("Create successed", func() {
				testPlayer, err := p.New(TestUsername, TestRandomNumber, channelPlayer, channels)
				Expect(err).To(BeNil())
				Expect(testPlayer.GetUsername()).To(Equal(TestUsername))
				Expect(testPlayer.GetRandomNumber()).To(Equal(TestRandomNumber))
				Expect(testPlayer.UserToString()).To(Equal(fmt.Sprintf("Username: %s, User number: %d", TestUsername, TestRandomNumber)))
				Expect(testPlayer.GetChannel()).ToNot(BeNil())
				channelsList := testPlayer.GetotherPlayersChannels()
				Expect(channelsList).ToNot(BeNil())
				Expect(len(channelsList)).To(Equal(2))
			})
			It("Create with no channels list", func() {
				_, err := p.New(TestUsername, TestRandomNumber, channelPlayer, nil)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal(p.ChannelsListNotGoodErrMsg))
			})
			It("Create with empty channels list", func() {
				var emptyChannels []cha.Channel
				_, err := p.New(TestUsername, TestRandomNumber, channelPlayer, emptyChannels)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal(p.ChannelsListNotGoodErrMsg))
			})
		})
		Context("Send messages to other players", func() {
			It("Send messages and retrieve sum from other channels with probability of 0", func() {
				testPlayer, _ := p.New(TestUsername, TestRandomNumber, channelPlayer, channels)
				amountLostMsg := testPlayer.SendMessagesToAllPlayers()
				Expect(amountLostMsg).To(Equal(0))
				channelsList := testPlayer.GetotherPlayersChannels()
				for _, channelElement := range channelsList {
					Expect(channelElement.GetSum()).To(Equal(TestRandomNumber))
				}
			})
			It("Send messages and retrieve sum from other channels with probability of 0", func() {
				channelPlayer, channels = createPlayersChannels(0)
				testPlayer, _ := p.New(TestUsername, TestRandomNumber, channelPlayer, channels)
				amountLostMsg := testPlayer.SendMessagesToAllPlayers()
				Expect(amountLostMsg).To(Equal(2))
			})
		})
		Context("Send message the same player", func() {
			It("Send message to the same player's channel and get sum", func() {
				testPlayerChannel, _ := cha.New(1, make(chan int, 2))
				var playersChannels []cha.Channel
				playersChannels = append(playersChannels, testPlayerChannel)
				testPlayer, _ := p.New(TestUsername, TestRandomNumber, testPlayerChannel, playersChannels)
				testPlayer.SendMessagesToAllPlayers()
				Expect(testPlayer.GetSum()).To(Equal(TestRandomNumber * 2))
			})
		})
	})
})

func createPlayersChannels(prob float64) (cha.Channel, []cha.Channel) {
	//Create player channel
	channel1, _ := cha.New(prob, make(chan int, 2))
	//Create othe players channels
	channel2, _ := cha.New(prob, make(chan int, 2))
	channel3, _ := cha.New(prob, make(chan int, 2))
	//Create a list of channels and add channels to the list
	var playersChannels []cha.Channel
	playersChannels = append(playersChannels, channel2)
	playersChannels = append(playersChannels, channel3)

	return channel1, playersChannels
}
