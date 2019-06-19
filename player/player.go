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

//GetNumberOfNodes - getting number of all the nodes
func (e Player) GetNumberOfNodes() int {
	return e.nodes
}

//SetNumberOfNodes - setting number of all the nodes
func (e *Player) SetNumberOfNodes(newSum int) {
	e.nodes = newSum
}

//SendMessagesToAllPlayers - Sends "ALIVE" or "START" to all the others players channels
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
	//Limit the algorithm run time
	for i := 0; i < NumberOfRounds; i++ {
		time.Sleep(1 * time.Second)
		//Check if receive a message from another player
		for _, element := range e.otherPlayersChannels {
			starts := e.ch.GetStartMsg()
			alives := e.ch.GetAliveMsg()
			if ((starts[element.GetID()]) > 0) || ((alives[element.GetID()]) > 0) {
				e.ch.InitialMessages(element.GetID())
				//Get the round of the sender
				otherRound := element.GetRound()
				currentRound := e.ch.GetRound()
				//Check which round is bigger
				if currentRound > otherRound {
					//Send "START" message if my round is bigger than sender round
					e.sendMessage(element, StartMsg)
					// Add message to the result object
					resultStr.AddMessage(fmt.Sprintf("%s,%s,%s", strconv.Itoa(e.GetNumber()), strconv.Itoa(element.GetID()), StartMsg))
				} else {
					if currentRound < otherRound {
						//Start the bigger round in case the sender round is bigger than mine
						e.startRound(otherRound, e.GetNumberOfNodes(), resultStr)
						//Update the round variables
						e.ch.SetRound(otherRound)
						currentRound = e.ch.GetRound()
					}
				}
			}
		}
		//Case of crash
		if alfa < i && beta > 0 {
			//decrease the number of players-nodes
			e.SetNumberOfNodes(e.GetNumberOfNodes() - 1)
			//Check the leader before the crashing
			if e.GetNumber() != (e.ch.GetRound() % (e.GetNumberOfNodes() + 1)) {
				//Start the next round after leader fail
				e.startRound(e.ch.GetRound()+1, e.GetNumberOfNodes(), resultStr)
				//Update the round variable
				e.ch.SetRound(e.ch.GetRound() + 1)
			} else {
				//Case I'am the leader that crashed
				return DefaultLeader
			}
		}
		//Case I'am the leader- send "ALIVE" message to all the players
		if e.GetNumber() == (e.ch.GetRound() % e.GetNumberOfNodes()) {
			e.SendMessagesToAllPlayers(AliveMsg, resultStr)
		}
		//Update the current leader
		Leader = (e.ch.GetRound() % e.GetNumberOfNodes())
	}
	return Leader
}

//-----------Private functions-----------

//sendMessage - send "ALIVE" or "START" to specific user
func (e Player) sendMessage(channel cha.Channel, msg string) {
	channel.InsertMessage(msg, e.ch.GetID())
}

//startRound - start new round
func (e Player) startRound(s int, len int, resultStr *r.ResultStrings) {
	//check if i'am the leader of the new round
	if e.GetNumber() != (s % len) {
		i := (s % len)
		//get the id of the new leader
		element := e.allChan[i]
		//send "START" message to the new leader
		e.sendMessage(element, "START")
		// Add message to the result object
		resultStr.AddMessage(fmt.Sprintf("%s,%s,%s", strconv.Itoa(e.GetNumber()), strconv.Itoa(element.GetID()), StartMsg))
	}

}
