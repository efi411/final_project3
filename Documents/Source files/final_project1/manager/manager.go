package manager

import (
	p "final_project1/player"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

//Manager - The manager of the game
//@Singelton
type Manager struct {
	players  []p.Player
	channels []chan int
}

var instance *Manager
var once sync.Once

//Concurrency wait group - for 'go' function (threads)
//var waitGroup sync.WaitGroup

//GetInstance - Get instance of the manager - Manager constructor
func GetInstance() *Manager {
	//Syncronize creation and return
	once.Do(func() {
		instance = &Manager{}
	})
	return instance
}

//-----------Public functions-----------

//StartGame - Starts the game
func (m Manager) StartGame(numOfPlayers int) {
	//Create uniq seed for this program - different random numbers every time
	rand.Seed(time.Now().UnixNano())

	m.printToConsole("-----------Starts the game...-----------")
	m.printToConsole("Adding players...")
	//Create channels
	for i := 0; i < numOfPlayers; i++ {
		addChannel(make(chan int, numOfPlayers))
	}
	//Create players
	for i := 0; i < numOfPlayers; i++ {
		e, _ := p.New(fmt.Sprintf("%s%d", "player", i), randomNumber(), instance.channels[i], getChannelsListWithoutIndex(i))
		addPlayer(e)
	}
	m.printToConsole("The players in the currnt game:\n" + m.getPlayersList())
	//Exchange messages between players
	m.printToConsole("Exchange messages...")
	for i := 0; i < numOfPlayers; i++ {
		instance.players[i].SendMessagesToAllPlayers()
	}
	//Print sums
	m.printToConsole("Print sums...")
	for i := 0; i < numOfPlayers; i++ {
		m.printToConsole(instance.players[i].GetSum(numOfPlayers))
	}
	m.printToConsole("-----------Exiting game...-----------")
}

//-----------Private functions-----------

//AddPlayer - Add a player to the players list
func addPlayer(player p.Player) {
	instance.players = append(instance.players, player)
}

//AddPlayers - Add a player to the players list
func addPlayers(players []p.Player) {
	instance.players = append(instance.players, players...)
}

//addChannel - Add a channel to the channels list
func addChannel(ch chan int) {
	instance.channels = append(instance.channels, ch)
}

//randomNumber - Create and returns a random number
func randomNumber() int {
	return rand.Intn(100) + 1
}

//returns the channels list without a certain index
func getChannelsListWithoutIndex(index int) []chan int {
	var newChannelsList []chan int
	for i, element := range instance.channels {
		if i != index {
			newChannelsList = append(newChannelsList, element)
		}
	}
	return newChannelsList
}

//getPlayersList - Returns the players list
func (m Manager) getPlayersList() string {
	playersList := ""
	for i, player := range instance.players {
		playersList += strconv.Itoa(i)
		playersList += ": " + player.UserToString() + "\n"
	}
	return playersList
}

//printToConsole - Prints strings to the debug console
func (m Manager) printToConsole(toPrint string) {
	fmt.Println(toPrint)
}
