package main

import (
	cache "github.com/patrickmn/go-cache"
)

// Candidate Model
type Candidate struct {
	ID             int
	Name           string
	PoliticalParty string
	Sex            string
}

// CandidateElectionResult type
type CandidateElectionResult struct {
	ID             int
	Name           string
	PoliticalParty string
	Sex            string
	VoteCount      int
}

// PartyElectionResult type
type PartyElectionResult struct {
	PoliticalParty string
	VoteCount      int
}

func getAllCandidate() (candidates []Candidate) {
	rows, err := db.Query("SELECT * FROM candidates")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		c := Candidate{}
		err = rows.Scan(&c.ID, &c.Name, &c.PoliticalParty, &c.Sex)
		if err != nil {
			panic(err.Error())
		}
		candidates = append(candidates, c)
	}
	return
}

func getCandidate(candidateID int) (c Candidate, err error) {
	row := db.QueryRow("SELECT * FROM candidates WHERE id = ?", candidateID)
	err = row.Scan(&c.ID, &c.Name, &c.PoliticalParty, &c.Sex)
	return
}

func getCandidateByName(name string) (c Candidate, err error) {
	row := db.QueryRow("SELECT * FROM candidates WHERE name = ?", name)
	err = row.Scan(&c.ID, &c.Name, &c.PoliticalParty, &c.Sex)
	return
}

func getAllPartyName() (partyNames []string) {
	rows, err := db.Query("SELECT political_party FROM candidates GROUP BY political_party")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			panic(err.Error())
		}
		partyNames = append(partyNames, name)
	}
	return
}

func getCandidatesIdsByPoliticalParty(party string) (ids []int) {
	rows, err := db.Query("SELECT id FROM candidates WHERE political_party = ?", party)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		c := Candidate{}
		err = rows.Scan(&c.ID)
		if err != nil {
			panic(err.Error())
		}
		ids = append(ids, c.ID)
	}
	return
}

func getCandidatesByPoliticalParty(party string) (candidates []Candidate) {
	rows, err := db.Query("SELECT * FROM candidates WHERE political_party = ?", party)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		c := Candidate{}
		err = rows.Scan(&c.ID, &c.Name, &c.PoliticalParty, &c.Sex)
		if err != nil {
			panic(err.Error())
		}
		candidates = append(candidates, c)
	}
	return
}

func getElectionResult() []CandidateElectionResult {
	cachedElection, found := cacheClient.Get("electionResult")

	if found {
		return cachedElection.([]CandidateElectionResult)
	}

	rows, err := db.Query(`
		SELECT c.id, c.name, c.political_party, c.sex, IFNULL(v.count, 0)
		FROM candidates AS c
		LEFT OUTER JOIN
	  	(SELECT candidate_id, COUNT(*) AS count
	  	FROM votes
	  	GROUP BY candidate_id) AS v
		ON c.id = v.candidate_id
		ORDER BY v.count DESC`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []CandidateElectionResult

	for rows.Next() {
		r := CandidateElectionResult{}
		err = rows.Scan(&r.ID, &r.Name, &r.PoliticalParty, &r.Sex, &r.VoteCount)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, r)
	}

	cacheClient.Set("electionResult", result, cache.DefaultExpiration)
	return result
}

func getElectionCountWithName(name string) int {
	count := 0

	rows := db.QueryRow(`
   select count(1)
   from votes
   where candidate_id in (
     select id from candidates where political_party = ?
   );
  `, name)

	rows.Scan(&count)

	return count
}
