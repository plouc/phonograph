package api

import (
	"fmt"
	"log"
	"github.com/jmcvetta/neoism"
)

type Track struct {
	node     *neoism.Node

	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Duration int         `json:"duration"`

	Links    Links       `json:"_links"`
	Embedded interface{} `json:"_embedded"`
}

type Tracks []*Track

type TracksCollection struct {
	HalCollection
	Pager   *Pager  `json:"pager"`
	Results *Tracks `json:"results"`
}

func (mc *TracksCollection) halify() {
	mc.Links.Self = fmt.Sprintf("http://localhost:2000/tracks?page=%d", mc.Pager.Page)
	mc.Links.Prev = fmt.Sprintf("http://localhost:2000/tracks?page=%d", mc.Pager.Page - 1)
	mc.Links.Next = fmt.Sprintf("http://localhost:2000/tracks?page=%d", mc.Pager.Page + 1)
}

type TracksManager struct {
	db *neoism.Database
}

func TrackFromNode(node *neoism.Node) *Track {
	name     := node.Data["name"].(string)
	duration := int(node.Data["duration"].(float64))

	return &Track{
		node:     node,
		Name:     name,
		Duration: duration,
	}
}

func (t *Track) halify() {
	t.Links.Self = fmt.Sprintf("http://localhost:2000/track/%d", t.Id)
}

func NewTracksManager(db *neoism.Database) *TracksManager {
	return &TracksManager{
		db: db,
	}
}

func (tm *TracksManager) Create(trackName string, trackDuration int) *Track {
	node, err := tm.db.CreateNode(neoism.Props{
		"name":     trackName,
		"duration": trackDuration,
	})
	if err != nil {
		log.Fatal(err)
	}

	node.AddLabel("Track")

	track := &Track{
		node:     node,
		Id:       node.Id(),
		Name:     trackName,
		Duration: trackDuration,
	}

	return track
}


func (tm *TracksManager) Find(pager *Pager) *TracksCollection {
	results := []struct {
		T       neoism.Node
		TrackId int
	}{}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (t:Track)
			RETURN t, id(t) AS trackId
			ORDER BY t.name
		`,
		Result: &results,
	}

	tm.db.Cypher(&cq)

	tracks := Tracks{}

	for _, res := range results {
		track := TrackFromNode(&res.T)
		track.Id = res.TrackId
		track.halify()

		tracks = append(tracks, track)
	}

	collection := TracksCollection{
		Pager:   pager,
		Results: &tracks,
	}

	collection.halify()

	return &collection
}
