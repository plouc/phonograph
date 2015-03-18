package api

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"log"
)

type Skill struct {
	ApiNode

	Name string `json:"name"`
}

func SkillFromNode(node *neoism.Node) *Skill {
	name := node.Data["name"].(string)

	return &Skill{
		ApiNode: ApiNode{node: node},
		Name:    name,
	}
}

func (s *Skill) halify() {
	s.Links.Self = fmt.Sprintf("http://localhost:2000/skills/%d", s.Id)
}

type Skills []Skill

type SkillsManager struct {
	db *neoism.Database
}

func NewSkillsManager(db *neoism.Database) *SkillsManager {
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
		ApiNode: ApiNode{
			node: node,
			Id:   node.Id(),
		},
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
