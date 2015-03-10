package api

import (
	"fmt"
	"log"
	"github.com/jmcvetta/neoism"
)

type Skill struct {
	node     *neoism.Node

	Id       int         `json:"id"`
	Name     string      `json:"name"`

	Links    Links       `json:"_links"`
	Embedded interface{} `json:"_embedded"`
}

func SkillFromNode(node *neoism.Node) *Skill {
	name := node.Data["name"].(string)

	return &Skill{
		node: node,
		Name: name,
	}
}

func (s *Skill) halify() {
	s.Links.Self = fmt.Sprintf("http://localhost:2000/skills/%d", s.Id)
}

type Skills []Skill

type SkillsManager struct {
	db *neoism.Database
}

func NewSkillsManager (db *neoism.Database) *SkillsManager {
	return &SkillsManager{
		db: db,
	}
}

func (sm *SkillsManager) Create(skillName string) *Skill {
	node, err := sm.db.CreateNode(neoism.Props{"name": skillName})
	if err != nil {
		log.Fatal(err)
	}

	node.AddLabel("Skill")

	skill := &Skill{
		node: node,
		Id:   node.Id(),
		Name: skillName,
	}

	return skill
}

func (sm *SkillsManager) Find() Skills {
	res := Skills{}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (a:Artist)
			RETURN a.name AS Name, a.skills AS Skills
			ORDER BY a.name
		`,
		Result: &res,
	}

	sm.db.Cypher(&cq)

	return res
}
