package channel

import (
	"fmt"
	"math/rand"
)

//Channel - Channel in the game
type Channel struct {
	probability float64
	ch          chan int
}

//ChannelsProbNotGoodErrMsg - Error message
const ChannelsProbNotGoodErrMsg = "Channels propbability must be between (0,1]"

//MessageLostErrMsg - Error message
const MessageLostErrMsg = "The message got lost in the way"

//New - Channel constructor
func New(probability float64, channel chan int) (Channel, error) {
	if probability > 1 || probability <= 0 {
		return Channel{}, fmt.Errorf("%s", ChannelsProbNotGoodErrMsg)
	}
	c := Channel{probability, channel}
	return c, nil
}

//-----------Public functions-----------

//GetChannel - return the channel of the user
func (c Channel) GetChannel() chan int {
	return c.ch
}

//InsertNumber - Insert number from the player to the channel
func (c Channel) InsertNumber(number int) error {
	randomNumber := rand.Float64()
	if randomNumber < c.probability {
		c.ch <- number
		return nil
	}
	return fmt.Errorf("%s", MessageLostErrMsg)
}

//GetSum - get all the numbers from the channel, summerizes and prints it
func (c Channel) GetSum() int {
	close(c.ch)
	sum := 0
	for elem := range c.ch {
		sum += elem
	}
	return sum
}
