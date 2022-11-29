package restserveuragent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	rad "gitlab.utc.fr/aguilber/ia04/agt"
	client "gitlab.utc.fr/aguilber/ia04/agt/restclientagent"

	"gitlab.utc.fr/aguilber/ia04/comsoc"
)

// ------------ SERVEUR ------------ //

type RestServeurAgent struct {
	sync.Mutex
	id       string
	reqCount int
	addr     string
	ballots  []rad.Ballot
	voters   []rad.Voter
}

func NewRestServeurAgent(addr string) *RestServeurAgent {
	return &RestServeurAgent{id: addr, addr: addr}
}

func (rsa *RestServeurAgent) Start() {
	// création du multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/new_ballot", rsa.doNewBallot)
	mux.HandleFunc("/vote", rsa.doNewVote)
	mux.HandleFunc("/result", rsa.doResult)
	mux.HandleFunc("/reqcount", rsa.doReqcount)

	// création du serveur http
	s := &http.Server{
		Addr:           rsa.addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	// lancement du serveur
	if client.IsDebugging {
		log.Println("Listening on", rsa.addr)
	}
	go log.Fatal(s.ListenAndServe())
}

func (rsa *RestServeurAgent) doReqcount(w http.ResponseWriter, r *http.Request) {
	if !rsa.checkMethod("GET", w, r) {
		return
	}

	w.WriteHeader(http.StatusOK)
	rsa.Lock()
	defer rsa.Unlock()
	serial, _ := json.Marshal(rsa.reqCount)
	w.Write(serial)
}

func (rsa *RestServeurAgent) checkMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method %q not allowed", r.Method)
		return false
	}
	return true
}

// ------------ BALLOT ------------ //

func (rsa *RestServeurAgent) GetBallots() []rad.Ballot {
	return rsa.ballots
}

func (*RestServeurAgent) decodeRequestBallot(r *http.Request) (req rad.Ballot, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (rsa *RestServeurAgent) doNewBallot(w http.ResponseWriter, r *http.Request) {
	// mise à jour du nombre de requêtes
	rsa.Lock()
	defer rsa.Unlock()
	rsa.reqCount++

	// vérification de la méthode de la requête
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := rsa.decodeRequestBallot(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// traitement de la requête
	var resp rad.ResponseNewBallot

	// check Rule
	if req.Rule != "majority" && req.Rule != "borda" && req.Rule != "approval" && req.Rule != "stv" && req.Rule != "copeland" && req.Rule != "condorcet" {
		w.WriteHeader(http.StatusNotImplemented)
		msg := fmt.Sprintf("Unkonwn rule '%s'", req.Rule)
		log.Println(msg)
		w.Write([]byte(msg))
		return
	}

	// check Deadline
	if req.Deadline != "" {
		today := time.Now()
		deadlineTime, errParseTime := time.Parse(time.UnixDate, req.Deadline)

		if errParseTime != nil {
			w.WriteHeader(http.StatusNotImplemented)
			msg := fmt.Sprintf("The deadline layout %s is invalid", req.Deadline)
			log.Println(msg)
			w.Write([]byte(msg))
			return
		} else {
			if deadlineTime.Before(today) {
				w.WriteHeader(http.StatusNotImplemented)
				msg := fmt.Sprintf("The deadline %s was exceeded.", req.Deadline)
				log.Println(msg)
				w.Write([]byte(msg))
			}
		}
	} else {
		w.WriteHeader(http.StatusNotImplemented)
		msg := fmt.Sprintf("Unknown deadline '%s'", req.Rule)
		log.Println(msg)
		w.Write([]byte(msg))
		return
	}

	// check Voter Ids
	if req.VoterIds != nil {
		for _, id := range req.VoterIds { // pour chaque agent_id
			if !strings.HasPrefix(id, "ag_id") { // vérifie s'il commence par "ag_id"
				w.WriteHeader(http.StatusNotImplemented)
				msg := fmt.Sprintf("Invalid Agent Id '%s'", id)
				log.Println(msg)
				w.Write([]byte(msg))
				return
			}
		}
	}

	// check nombre d'Alts
	if req.Alts <= 0 {
		w.WriteHeader(http.StatusNotImplemented)
		msg := fmt.Sprintf("Invalid number of alternatives '%d'", req.Alts)
		log.Println(msg)
		w.Write([]byte(msg))
		return
	}

	resp.BallotId = fmt.Sprintf("vote%d", req.Alts)

	newBallot := rad.Ballot{
		Rule:     req.Rule,
		Deadline: req.Deadline,
		VoterIds: req.VoterIds,
		Alts:     req.Alts,
		BallotId: resp.BallotId,
	}
	rsa.ballots = append(rsa.ballots, newBallot)

	w.WriteHeader(http.StatusCreated)
	serial, _ := json.Marshal(resp)
	w.Write(serial)
}

// ------------ VOTE ------------ //

func (rsa *RestServeurAgent) GetVoters() []rad.Voter {
	return rsa.voters
}

func (*RestServeurAgent) decodeRequestVote(r *http.Request) (req rad.Voter, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (rsa *RestServeurAgent) doNewVote(w http.ResponseWriter, r *http.Request) {
	// mise à jour du nombre de requêtes
	rsa.Lock()
	defer rsa.Unlock()
	rsa.reqCount++

	// vérification de la méthode de la requête
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := rsa.decodeRequestVote(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// traitement de la requête

	// check vote-id dans ballot

	var found bool = false

	for i := range rsa.ballots {
		structure := rsa.ballots[i]
		if structure.BallotId == req.VoteId {
			found = true
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotImplemented)
		msg := fmt.Sprintf("Unkonwn ballot '%s'", req.VoteId)
		w.Write([]byte(msg))
		return
	}

	//check que agent id et vote id pas deja dans tableau
	//sinon vote deja effectue

	found = false

	for i := range rsa.voters {
		if rsa.voters[i].AgentId == req.AgentId && rsa.voters[i].VoteId == req.VoteId {
			found = true
		}
	}

	if found {
		w.WriteHeader(http.StatusForbidden) //403
		msg := fmt.Sprintf("Agent '%s' has already voted  for ballot '%s'", req.AgentId, req.VoteId)
		w.Write([]byte(msg))
		return
	}

	//check que la deadline est pas depassee
	//trouver la deadline
	var Deadline string

	for i := range rsa.ballots { //gerer les doublons
		structure := rsa.ballots[i]
		if structure.BallotId == req.VoteId {
			Deadline = structure.Deadline
		}
	}

	if Deadline != "" {
		today := time.Now()
		deadlineTime, errParseTime := time.Parse(time.UnixDate, Deadline)

		if errParseTime != nil {
			w.WriteHeader(http.StatusNotImplemented)
			msg := fmt.Sprintf("The deadline layout %s is invalid", Deadline)
			log.Println(msg)
			w.Write([]byte(msg))
			return
		} else {
			if deadlineTime.Before(today) {
				w.WriteHeader(http.StatusNotImplemented)
				msg := fmt.Sprintf("The deadline %s was exceeded.", Deadline)
				log.Println(msg)
				w.Write([]byte(msg))
			}
		}
	} else {
		w.WriteHeader(http.StatusNotImplemented)
		msg := fmt.Sprintf("Unkonwn deadline")
		log.Println(msg)
		w.Write([]byte(msg))
		return
	}

	//si tous les checks sont OK
	//on ajoute le votant dans rsa.voters

	newVoter := rad.Voter{
		AgentId: req.AgentId,
		VoteId:  req.VoteId,
		Prefs:   req.Prefs,
		Options: req.Options}

	rsa.voters = append(rsa.voters, newVoter)

	w.WriteHeader(http.StatusOK)

}

// ------------ RESULT ------------ //

func (*RestServeurAgent) decodeRequestResult(r *http.Request) (req rad.Result, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (rsa *RestServeurAgent) doResult(w http.ResponseWriter, r *http.Request) {
	// mise à jour du nombre de requêtes
	rsa.Lock()
	defer rsa.Unlock()
	rsa.reqCount++

	// vérification de la méthode de la requête
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := rsa.decodeRequestResult(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	var rule string
	// Check BallotId and Deadline
	for _, ballot := range rsa.ballots {
		if ballot.BallotId == req.BallotId {

			/*
			 Code qui vérifie si la deadline pour aller voter est dépassée.
			 Si c'est le cas, on peut lancer les dépouilles et donner le résultat
			 Si ce n'est pas le cas, on renvoie un erreur http.StatusTooEarly
			 => Fonctionalité testable uniquement si on lance le fichier serveur et
			 le fichier client séparément et que l'utilisateur puisse exécuter
			 les end-points en temps-réel.
			*/

			/*
			 today := time.Now()
			 deadlineTime, errParseTime := time.Parse(time.UnixDate, ballot.Deadline)
			 if errParseTime != nil {
			 	fmt.Println(errParseTime)
			 } else {
			 	if deadlineTime.After(today) {
			 		w.WriteHeader(http.StatusTooEarly)
			 		msg := fmt.Sprintf("The deadline %s is NOT YET exceeded.", ballot.Deadline)
			 		log.Println(msg)
			 		w.Write([]byte(msg))
			 	} else {
			*/

			rule = ballot.Rule

			// }
			// }
		}
	}

	// traitement de la requête
	var resp rad.ResponseResult

	// Create Profile and stock threshold if exists
	var profile comsoc.Profile
	var thresholds []int
	for _, voter := range rsa.voters {
		if voter.VoteId == req.BallotId {
			profile = append(profile, voter.Prefs)
			if voter.Options != nil {
				thresholds = append(thresholds, voter.Options[0])
			}
		}
	}

	switch rule {
	case "majority":
		// Ranking
		giveRanking := comsoc.SWFFactory(comsoc.MajoritySWF, comsoc.TieBreakFactory([]comsoc.Alternative{}))
		ranking, errRanking := giveRanking(profile)
		if errRanking == nil {
			resp.Ranking = ranking
		} else {
			fmt.Println(errRanking)
		}

		// Winner
		resSCF, errSCF := comsoc.MajoritySCF(profile)
		if errSCF == nil {
			tiebreak := comsoc.TieBreakFactory(resSCF)
			winner, errTieBreak := tiebreak(resSCF)
			if errTieBreak == nil {
				resp.Winner = winner
			} else {
				fmt.Println(errTieBreak)
			}
		} else {
			fmt.Println(errSCF)
		}
		break

	case "borda":
		// Ranking
		giveRanking := comsoc.SWFFactory(comsoc.BordaSWF, comsoc.TieBreakFactory([]comsoc.Alternative{}))
		ranking, errRanking := giveRanking(profile)
		if errRanking == nil {
			resp.Ranking = ranking
		} else {
			fmt.Println(errRanking)
		}

		// Winner
		resSCF, errSCF := comsoc.BordaSCF(profile)
		if errSCF == nil {
			tiebreak := comsoc.TieBreakFactory(resSCF)
			winner, errTieBreak := tiebreak(resSCF)
			if errTieBreak == nil {
				resp.Winner = winner
			} else {
				fmt.Println(errTieBreak)
			}
		} else {
			fmt.Println(errSCF)
		}
		break

	case "approval":
		// Ranking
		giveWinner := comsoc.SWFFactoryApproval(comsoc.ApprovalSWF, comsoc.TieBreakFactory([]comsoc.Alternative{}))
		ranking, errRanking := giveWinner(profile, thresholds)
		if errRanking == nil {
			resp.Ranking = ranking
		} else {
			fmt.Println(errRanking)
		}

		// Winner
		winner, err := comsoc.ApprovalSCF(profile, thresholds)
		if err == nil {
			resp.Winner = winner[0]
		} else {
			fmt.Println(err)
		}

		break
	case "stv":
		// Ranking
		giveRanking := comsoc.SWFFactory(comsoc.STV_SWF, comsoc.TieBreakFactory([]comsoc.Alternative{}))
		ranking, errRanking := giveRanking(profile)
		if errRanking == nil {
			resp.Ranking = ranking
		} else {
			fmt.Println(errRanking)
		}

		// Winner
		resSCF, errSCF := comsoc.STV_SCF(profile)
		if errSCF == nil {
			tiebreak := comsoc.TieBreakFactory(resSCF)
			winner, errTieBreak := tiebreak(resSCF)
			if errTieBreak == nil {
				resp.Winner = winner
			} else {
				fmt.Println(errTieBreak)
			}
		} else {
			fmt.Println(errSCF)
		}
		break
	case "copeland":
		// Ranking
		giveRanking := comsoc.SWFFactory(comsoc.CopelandSWF, comsoc.TieBreakFactory([]comsoc.Alternative{}))
		ranking, errRanking := giveRanking(profile)
		if errRanking == nil {
			resp.Ranking = ranking
		} else {
			fmt.Println(errRanking)
		}

		// Winner
		resSCF, errSCF := comsoc.CopelandSCF(profile)
		if errSCF == nil {
			tiebreak := comsoc.TieBreakFactory(resSCF)
			winner, errTieBreak := tiebreak(resSCF)
			if errTieBreak == nil {
				resp.Winner = winner
			} else {
				fmt.Println(errTieBreak)
			}
		} else {
			fmt.Println(errSCF)
		}
		break
	case "condorcet":
		// Ranking
		resp.Ranking = nil

		// Winner
		winner, err := comsoc.CondorcetWinner(profile)
		if err == nil {
			resp.Winner = winner[0]
		} else {
			fmt.Println(err)
		}
		break
	}

	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
