package restclientagent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	rad "gitlab.utc.fr/aguilber/ia04/agt"
	"gitlab.utc.fr/aguilber/ia04/comsoc"
)

// En debug, pour afficher plus d'infos, mettre à true :
const IsDebugging bool = false

type BallotAgent struct {
	id       string
	url      string
	rule     string
	deadline string
	voterIds []string
	alts     int
}

func NewBallotAgent(id string, url string, rule string, deadline string, voterIds []string, alts int) *BallotAgent {
	if IsDebugging {
		log.Printf("Ballot Agent created, id  ='%s'", id)
	}
	return &BallotAgent{id, url, rule, deadline, voterIds, alts}
}

func (rca *BallotAgent) doRequestNewBallot() (resBallotId string, err error) {
	req := rad.Ballot{
		Rule:     rca.rule,
		Deadline: rca.deadline,
		VoterIds: rca.voterIds,
		Alts:     rca.alts,
	}

	// sérialisation de la requête
	url := rca.url + "/new_ballot"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusCreated {
		log.Println("Erreur status code")
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}
	resBallotId = rca.treatResponse(resp)

	return
}

func (rca *BallotAgent) treatResponse(r *http.Response) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	var resp rad.ResponseNewBallot
	json.Unmarshal(buf.Bytes(), &resp)

	return resp.BallotId
}

func (rca *BallotAgent) StartBallotAgent() {
	if IsDebugging {
		log.Printf("démarrage de %s", rca.id)
	}
	res, err := rca.doRequestNewBallot()
	if err != nil {
		log.Fatal(rca.id, " Error:", err.Error())
	} else {
		if IsDebugging {
			log.Printf("[%s] %s\n", rca.id, res)
		}
	}
}

// ------------ VOTE ------------ //

type VoterAgent struct {
	id      string
	url     string
	agentId string
	voteId  string
	prefs   []comsoc.Alternative
	options []int
}

func NewVoterAgent(id string, url string, agentId string, voteId string, prefs []comsoc.Alternative, options []int) *VoterAgent {
	if IsDebugging {
		log.Printf("Voter Agent created, id  ='%s'", agentId)
	}
	return &VoterAgent{id, url, agentId, voteId, prefs, options}
}

func (rca *VoterAgent) doRequestVoter() (err error) {
	req := rad.Voter{
		AgentId: rca.agentId,
		VoteId:  rca.voteId,
		Prefs:   rca.prefs,
		Options: rca.options,
	}

	// sérialisation de la requête
	url := rca.url + "/vote"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}

	return
}

func (rca *VoterAgent) StartVoter() {
	if IsDebugging {
		log.Printf("démarrage de %s", rca.id)
	}
	err := rca.doRequestVoter()

	if err != nil {
		log.Fatal(rca.id, "error:", err.Error())
	} else {
		if IsDebugging {
			log.Printf("[%s] %s %s %s, prefs = %d\n", rca.id, rca.url, rca.agentId, rca.voteId, rca.prefs)
		}
	}
}

// ------------ RESULT ------------ //

type ResultAgent struct {
	id        string
	url       string
	ballot_id string
}

func NewResultAgent(id string, url string, ballot_id string) *ResultAgent {
	if IsDebugging {
		log.Printf("Result Agent created, id  ='%s'", id)
	}
	return &ResultAgent{id, url, ballot_id}
}

func (rca *ResultAgent) doRequestResult() (resWinner comsoc.Alternative, resRanking []comsoc.Alternative, err error) {
	req := rad.Result{
		BallotId: rca.ballot_id,
	}

	// sérialisation de la requête
	url := rca.url + "/result"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("erreur status code")
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}
	resWinner, resRanking = rca.treatResponse(resp)

	return
}

func (rca *ResultAgent) treatResponse(r *http.Response) (resWinner comsoc.Alternative, resRanking []comsoc.Alternative) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	var resp rad.ResponseResult
	json.Unmarshal(buf.Bytes(), &resp)

	return resp.Winner, resp.Ranking
}

func (rca *ResultAgent) StartResultAgent() {
	if IsDebugging {
		log.Printf("démarrage de %s", rca.id)
	}
	resWinner, resRanking, err := rca.doRequestResult()
	if err != nil {
		log.Fatal(rca.id, " Error:", err.Error())
	} else {
		log.Printf("[%s] Winner : %d\n ", rca.id, resWinner)
		log.Printf("[%s] Ranking : \n", rca.id)
		if resRanking != nil {
			for _, r := range resRanking {
				log.Printf("%d", r)
			}
		} else {
			fmt.Println("Condorcet Winner => No Ranking")
		}
	}
}
