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

type StylesCollection struct {
	ApiCollection
	Results *Styles `json:"results"`
}

func (sc *StylesCollection) halify() {
	sc.Links.Self = fmt.Sprintf("http://localhost:2000/styles?page=%d", sc.Pager.Page)
	sc.Links.Prev = fmt.Sprintf("http://localhost:2000/styles?page=%d", sc.Pager.Page-1)
	sc.Links.Next = fmt.Sprintf("http://localhost:2000/styles?page=%d", sc.Pager.Page+1)
}

type Styles []*Style

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

func (sm *StylesManager) FindById(id int) (*Style, error) {
	results := []struct {
		S neoism.Node
	}{}

	cq := neoism.CypherQuery{
		Statement: `
MATCH (s:Style)
WHERE id(s) = {nodeId}
RETURN s`,
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

	style := StyleFromNode(&results[0].S)
	style.Id = id
	style.halify()

	return style, nil
}

func (sm *StylesManager) Find(pager *Pager) *StylesCollection {
	results := []struct {
		S       neoism.Node
		StyleId int
		RelType string
		N       neoism.Node
		NodeId  int
		Total   int
	}{}

	cq := neoism.CypherQuery{
		Statement: `
MATCH (s:Style), (t:Style)
WITH s, count(t) AS total
ORDER BY s.name
SKIP {offset} LIMIT {limit}
RETURN s, id(s) AS styleId, total`,
		Parameters: neoism.Props{
			"offset": pager.Offset(),
			"limit":  pager.PerPage,
		},
		Result: &results,
	}

	sm.db.Cypher(&cq)

	styles := Styles{}

	if len(results) > 0 {
		pager.SetTotal(results[0].Total)
		for _, res := range results {
			style := StyleFromNode(&res.S)
			style.Id = res.StyleId
			style.halify()

			styles = append(styles, style)
		}
	}

	collection := StylesCollection{
		ApiCollection: ApiCollection{Pager: pager},
		Results:       &styles,
	}

	collection.halify()

	return &collection
}
