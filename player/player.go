package player

import (
	cha "final_project3/channel"
	"fmt"
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
const ChannelsListNotGoodErrMsg = "Channels propbability must be between (0,1]"

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

// //GetChannel - return the channel of the user
// func (e Player) GetChannel() chan int {
// 	return e.ch.GetChannel()
// }

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
func (e Player) SendMessagesToAllPlayers(msg string) {
	//countLostMessages := 0
	for _, element := range e.otherPlayersChannels {
		e.sendMessage(element, msg)
	}
}

//LeaderAlgo - execute the second algorithm
func (e Player) LeaderAlgo(alfa int, beta int) int {
	//var a = alfa
	//var b = beta
	//var d = delta
	var Leader = -1
	//e.GetChannel() <- round
	//x := <-e.GetChannel()
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		for _, element := range e.otherPlayersChannels {
			starts := e.ch.GetStartMsg()
			alives := e.ch.GetAliveMsg()
			//s := strings.Split(msg, ",")
			//messageFrom, otherRound := s[0], s[1]
			if ((starts[element.GetID()]) > 0) || ((alives[element.GetID()]) > 0) {
				fmt.Printf("run:%d, player:%d in the condition, get message from:%d\n", i, e.GetNumber(), element.GetID())
				e.ch.InitialMessages(element.GetID())
				otherRound := element.GetRound()
				currentRound := e.ch.GetRound()
				fmt.Printf("run:%d, player:%d, my round is:%d, other round is:%d\n", i, e.GetNumber(), currentRound, otherRound)
				if currentRound > otherRound {
					fmt.Printf("run:%d, player:%d send start I am the leader to, %d \n", i, e.GetNumber(), element.GetID())
					e.sendMessage(element, "START")
				} else {

					if currentRound < otherRound {
						fmt.Printf("run:%d, player:%d is going to start new round, %d\n", i, e.GetNumber(), otherRound)
						e.startRound(otherRound, e.GetNumberOfNodes())
						e.ch.SetRound(otherRound)
						fmt.Printf("run:%d, player:%d change is round to, %d\n", e.GetNumber(), i, e.ch.GetRound())
						currentRound = e.ch.GetRound()
					} else {
						fmt.Printf("run:%d, player:%d and number, %d are equals\n", i, e.GetNumber(), element.GetID())
					}
				}
			}
		}
		if alfa < i && beta > 0 {
			fmt.Printf("run:%d, player:%d, alfa is %d\n", i, e.GetNumber(), alfa)
			fmt.Printf("run:%d, player:%d, player %d is now crashing\n", i, e.GetNumber(), (e.ch.GetRound() % e.GetNumberOfNodes()))
			e.SetNumberOfNodes(e.GetNumberOfNodes() - 1)
			if e.GetNumber() != (e.ch.GetRound() % (e.GetNumberOfNodes() + 1)) {
				e.startRound(e.ch.GetRound()+1, e.GetNumberOfNodes())
				e.ch.SetRound(e.ch.GetRound() + 1)
				fmt.Printf("run:%d, player:%d, after return from function the number of nodes is: %d\n", i, e.GetNumber(), e.GetNumberOfNodes())
			} else {
				fmt.Printf("run:%d, player:%d, the current leader %d exit from game\n", i, e.GetNumber(), e.ch.GetRound()%e.GetNumberOfNodes())
				return -1
			}
		}
		fmt.Printf("run:%d, player:%d, after end if the number of nodes is: %d\n", i, e.GetNumber(), e.GetNumberOfNodes())
		if e.GetNumber() == (e.ch.GetRound() % e.GetNumberOfNodes()) {
			fmt.Printf("run:%d, player:%d, send alive to all players, my current round is:%d\n", i, e.GetNumber(), e.ch.GetRound())
			e.SendMessagesToAllPlayers("ALIVE")
		}
		fmt.Printf("run:%d, player:%d, before initial leader the round is:%d, and the nodes is:%d\n", i, e.GetNumber(), e.ch.GetRound(), e.GetNumberOfNodes())
		Leader = (e.ch.GetRound() % e.GetNumberOfNodes())
		fmt.Printf("run:%d, player:%d, my leader in run %d, is :%d\n", i, e.GetNumber(), i, Leader)
	}
	return Leader
}

func (e Player) sendMessage(channel cha.Channel, msg string) {
	channel.InsertMessage(msg, e.ch.GetID())
}

func (e Player) startRound(s int, len int) {
	fmt.Printf("number %d start round number, %d\n", e.GetNumber(), s)
	fmt.Printf("the len of nodes is %d\n", len)
	if e.GetNumber() != (s % len) {
		i := (s % len)
		fmt.Printf("number %d start round number, %d, the result of mod is,%d\n", e.GetNumber(), s, i)
		element := e.allChan[i]
		fmt.Printf("number %d send start i am with you to, %d\n", e.GetNumber(), element.GetID())
		e.sendMessage(element, "START")
	}

}
