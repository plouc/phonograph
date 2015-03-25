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

type SkillsCollection struct {
	ApiCollection
	Results *Skills `json:"results"`
}

func (sc *SkillsCollection) halify() {
	sc.Links.Self = fmt.Sprintf("http://localhost:2000/skills?page=%d", sc.Pager.Page)
	sc.Links.Prev = fmt.Sprintf("http://localhost:2000/skills?page=%d", sc.Pager.Page-1)
	sc.Links.Next = fmt.Sprintf("http://localhost:2000/skills?page=%d", sc.Pager.Page+1)
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

type Skills []*Skill

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

func (sm *SkillsManager) FindById(id int) (*Skill, error) {
	results := []struct {
		S neoism.Node
	}{}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (s:Skill)
			WHERE id(s) = {nodeId}
			RETURN s
		`,
		Parameters: neoism.Props{"nodeId": id},
		Result:     &results,
	}

	err := sm.db.Cypher(&cq)
	if err != nil {
		log.Fatal(err)
	}

	if len(results) == 0 {
		return nil, NotFound
	}

	skill := SkillFromNode(&results[0].S)
	skill.Id = id
	skill.halify()

	return skill, nil
}


func (sm *SkillsManager) Find(pager *Pager) *SkillsCollection {
	results := []struct {
		S       neoism.Node
		SkillId int
		RelType string
		N       neoism.Node
		NodeId  int
		Total   int
	}{}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (s:Skill), (t:Skill)
			WITH s, count(t) AS total
			ORDER BY s.name
			SKIP {offset} LIMIT {limit}
			RETURN s, id(s) AS skillId, total
		`,
		Parameters: neoism.Props{
			"offset": pager.Offset(),
			"limit":  pager.PerPage,
		},
		Result: &results,
	}

	sm.db.Cypher(&cq)

	skills := Skills{}

	if len(results) > 0 {
		pager.SetTotal(results[0].Total)
		for _, res := range results {
			skill := SkillFromNode(&res.S)
			skill.Id = res.SkillId
			skill.halify()

			skills = append(skills, skill)
		}
	}

	collection := SkillsCollection{
		ApiCollection: ApiCollection{Pager: pager},
		Results:       &skills,
	}

	collection.halify()

	return &collection
}
