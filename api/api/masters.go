package api

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"log"
)

type Master struct {
	ApiNode

	Name     string     `json:"name"`
	Year     int        `json:"year"`
	Releases []*Release `json:"releases"`
	Artists  []*Artist  `json:"artists"`
	Tracks   []*Track   `json:"tracks"`
}

type Masters []*Master

type MastersCollection struct {
	ApiCollection
	Results *Masters `json:"results"`
}

func (mc *MastersCollection) halify() {
	mc.Links.Self = fmt.Sprintf("http://localhost:2000/masters?page=%d", mc.Pager.Page)
	mc.Links.Prev = fmt.Sprintf("http://localhost:2000/masters?page=%d", mc.Pager.Page-1)
	mc.Links.Next = fmt.Sprintf("http://localhost:2000/masters?page=%d", mc.Pager.Page+1)
}

type MastersManager struct {
	db *neoism.Database
}

func MasterFromNode(node *neoism.Node) *Master {
	name := node.Data["name"].(string)
	year := int(node.Data["year"].(float64))

	return &Master{
		ApiNode:  ApiNode{node: node},
		Name:     name,
		Year:     year,
		Releases: []*Release{},
		Artists:  []*Artist{},
		Tracks:   []*Track{},
	}
}

func (m *Master) AddTrack(track *Track) *Master {
	m.node.Relate("HAS_TRACK", track.Id, nil)

	return m
}

func (m *Master) AddRelease(release *Release) *Master {
	m.node.Relate("HAS_RELEASE", release.Id, nil)

	return m
}

func (m *Master) AddStyle(style *Style) *Master {
	m.node.Relate("CLASSIFIED_IN", style.Id, nil)

	return m
}

func (m *Master) halify() {
	m.Links.Self = fmt.Sprintf("http://localhost:2000/masters/%d", m.Id)
}

func NewMastersManager(db *neoism.Database) *MastersManager {
	return &MastersManager{
		db: db,
	}
}

func (mm *MastersManager) Create(masterName string, year int) *Master {
	node, err := mm.db.CreateNode(neoism.Props{
		"name": masterName,
		"year": year,
	})
	if err != nil {
		log.Fatal(err)
	}

	node.AddLabel("Master")

	master := &Master{
		ApiNode: ApiNode{
			node: node,
			Id:   node.Id(),
		},
		Name:     masterName,
		Year:     year,
		Releases: []*Release{},
	}

	return master
}

func (mm *MastersManager) Find(pager *Pager) *MastersCollection {
	results := []struct {
		M        neoism.Node
		MasterId int
	}{}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (m:Master)
			RETURN m, id(m) AS masterId
			ORDER BY m.name
		`,
		Result: &results,
	}

	mm.db.Cypher(&cq)

	masters := Masters{}

	for _, res := range results {
		master := MasterFromNode(&res.M)
		master.Id = res.MasterId
		master.halify()

		masters = append(masters, master)
	}

	collection := MastersCollection{
		ApiCollection: ApiCollection{Pager: pager},
		Results:       &masters,
	}

	collection.halify()

	return &collection
}

func (mm *MastersManager) FindById(id int) (*Master, error) {
	results := []struct {
		M       neoism.Node
		RelType string
		NodeId  int
		N       neoism.Node
		LabelId int
		L       neoism.Node
	}{}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (m:Master)
			WHERE id(m) = {nodeId}
			OPTIONAL MATCH (m)-[r:HAS_RELEASE|PLAYED_IN|HAS_TRACK]-(n)
			OPTIONAL MATCH (n:Release)-[BY_LABEL]->(l:Label)
			RETURN m, type(r) AS relType, id(n) AS nodeId, n, id(l) AS labelId, l
			ORDER BY type(r), n.name
		`,
		Parameters: neoism.Props{"nodeId": id},
		Result:     &results,
	}

	err := mm.db.Cypher(&cq)
	if err != nil {
		log.Fatal(err)
	}

	if len(results) == 0 {
		return nil, NotFound
	}

	master := MasterFromNode(&results[0].M)
	master.Id = id

	nodesById := make(map[int]interface{})

	for _, res := range results {
		if res.NodeId != 0 {
			if res.RelType == "HAS_RELEASE" {
				_, ok := nodesById[res.NodeId]
				if !ok {
					rel := ReleaseFromNode(&res.N)
					nodesById[res.NodeId] = rel
					rel.Id = res.NodeId
					rel.halify()
					master.Releases = append(master.Releases, rel)
				}
				if res.LabelId != 0 {
					lab := LabelFromNode(&res.L)
					lab.Id = res.LabelId
					lab.halify()
					rel := nodesById[res.NodeId].(*Release)
					rel.Labels = append(rel.Labels, lab)
				}
			}
			if res.RelType == "PLAYED_IN" {
				art := ArtistFromNode(&res.N)
				nodesById[res.NodeId] = art
				art.Id = res.NodeId
				art.halify()
				master.Artists = append(master.Artists, art)
			}
			if res.RelType == "HAS_TRACK" {
				tra := TrackFromNode(&res.N)
				nodesById[res.NodeId] = tra
				tra.Id = res.NodeId
				tra.halify()
				master.Tracks = append(master.Tracks, tra)
			}
		}
	}

	master.halify()

	return master, nil
}

func (mm *MastersManager) PlayedBy(artist *Artist) Masters {
	results := []struct {
		M        neoism.Node
		MasterId int
	}{}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (a:Artist)
			WHERE id(a) = {artistId}
			MATCH (a)-[r:MEMBER_OF|PLAYED_IN*]->(m:Master)
			RETURN m, id(m) AS masterId
			ORDER BY m.name
		`,
		Parameters: neoism.Props{"artistId": artist.Id},
		Result:     &results,
	}

	mm.db.Cypher(&cq)

	masters := Masters{}

	for _, res := range results {
		master := MasterFromNode(&res.M)
		master.Id = res.MasterId
		master.halify()

		masters = append(masters, master)
	}

	return masters
}
