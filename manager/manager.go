package manager

import (
	cha "final_project3/channel"
	p "final_project3/player"
	r "final_project3/resultStrings"
	"fmt"
	"strconv"
	"sync"
)

//Manager - The manager of the game
//@Singelton
type Manager struct {
	players  []p.Player
	channels []cha.Channel
}

var instance *Manager
var once sync.Once

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
func (m Manager) StartGame(numOfPlayers int, crash int, resStr *r.ResultStrings) error {

	m.printToConsole("--------Starting game--------")
	m.printToConsole("Adding players...")
	//Create channels
	for i := 0; i < numOfPlayers; i++ {
		//List for saving the alive messages from the other players
		alives := make([]int, numOfPlayers)
		//List for saving the start messages from the other players
		starts := make([]int, numOfPlayers)
		channel, _ := cha.New(i, starts, alives, i)
		addChannel(channel)
	}
	//Create players
	for i := 0; i < numOfPlayers; i++ {
		//Player gets his channel in the constructor
		e, _ := p.New(fmt.Sprintf("%s%d", "player", i), i, instance.channels[i], getChannelsListWithoutIndex(i), instance.channels, numOfPlayers)
		addPlayer(e)
	}

	//Running the algorithm in each of the players in parallel
	var wg sync.WaitGroup
	wg.Add(numOfPlayers)
	for _, element := range instance.players {
		//The algorithm runs on goroutines
		go func(e p.Player, resultStr *r.ResultStrings) {
			defer wg.Done()
			//Sending to the algorithm the relevant parmeters
			leader := e.LeaderAlgo(4, crash, resultStr)
			resultStr.AddLeader(strconv.Itoa(leader))
		}(element, resStr)
	}
	//Waiting for all the players to finish the algorithm and return an answer
	wg.Wait()

	m.printToConsole("--------Exiting game--------")
	return nil
}

//-----------Private functions-----------

//addPlayer - Add a player to the players list
func addPlayer(player p.Player) {
	instance.players = append(instance.players, player)
}

//addChannel - Add a channel to the channels list
func addChannel(ch cha.Channel) {
	instance.channels = append(instance.channels, ch)
}

//getChannelsListWithoutIndex - Returns the channels list without a certain index
func getChannelsListWithoutIndex(index int) []cha.Channel {
	var newChannelsList []cha.Channel
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
