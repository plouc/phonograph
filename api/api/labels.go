package api

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"log"
)

type Label struct {
	ApiNode

	Name string `json:"name"`
}

func LabelFromNode(node *neoism.Node) *Label {
	name := node.Data["name"].(string)

	return &Label{
		ApiNode: ApiNode{node: node},
		Name:    name,
	}
}

func (l *Label) halify() {
	l.Links.Self = fmt.Sprintf("http://localhost:2000/labels/%d", l.Id)
}

type Labels []Label

type LabelsManager struct {
	db *neoism.Database
}

func NewLabelsManager(db *neoism.Database) *LabelsManager {
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
		ApiNode: ApiNode{
			node: node,
			Id:   node.Id(),
		},
		Name:    labelName,
	}
}

func (lm *LabelsManager) FindById(id int) (*Label, error) {
	results := []struct {
		L neoism.Node
	}{}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (l:Label)
			WHERE id(l) = {nodeId}
			RETURN l
		`,
		Parameters: neoism.Props{"nodeId": id},
		Result:     &results,
	}

	err := lm.db.Cypher(&cq)
	if err != nil {
		log.Fatal(err)
	}

	if len(results) == 0 {
		return nil, NotFound
	}

	label := LabelFromNode(&results[0].L)
	label.Id = id

	label.halify()

	return label, nil
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
		labels[k].halify()
	}

	return labels
}
