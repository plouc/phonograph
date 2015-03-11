package api

import (
	"fmt"
	"log"
	"github.com/jmcvetta/neoism"
)

type Master struct {
	node     *neoism.Node

	Id       int         `json:"id"`
	Name     string      `json:"name"`

	Links    Links       `json:"_links"`
	Embedded interface{} `json:"_embedded"`
}

func MasterFromNode(node *neoism.Node) *Master {
	name := node.Data["name"].(string)

	return &Master{
		node: node,
		Name: name,
	}
}

func (m *Master) AddRelease(release *Release) *Master {
	m.node.Relate("HAS_RELEASE", release.Id, nil)

	return m
}

func (m *Master) halify() {
	m.Links.Self = fmt.Sprintf("http://localhost:2000/masters/%d", m.Id)
}


type Masters []*Master

type MastersManager struct {
	db *neoism.Database
}

func NewMastersManager(db *neoism.Database) *MastersManager {
	return &MastersManager{
		db: db,
	}
}

func (mm *MastersManager) Create(masterName string) *Master {
	node, err := mm.db.CreateNode(neoism.Props{"name": masterName})
	if err != nil {
		log.Fatal(err)
	}

	node.AddLabel("Master")

	master := &Master{
		node: node,
		Id:   node.Id(),
		Name: masterName,
	}

	return master
}


func (mm *MastersManager) Find() Masters {
	results := []struct {
		M        neoism.Node
		MasterId int
	}{}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (m:Master)
			RETURN m, id(m) AS masterId
			ORDER BY m.name
		`,
		Result: &results,
	}

	mm.db.Cypher(&cq)

	masters := Masters{}

	for _, res := range results {
		master := MasterFromNode(&res.M)
		master.Id = res.MasterId
		master.halify()

		masters = append(masters, master)
	}

	return masters
}


func (mm *MastersManager) PlayedBy(artist *Artist) Masters {
	results := []struct {
		M        neoism.Node
		MasterId int
	}{}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (a:Artist)
			WHERE id(a) = {artistId}
			MATCH (a)-[PLAYED_IN]->(m:Master)
			RETURN m, id(m) AS masterId
			ORDER BY m.name
		`,
		Parameters: neoism.Props{"artistId": artist.Id},
		Result: &results,
	}

	mm.db.Cypher(&cq)

	masters := Masters{}

	for _, res := range results {
		master := MasterFromNode(&res.M)
		master.Id = res.MasterId
		master.halify()

		masters = append(masters, master)
	}

	return masters
}
