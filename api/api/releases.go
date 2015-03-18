package api

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"log"
)

type Release struct {
	ApiNode

	Year    int      `json:"year"`
	Country string   `json:"country"`
	Labels  []*Label `json:"labels"`
}

func ReleaseFromNode(node *neoism.Node) *Release {
	year := int(node.Data["year"].(float64))
	country := node.Data["country"].(string)

	return &Release{
		ApiNode: ApiNode{node: node},
		Year:    year,
		Country: country,
		Labels:  []*Label{},
	}
}

func (r *Release) AddLabel(label *Label) *Release {
	r.node.Relate("BY_LABEL", label.Id, nil)

	return r
}

func (r *Release) halify() {
	r.Links.Self = fmt.Sprintf("http://localhost:2000/releases/%d", r.Id)
}

type Releases []Release

type ReleasesManager struct {
	db *neoism.Database
}

func NewReleasesManager(db *neoism.Database) *ReleasesManager {
	return &ReleasesManager{
		db: db,
	}
}

func (rm *ReleasesManager) Create(year int, country string) *Release {
	node, err := rm.db.CreateNode(neoism.Props{
		"year":    year,
		"country": country,
	})
	if err != nil {
		log.Fatal(err)
	}

	node.AddLabel("Release")

	release := &Release{
		ApiNode: ApiNode{
			node: node,
			Id:   node.Id(),
		},
		Year:    year,
		Country: country,
	}

	return release
}
