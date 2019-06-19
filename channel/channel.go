package channel

import (
	"fmt"
)

//Channel - Channel in the game
type Channel struct {
	id       int
	startMsg []int
	aliveMsg []int
	round    int
}

//ChannelErrMsg - Error message
const ChannelErrMsg = "the id or the round is invalid"

//New - Channel constructor
func New(id int, startMsg []int, aliveMsg []int,
	round int) (Channel, error) {
	//Checking the validation of the channel
	if id < 0 || round < 0 {
		return Channel{}, fmt.Errorf("%s", ChannelErrMsg)
	}
	c := Channel{id, startMsg, aliveMsg, round}
	return c, nil
}

//-----------Public functions-----------

//GetRound - return the current channel-player round
func (c Channel) GetRound() int {
	return c.round
}

//GetID - return channel id
func (c Channel) GetID() int {
	return c.id
}

//SetRound - setting current channel-player round
func (c *Channel) SetRound(i int) {
	c.round = i
}

//GetStartMsg - return the array of start messages
func (c Channel) GetStartMsg() []int {
	return c.startMsg
}

//GetAliveMsg - return the array of alive messages
func (c Channel) GetAliveMsg() []int {
	return c.aliveMsg
}

//InsertMessage - insert the start or alive message to array
func (c Channel) InsertMessage(msg string, fromCh int) {
	//Check the type of the message and add in the sender id index
	if msg == "ALIVE" {
		c.aliveMsg[fromCh] = 1
	} else {
		c.startMsg[fromCh] = 1
	}
}

//InitialMessages - initialize the array of messages
func (c Channel) InitialMessages(id int) {
	c.aliveMsg[id] = 0
	c.startMsg[id] = 0
}
