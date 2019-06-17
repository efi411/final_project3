# final_project

In this repository, we implements - Robust Stabilizing Leader Election algorithm.
1. A manager creates players with a random number for each player (doesn't have to be uniqe).
2. The players send messages with their numbers to one another throgh channels.
3. After they get all the messages, they summerize the numbers and print the result.

Each time, the probability of successful recieving messages by the other players is random.
We expect the results to not always be equal. If the probability is low, we expect many messages to be lost and if its high, we expect fewer messages to be lost.



Code testing:<br>
go test ./...<br>
Coverage flag:<br>
-cover
