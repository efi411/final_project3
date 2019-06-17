package main

import (
	m "final_project3/manager"
	"fmt"
	"github.com/gin-contrib/cors"
	"math/rand"
	"net/http"
	"strings"
	"time"

	r "final_project3/resultStrings"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Use cors for go server (Cross-Origin Resource Sharing)
	r.Use(cors.Default())

	projectAPI := r.Group("/project")
	projectAPI.GET("/", handleRequest)

	r.Run()
}

func handleRequest(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.WriteString("3:2,3,Alive-4,2,Alive-5,1,Start-4,5,Alive-4,3,Start-4,1,Start-2,3,Alive-3,2,Start-1,3,Alive")
}

func runAlgorithmAndGetResults(numOfPlayers int, crash int) string {
	leader, results := runScenario(numOfPlayers, crash)
	stringToClient := fmt.Sprintf("The leader is: %s\nMessages: %s", leader, results)
	return stringToClient
}

func runScenario(numOfPlayers int, crash int) (string, string){
	//Create unique seed for this program - different random numbers every time
	rand.Seed(time.Now().UnixNano())

	//Create the manager and start the game
	manager := m.GetInstance()

	//Create the string result object
	res := r.NewResultString()

	var leader string
	var scenarioResults string
	//Arguments: Number of players, Probability of loosing messages
	if err := manager.StartGame(numOfPlayers, crash, &res); err != nil {
		fmt.Println("Something went wrong!")
	} else {
		resultStr, lead := res.GetResultString()
		scenarioResults = resultStr

		leadersArray := strings.Split(lead, ":")

		if crash == 1 {
			leader = "1"
		} else {
			leader = leadersArray[0]
		}
	}
	return leader, scenarioResults
}
