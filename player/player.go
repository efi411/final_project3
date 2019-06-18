package player

import (
	cha "final_project3/channel"
	r "final_project3/resultStrings"
	"fmt"
	"strconv"
	"time"
)

//Player - Player in the game
type Player struct {
	username             string
	number               int
	ch                   cha.Channel
	otherPlayersChannels []cha.Channel
	allChan              []cha.Channel
	nodes                int
}

//ChannelsListNotGoodErrMsg - Error message
const ChannelsListNotGoodErrMsg = "Channels list is not good"

//NumberOfRounds - The number of rounds for the players loop
const NumberOfRounds = 6

//DefaultLeader - The leader that is being selected when there is a crash
const DefaultLeader = -1

//AliveMsg - Alive string message
const AliveMsg = "ALIVE"

//StartMsg - Start string message
const StartMsg = "START"

//New - Player constructor
func New(username string, number int, userChannel cha.Channel, channelsList []cha.Channel, allChan []cha.Channel, nodes int) (Player, error) {
	if channelsList == nil || len(channelsList) == 0 {
		return Player{}, fmt.Errorf("%s", ChannelsListNotGoodErrMsg)
	}
	e := Player{username, number, userChannel, channelsList, allChan, nodes}
	return e, nil
}

//-----------Public functions-----------

//UserToString - Returns a representation of the player
func (e Player) UserToString() string {
	return fmt.Sprintf("Username: %d, User number: %d", e.GetNumber(), e.number)
}

//GetUsername - return the username
func (e Player) GetUsername() string {
	return e.username
}

//GetNumber - return the random number of the user
func (e Player) GetNumber() int {
	return e.number
}

//GetotherPlayersChannels - return other players channels
func (e Player) GetotherPlayersChannels() []cha.Channel {
	return e.otherPlayersChannels
}

//GetNumberOfNodes - getting number of all nodes
func (e Player) GetNumberOfNodes() int {
	return e.nodes
}

//SetNumberOfNodes - setting number of all nodes
func (e *Player) SetNumberOfNodes(newSum int) {
	e.nodes = newSum
}

//SendMessagesToAllPlayers - Sends the random number of the user to all the others players channels
func (e Player) SendMessagesToAllPlayers(msg string, resultStr *r.ResultStrings) {
	for _, element := range e.otherPlayersChannels {
		e.sendMessage(element, msg)
		// Add message to the result object
		resultStr.AddMessage(fmt.Sprintf("%s,%s,%s", strconv.Itoa(e.GetNumber()), strconv.Itoa(element.GetID()), msg))
	}
}

//LeaderAlgo - execute the second algorithm
func (e Player) LeaderAlgo(alfa int, beta int, resultStr *r.ResultStrings) int {
	var Leader = DefaultLeader
	for i := 0; i < NumberOfRounds; i++ {
		time.Sleep(1 * time.Second)
		for _, element := range e.otherPlayersChannels {
			starts := e.ch.GetStartMsg()
			alives := e.ch.GetAliveMsg()
			if ((starts[element.GetID()]) > 0) || ((alives[element.GetID()]) > 0) {
				e.ch.InitialMessages(element.GetID())
				otherRound := element.GetRound()
				currentRound := e.ch.GetRound()
				if currentRound > otherRound {
					e.sendMessage(element, StartMsg)
					// Add message to the result object
					resultStr.AddMessage(fmt.Sprintf("%s,%s,%s", strconv.Itoa(e.GetNumber()), strconv.Itoa(element.GetID()), StartMsg))
				} else {
					if currentRound < otherRound {
						e.startRound(otherRound, e.GetNumberOfNodes(), resultStr)
						e.ch.SetRound(otherRound)
						currentRound = e.ch.GetRound()
					}
				}
			}
		}
		if alfa < i && beta > 0 {
			e.SetNumberOfNodes(e.GetNumberOfNodes() - 1)
			if e.GetNumber() != (e.ch.GetRound() % (e.GetNumberOfNodes() + 1)) {
				e.startRound(e.ch.GetRound()+1, e.GetNumberOfNodes(), resultStr)
				e.ch.SetRound(e.ch.GetRound() + 1)
			} else {
				return DefaultLeader
			}
		}
		if e.GetNumber() == (e.ch.GetRound() % e.GetNumberOfNodes()) {
			e.SendMessagesToAllPlayers(AliveMsg, resultStr)
		}
		Leader = (e.ch.GetRound() % e.GetNumberOfNodes())
	}
	return Leader
}

func (e Player) sendMessage(channel cha.Channel, msg string) {
	channel.InsertMessage(msg, e.ch.GetID())
}

func (e Player) startRound(s int, len int, resultStr *r.ResultStrings) {
	if e.GetNumber() != (s % len) {
		i := (s % len)
		element := e.allChan[i]
		e.sendMessage(element, "START")
		// Add message to the result object
		resultStr.AddMessage(fmt.Sprintf("%s,%s,%s", strconv.Itoa(e.GetNumber()), strconv.Itoa(element.GetID()), StartMsg))
	}

}
