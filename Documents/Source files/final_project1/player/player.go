package player

import (
	"fmt"
)

//Player - Player in the game
type Player struct {
	username             string
	number               int
	ch                   chan int
	otherPlayersChannels []chan int
}

var CHANELS_LIST_NOT_GOOD_ERR = "Channels list has no channels in it"

//New - Player constructor
func New(username string, number int, userChannel chan int, channelsList []chan int) (Player, error) {
	if channelsList == nil {
		return Player{}, fmt.Errorf("%s", CHANELS_LIST_NOT_GOOD_ERR)
	}
	e := Player{username, number, userChannel, channelsList}
	return e, nil
}

//-----------Public functions-----------

//UserToString - Returns a representation of the player
func (e Player) UserToString() string {
	return fmt.Sprintf("username: %s, user number: %d", e.username, e.number)
}

//GetUsername - return the username
func (e Player) GetUsername() string {
	return e.username
}

//GetRandomNumber - return the random number of the user
func (e Player) GetRandomNumber() int {
	return e.number
}

//GetChannel - return the channel of the user
func (e Player) GetChannel() chan int {
	return e.ch
}

//GetotherPlayersChannels - return other players channels
func (e Player) GetotherPlayersChannels() []chan int {
	return e.otherPlayersChannels
}

//SendMessagesToAllPlayers - Sends the random number of the user to all the others players channels
func (e Player) SendMessagesToAllPlayers() {
	for _, element := range e.otherPlayersChannels {
		e.sendNumber(element)
	}
}

//GetSum - get all the numbers from the channel, summerizes and prints it
func (e Player) GetSum(numOfPlayers int) string {
	close(e.ch)
	sum := e.number
	for elem := range e.ch {
		sum += elem
	}
	return fmt.Sprintf("Username: %s, Sum: %d", e.username, sum)
}

//-----------Private functions-----------

//sendNumber - Sends the random number of the user to the channel (argument)
func (e Player) sendNumber(channel chan int) {
	//e.number is always the same!!!
	channel <- e.number
}
