package api

import (
	"fmt"
	"log"
	"github.com/jmcvetta/neoism"
)

type Release struct {
	node *neoism.Node

	Id      int    `json:"id"`
	Year    int    `json:"year"`
	Country string `json:"country"`

	Links   Links  `json:"_links"`
}

func ReleaseFromNode(node *neoism.Node) *Release {
	year    := int(node.Data["year"].(float64))
	country := node.Data["country"].(string)

	return &Release{
		node:    node,
		Year:    year,
		Country: country,
	}
}

func (r *Release) AddLabel(label *Label) *Release {
	r.node.Relate("BY_LABEL", label.Id(), nil)

	return r
}

func (r *Release) halify() {
	r.Links.Self = fmt.Sprintf("http://localhost:2000/releases/%d", r.Id)
}

type Releases []Release

type ReleasesManager struct {
	db *neoism.Database
}

func NewReleasesManager (db *neoism.Database) *ReleasesManager {
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
		node:    node,
		Id:      node.Id(),
		Year:    year,
		Country: country,
	}

	return release
}
