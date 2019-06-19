package main

import (
	m "final_project2/manager"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Create uniq seed for this program - different random numbers every time
	rand.Seed(time.Now().UnixNano())

	//Create the manager and start the game
	manager := m.GetInstance()

	//Arguments: Number of players, Probability of loosing messages
	if err := manager.StartGame(5, 0.9); err != nil {
		fmt.Println("Something went wrong!")
	}
}
