package main

import (
	"fmt"
	"log"

	"gitlab.utc.fr/aguilber/ia04/agt/restclientagent"
	"gitlab.utc.fr/aguilber/ia04/agt/restserveuragent"
	"gitlab.utc.fr/aguilber/ia04/comsoc"
)

func main() {
	const url1 = ":8080"
	const url2 = "http://localhost:8080"

	serveurAgt := restserveuragent.NewRestServeurAgent(url1)
	log.Println("Démarrage des clients votants...")
	go serveurAgt.Start()

	// Voters
	var prefs1 = []comsoc.Alternative{1, 2, 5, 12}
	var prefs2 = []comsoc.Alternative{2, 12, 1, 5}
	var prefs3 = []comsoc.Alternative{2, 5, 12, 1}
	var prefs4 = []comsoc.Alternative{12, 1, 5, 2}
	var prefs5 = []comsoc.Alternative{5, 12, 1, 2}

	// -------------- TEST NUMERO 1 : 5 VOTANTS, Vote majoritaire simple -------------- //
	fmt.Println("")
	fmt.Println("")
	fmt.Println("------------ MAJORITE SIMPLE ------------")

	// Ballot
	fmt.Println("")
	fmt.Println("Ballot : ")
	voter_ids := []string{"ag_id1", "ag_id2", "ag_id3", "ag_id4", "ag_id5"}
	agNewBallot := restclientagent.NewBallotAgent("idb1", url2, "majority", "Wed Nov 30 23:00:00 UTC 2022", voter_ids, 1)
	agNewBallot.StartBallotAgent()

	fmt.Println(serveurAgt.GetBallots()[len(serveurAgt.GetBallots())-1])

	// Votants
	fmt.Println("")
	fmt.Println("Votants : ")
	agtNewVote1 := restclientagent.NewVoterAgent("idv1", url2, "ag_id1", "vote1", prefs1, nil)
	agtNewVote1.StartVoter()

	agtNewVote2 := restclientagent.NewVoterAgent("idv2", url2, "ag_id2", "vote1", prefs2, nil)
	agtNewVote2.StartVoter()

	agtNewVote3 := restclientagent.NewVoterAgent("idv3", url2, "ag_id3", "vote1", prefs3, nil)
	agtNewVote3.StartVoter()

	agtNewVote4 := restclientagent.NewVoterAgent("idv4", url2, "ag_id4", "vote1", prefs4, nil)
	agtNewVote4.StartVoter()

	agtNewVote5 := restclientagent.NewVoterAgent("idv5", url2, "ag_id5", "vote1", prefs5, nil)
	agtNewVote5.StartVoter()

	for i := 0; i < 5; i++ {
		fmt.Println(serveurAgt.GetVoters()[i])
	}

	// Result
	fmt.Println("")
	fmt.Println("Résultat : ")
	agResult := restclientagent.NewResultAgent("idr1", url2, "vote1")
	agResult.StartResultAgent()
	fmt.Println("")

	// -------------- TEST NUMERO 2 : 5 VOTANTS, Borda -------------- //
	fmt.Println("")
	fmt.Println("")
	fmt.Println("------------ BORDA ------------")

	// Ballot
	fmt.Println("")
	fmt.Println("Ballot : ")
	voter_ids = []string{"ag_id6", "ag_id7", "ag_id8", "ag_id9", "ag_id10"}
	agNewBallot = restclientagent.NewBallotAgent("idb2", url2, "borda", "Wed Nov 30 23:00:00 UTC 2022", voter_ids, 2)
	agNewBallot.StartBallotAgent()

	fmt.Println(serveurAgt.GetBallots()[len(serveurAgt.GetBallots())-1])

	// Votants
	fmt.Println("")
	fmt.Println("Votants : ")
	agtNewVote6 := restclientagent.NewVoterAgent("idv6", url2, "ag_id6", "vote2", prefs1, nil)
	agtNewVote6.StartVoter()

	agtNewVote7 := restclientagent.NewVoterAgent("idv7", url2, "ag_id7", "vote2", prefs2, nil)
	agtNewVote7.StartVoter()

	agtNewVote8 := restclientagent.NewVoterAgent("idv8", url2, "ag_id8", "vote2", prefs3, nil)
	agtNewVote8.StartVoter()

	agtNewVote9 := restclientagent.NewVoterAgent("idv9", url2, "ag_id9", "vote2", prefs4, nil)
	agtNewVote9.StartVoter()

	agtNewVote10 := restclientagent.NewVoterAgent("idv10", url2, "ag_id10", "vote2", prefs5, nil)
	agtNewVote10.StartVoter()

	for i := 5; i < 10; i++ {
		fmt.Println(serveurAgt.GetVoters()[i])
	}

	// Result
	fmt.Println("")
	fmt.Println("Résultat : ")
	agResult = restclientagent.NewResultAgent("idr2", url2, "vote2")
	agResult.StartResultAgent()

	// -------------- TEST NUMERO 3 : 5 VOTANTS, vote par approbation -------------- //
	fmt.Println("")
	fmt.Println("")
	fmt.Println("------------ APPROBATION ------------")

	// Ballot
	fmt.Println("")
	fmt.Println("Ballot : ")
	voter_ids = []string{"ag_id11", "ag_id12", "ag_id13", "ag_id14", "ag_id15"}
	agNewBallot = restclientagent.NewBallotAgent("idb3", url2, "approval", "Wed Nov 30 23:00:00 UTC 2022", voter_ids, 3)
	agNewBallot.StartBallotAgent()

	fmt.Println(serveurAgt.GetBallots()[len(serveurAgt.GetBallots())-1])

	// Votants
	fmt.Println("")
	fmt.Println("Votants : ")
	var tab2 = []int{2}
	var tab3 = []int{3}

	agtNewVote11 := restclientagent.NewVoterAgent("idv11", url2, "ag_id11", "vote3", prefs1, tab2)
	agtNewVote11.StartVoter()

	agtNewVote12 := restclientagent.NewVoterAgent("idv12", url2, "ag_id12", "vote3", prefs2, tab2)
	agtNewVote12.StartVoter()

	agtNewVote13 := restclientagent.NewVoterAgent("idv13", url2, "ag_id13", "vote3", prefs3, tab3)
	agtNewVote13.StartVoter()

	agtNewVote14 := restclientagent.NewVoterAgent("idv14", url2, "ag_id14", "vote3", prefs4, tab2)
	agtNewVote14.StartVoter()

	agtNewVote15 := restclientagent.NewVoterAgent("idv15", url2, "ag_id15", "vote3", prefs5, tab3)
	agtNewVote15.StartVoter()

	for i := 10; i < 15; i++ {
		fmt.Println(serveurAgt.GetVoters()[i])
	}

	// Result
	fmt.Println("")
	fmt.Println("Résultat : ")
	agResult = restclientagent.NewResultAgent("idr3", url2, "vote3")
	agResult.StartResultAgent()

	// -------------- TEST NUMERO 4 : 5 VOTANTS, gagnant de Condorcet -------------- //
	fmt.Println("")
	fmt.Println("")
	fmt.Println("------------ CONDORCET ------------")

	// Ballot
	fmt.Println("")
	fmt.Println("Ballot : ")
	voter_ids = []string{"ag_id16", "ag_id17", "ag_id18", "ag_id19", "ag_id20"}
	agNewBallot = restclientagent.NewBallotAgent("idb4", url2, "condorcet", "Wed Nov 30 23:00:00 UTC 2022", voter_ids, 4)
	agNewBallot.StartBallotAgent()

	fmt.Println(serveurAgt.GetBallots()[len(serveurAgt.GetBallots())-1])

	// Votants
	fmt.Println("")
	fmt.Println("Votants : ")
	agtNewVote16 := restclientagent.NewVoterAgent("idv16", url2, "ag_id16", "vote4", prefs1, nil)
	agtNewVote16.StartVoter()

	agtNewVote17 := restclientagent.NewVoterAgent("idv17", url2, "ag_id17", "vote4", prefs2, nil)
	agtNewVote17.StartVoter()

	agtNewVote18 := restclientagent.NewVoterAgent("idv18", url2, "ag_id18", "vote4", prefs3, nil)
	agtNewVote18.StartVoter()

	agtNewVote19 := restclientagent.NewVoterAgent("idv19", url2, "ag_id19", "vote4", prefs4, nil)
	agtNewVote19.StartVoter()

	agtNewVote20 := restclientagent.NewVoterAgent("idv20", url2, "ag_id20", "vote4", prefs5, nil)
	agtNewVote20.StartVoter()

	for i := 15; i < 20; i++ {
		fmt.Println(serveurAgt.GetVoters()[i])
	}

	// Result
	fmt.Println("")
	fmt.Println("Résultat : ")
	agResult = restclientagent.NewResultAgent("idr4", url2, "vote4")
	agResult.StartResultAgent()

	// -------------- TEST NUMERO 5 : 5 VOTANTS, Copeland -------------- //
	fmt.Println("")
	fmt.Println("")
	fmt.Println("------------ COPELAND ------------")

	// Ballot
	fmt.Println("")
	fmt.Println("Ballot : ")
	voter_ids = []string{"ag_id21", "ag_id22", "ag_id23", "ag_id24", "ag_id25"}
	agNewBallot = restclientagent.NewBallotAgent("idb5", url2, "copeland", "Wed Nov 30 23:00:00 UTC 2022", voter_ids, 5)
	agNewBallot.StartBallotAgent()

	fmt.Println(serveurAgt.GetBallots()[len(serveurAgt.GetBallots())-1])

	// Votants
	fmt.Println("")
	fmt.Println("Votants : ")
	agtNewVote21 := restclientagent.NewVoterAgent("idv21", url2, "ag_id21", "vote5", prefs1, nil)
	agtNewVote21.StartVoter()

	agtNewVote22 := restclientagent.NewVoterAgent("idv22", url2, "ag_id22", "vote5", prefs2, nil)
	agtNewVote22.StartVoter()

	agtNewVote23 := restclientagent.NewVoterAgent("idv23", url2, "ag_id23", "vote5", prefs3, nil)
	agtNewVote23.StartVoter()

	agtNewVote24 := restclientagent.NewVoterAgent("idv24", url2, "ag_id24", "vote5", prefs4, nil)
	agtNewVote24.StartVoter()

	agtNewVote25 := restclientagent.NewVoterAgent("idv25", url2, "ag_id25", "vote5", prefs5, nil)
	agtNewVote25.StartVoter()

	for i := 20; i < 25; i++ {
		fmt.Println(serveurAgt.GetVoters()[i])
	}

	// Result
	fmt.Println("")
	fmt.Println("Résultat : ")
	agResult = restclientagent.NewResultAgent("idr5", url2, "vote5")
	agResult.StartResultAgent()

	// -------------- TEST NUMERO 6 : 5 VOTANTS, STV -------------- //
	fmt.Println("")
	fmt.Println("")
	fmt.Println("------------ STV ------------")

	// Ballot
	voter_ids = []string{"ag_id26", "ag_id27", "ag_id28", "ag_id29", "ag_id30"}
	agNewBallot = restclientagent.NewBallotAgent("idb6", url2, "stv", "Wed Nov 30 23:00:00 UTC 2022", voter_ids, 6)
	agNewBallot.StartBallotAgent()

	fmt.Println("")
	fmt.Println("Ballot : ")
	fmt.Println(serveurAgt.GetBallots()[len(serveurAgt.GetBallots())-1])

	// Votants
	fmt.Println("")
	fmt.Println("Votants : ")
	agtNewVote26 := restclientagent.NewVoterAgent("idv26", url2, "ag_id16", "vote6", prefs1, nil)
	agtNewVote26.StartVoter()

	agtNewVote27 := restclientagent.NewVoterAgent("idv27", url2, "ag_id17", "vote6", prefs2, nil)
	agtNewVote27.StartVoter()

	agtNewVote28 := restclientagent.NewVoterAgent("idv28", url2, "ag_id18", "vote6", prefs3, nil)
	agtNewVote28.StartVoter()

	agtNewVote29 := restclientagent.NewVoterAgent("idv29", url2, "ag_id19", "vote6", prefs4, nil)
	agtNewVote29.StartVoter()

	agtNewVote30 := restclientagent.NewVoterAgent("idv30", url2, "ag_id20", "vote6", prefs5, nil)
	agtNewVote30.StartVoter()

	for i := 25; i < 30; i++ {
		fmt.Println(serveurAgt.GetVoters()[i])
	}

	// Result
	fmt.Println("")
	fmt.Println("Résultat : ")
	agResult = restclientagent.NewResultAgent("idr6", url2, "vote6")
	agResult.StartResultAgent()

	fmt.Scanln()
}
