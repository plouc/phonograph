package api

import (
	"fmt"
	"log"
	"github.com/jmcvetta/neoism"
)

type Style struct {
	node     *neoism.Node

	Id       int         `json:"id"`
	Name     string      `json:"name"`

	Links    Links       `json:"_links"`
	Embedded interface{} `json:"_embedded"`
}

func StyleFromNode(node *neoism.Node) *Style {
	name := node.Data["name"].(string)

	return &Style{
		node: node,
		Name: name,
	}
}

func (s *Style) halify() {
	s.Links.Self = fmt.Sprintf("http://localhost:2000/style/%d", s.Id)
}

type Styles []Style

type StylesManager struct {
	db *neoism.Database
}

func NewStylesManager (db *neoism.Database) *StylesManager {
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
		node: node,
		Id:   node.Id(),
		Name: styleName,
	}

	return style
}
