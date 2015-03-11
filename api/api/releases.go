package api

import (
	//"fmt"
	"log"
	"github.com/jmcvetta/neoism"
)

type Release struct {
	node *neoism.Node

	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (r *Release) ProducedBy(label *Label) *Release {
	r.node.Relate("PRODUCED_BY", label.Id(), nil)

	return r
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

func (rm *ReleasesManager) Create(releaseName string) *Release {
	node, err := rm.db.CreateNode(neoism.Props{"name": releaseName})
	if err != nil {
		log.Fatal(err)
	}

	node.AddLabel("Release")

	release := &Release{
		node: node,
		Id:   node.Id(),
		Name: releaseName,
	}

	return release
}
