package api

import (
	"fmt"
	"log"
	"github.com/jmcvetta/neoism"
)

type Label struct {
	node        *neoism.Node
	Name        string      `json:"name"`
	Links       Links       `json:"_links"`
	Embedded    interface{} `json:"_embedded"`
}

func (l *Label) Id() int {
	return l.node.Id()
}

func (l *Label) Halify() {
	l.Links.Self = fmt.Sprintf("http://localhost:2000/labels/%d", 3)
	l.Embedded = map[string]string{
		"productions": "test",
	}
}

type Labels []Label

type LabelsManager struct {
	db *neoism.Database
}

func NewLabelsManager (db *neoism.Database) *LabelsManager {
	return &LabelsManager{
		db: db,
	}
}

func (lm *LabelsManager) Create(labelName string) *Label {
	node, err := lm.db.CreateNode(neoism.Props{"name": labelName})
	if err != nil {
		log.Fatal(err)
	}

	node.AddLabel("Label")

	return &Label{
		node: node,
		Name: labelName,
	}
}

func (lm *LabelsManager) Find() Labels {
	labels := Labels{}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (l:Label)
			RETURN l.name AS Name
			ORDER BY l.name
		`,
		Result: &labels,
	}

	lm.db.Cypher(&cq)

	for k, _ := range labels {
		labels[k].Halify()
	}

	return labels
}
