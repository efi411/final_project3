# Robust Stabilizing Leader Election Algorithm

In this repository, we implements - Robust Stabilizing Leader Election Algorithm.

Abstract (from the article):<br>
We mix two approaches of the fault-tolerance: robustness and stabilization.<br>
Using these approaches, we propose leader election algorithms that tolerate both transient and crash failures.<br>
Our goal is to show the implementability of the robust self- and/or pseudo- stabilizing leader election in various systems with weak reliability and synchrony assumptions.<br>
We try to propose, when it is possible, communicationefficient implementations.<br>
Also, we exhibit some assumptions required to obtain robust stabilizing leader election algorithms.<br>
Our results show that the gap between robustness and stabilizing robustness is not really significant when we consider fix-point problems such as leader election.<br>
<br><br>
In order to use this library and run the leader election algorithm, follow these instructions:

1. Install go: Go to https://golang.org/dl/ and install Go on your computer.<br>
*You can learn more about Golang installation in this link - https://golang.org/doc/install*<br>
2. Install all the dependencies of the library: Run go get github.com/efi411/final_project3/<br>
3. Run the algorithm: Open ui/index.html file -> Choose if you want to have a crash or not and the number of palyers in the current run and click on "Find the leader". The results will appear on the screen.<br>


Code testing:<br>
go test ./...<br>
Coverage flag:<br>
-cover

Owners: Efrat Tsadok and Shira Mandelbaum
