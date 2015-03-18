package api

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"log"
)

type Style struct {
	ApiNode

	Name string `json:"name"`
}

func StyleFromNode(node *neoism.Node) *Style {
	name := node.Data["name"].(string)

	return &Style{
		ApiNode: ApiNode{node: node},
		Name:    name,
	}
}

func (s *Style) halify() {
	s.Links.Self = fmt.Sprintf("http://localhost:2000/style/%d", s.Id)
}

type Styles []Style

type StylesManager struct {
	db *neoism.Database
}

func NewStylesManager(db *neoism.Database) *StylesManager {
	return &StylesManager{
		db: db,
	}
}

func (sm *StylesManager) Create(styleName string) *Style {
	node, err := sm.db.CreateNode(neoism.Props{"name": styleName})
	if err != nil {
		log.Fatal(err)
	}

	node.AddLabel("Style")

	style := &Style{
		ApiNode: ApiNode{
			node: node,
			Id:   node.Id(),
		},
		Name: styleName,
	}

	return style
}
