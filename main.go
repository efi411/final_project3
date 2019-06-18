package main

import (
	m "final_project3/manager"
	"fmt"
	"github.com/gin-contrib/cors"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	p "final_project3/preRunResults"
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
	//Save the pre run results for the algorithm
	preRunResults := p.NewPreRunResults()

	//Get num of players
	numOfPlayers, _ := c.Request.URL.Query()["numOfPlayers"]
	num := numOfPlayers[0]
	numOfPlayersInt, _ := strconv.Atoi(num)
	//Get crashes status
	crashes, _ := c.Request.URL.Query()["crashes"]
	crash := crashes[0]
	crashInt, _ := strconv.Atoi(crash)

	resultsForUi := preRunResults.GetFullResults(numOfPlayersInt, crashInt)
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.WriteString(resultsForUi)
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
