package api

import (
	"fmt"
	"log"
	"github.com/jmcvetta/neoism"
)

type Artist struct {
	node     *neoism.Node

	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Skills   []*Skill    `json:"skills"`
	Groups   []*Artist   `json:"groups"`

	Links    Links       `json:"_links"`
	Embedded interface{} `json:"_embedded"`
}

type Artists []*Artist

func ArtistFromNode(node *neoism.Node) *Artist {
	name := node.Data["name"].(string)

	return &Artist{
		node:   node,
		Name:   name,
		Skills: []*Skill{},
		Groups: []*Artist{},
	}
}

func (a *Artist) hasSkill(skill *Skill) *Artist {
	a.Skills = append(a.Skills, skill)

	return a
}

func (a *Artist) AddSkill(skill *Skill) *Artist {
	a.node.Relate("HAS_SKILL", skill.Id, nil)

	return a
}

func (a *Artist) PlayedIn(master *Master) *Artist {
	a.node.Relate("PLAYED_IN", master.Id, nil)

	return a
}

func (a *Artist) AddMembership(artist *Artist) *Artist {
	a.node.Relate("MEMBER_OF", artist.Id, nil)

	return a
}

func (a *Artist) memberOf(group *Artist) *Artist {
	a.Groups = append(a.Groups, group)

	return a
}

func (a *Artist) AddStyle(style *Style) *Artist {
	a.node.Relate("CLASSIFIED_IN", style.Id, nil)

	return a
}

func (a *Artist) halify() {
	a.Links.Self = fmt.Sprintf("http://localhost:2000/artists/%d", a.Id)
}



type ArtistsManager struct {
	db *neoism.Database
}

func NewArtistsManager (db *neoism.Database) *ArtistsManager {
	return &ArtistsManager{
		db: db,
	}
}

func (am *ArtistsManager) FindById(id int) (*Artist, error) {
	results := []struct {
		A       neoism.Node
		RelType string
		NodeId  int
		N       neoism.Node
	}{}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (a:Artist)
			WHERE id(a) = {nodeId}
			OPTIONAL MATCH (a)-[r:HAS_SKILL|MEMBER_OF]->(n)
			RETURN a, type(r) AS relType, id(n) AS nodeId, n
			ORDER BY type(r), n.name
		`,
		Parameters: neoism.Props{"nodeId": id},
		Result:     &results,
	}

	err := am.db.Cypher(&cq)
	if err != nil {
		log.Fatal(err)
	}

	if len(results) == 0 {
		return nil, NotFound
	}

	artist := ArtistFromNode(&results[0].A)
	artist.Id = id

	for _, res := range results {
		if res.NodeId != 0 {
			fmt.Printf("%s\n", res.RelType)
			if res.RelType == "HAS_SKILL" {
				skill := SkillFromNode(&res.N)
				skill.Id = res.NodeId
				skill.halify()
				artist.hasSkill(skill)
			}

			if res.RelType == "MEMBER_OF" {
				group := ArtistFromNode(&res.N)
				group.Id = res.NodeId
				group.halify()
				artist.memberOf(group)
			}
		}
	}

	artist.halify()

	return artist, nil
}


func (am *ArtistsManager) Create(artistName string) *Artist {
	node, err := am.db.CreateNode(neoism.Props{"name": artistName})
	if err != nil {
		log.Fatal(err)
	}

	node.AddLabel("Artist")

	return &Artist{
		node:   node,
		Id:     node.Id(),
		Name:   artistName,
		Skills: []*Skill{},
	}
}


func (am *ArtistsManager) Find() Artists {
	results := []struct {
		A        neoism.Node
		ArtistId int
		S        neoism.Node
		SkillId  int
	}{}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (a:Artist)
			OPTIONAL MATCH (a)-[HAS_SKILL]->(s:Skill)
			RETURN a, id(a) AS artistId, s, id(s) AS skillId
			ORDER BY a.name, s.name
		`,
		Result: &results,
	}

	am.db.Cypher(&cq)

	artistsById := make(map[int]*Artist)

	artists := Artists{}

	for _, res := range results {
		_, ok := artistsById[res.ArtistId]
		if !ok {
			artistsById[res.ArtistId] = ArtistFromNode(&res.A)
			artistsById[res.ArtistId].Id = res.ArtistId
			artistsById[res.ArtistId].halify()

			artists = append(artists, artistsById[res.ArtistId])
		}

		if res.SkillId != 0 {
			skill := SkillFromNode(&res.S)
			skill.Id = res.SkillId
			skill.halify()
			artistsById[res.ArtistId].hasSkill(skill)
		}
	}

	return artists
}


func (am *ArtistsManager) Similars(a *Artist) Artists {
	results := []struct {
		B        neoism.Node
		ArtistId int
	}{}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (a:Artist)
			WHERE id(a) = {artistId}
			MATCH (a)-[c0:CLASSIFIED_IN]->(s:Style)<-[c1:CLASSIFIED_IN]-(b:Artist)
			RETURN b, id(b) AS artistId
			ORDER BY b.name
		`,
		Parameters: neoism.Props{"artistId": a.Id},
		Result: &results,
	}

	am.db.Cypher(&cq)

	artists := Artists{}

	for _, res := range results {
		artist := ArtistFromNode(&res.B)
		artist.Id = res.ArtistId
		artist.halify()

		artists = append(artists, artist)
	}
	
	return artists
}
