package agt

import "gitlab.utc.fr/aguilber/ia04/comsoc"

type Alternative int

type AgentID string

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []Alternative
}

type AgentI interface {
	Equal(ag AgentI) bool
	DeepEqual(ag AgentI) bool
	Clone() AgentI
	String() string
	Prefers(a Alternative, b Alternative)
	Start()
}

type Ballot struct {
	Rule     string   `json:"rule"`
	Deadline string   `json:"deadline"`
	VoterIds []string `json:"voter-ids"`
	Alts     int      `json:"#alts"`
	BallotId string   `json:"ballot-id"`
}

type ResponseNewBallot struct {
	BallotId string `json:"ballot-id"`
}

type Profile [][]int

type Voter struct {
	AgentId string               `json:"agentId"`
	VoteId  string               `json:"voteId"`
	Prefs   []comsoc.Alternative `json:"prefs"`
	Options []int                `json:"options"`
}

type Result struct {
	BallotId string `json:"ballot-id"`
}

type ResponseResult struct {
	Winner  comsoc.Alternative   `json:"winner"`
	Ranking []comsoc.Alternative `json:"ranking"`
}
