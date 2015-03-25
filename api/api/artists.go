package api

import (
	"fmt"
	"log"

	"github.com/jmcvetta/neoism"
)

type Artist struct {
	ApiNode
	Name    string    `json:"name"`
	Picture string    `json:"picture"`
	Skills  []*Skill  `json:"skills"`
	Groups  []*Artist `json:"groups"`
	Styles  []*Style  `json:"styles"`
}

type ArtistsCollection struct {
	ApiCollection
	Results *Artists `json:"results"`
}

func (ac *ArtistsCollection) halify() {
	ac.Links.Self = fmt.Sprintf("http://localhost:2000/artists?page=%d", ac.Pager.Page)
	ac.Links.Prev = fmt.Sprintf("http://localhost:2000/artists?page=%d", ac.Pager.Page-1)
	ac.Links.Next = fmt.Sprintf("http://localhost:2000/artists?page=%d", ac.Pager.Page+1)
}

type Artists []*Artist

func ArtistFromNode(node *neoism.Node) *Artist {
	name    := node.Data["name"].(string)
	picture := node.Data["picture"].(string)

	return &Artist{
		ApiNode: ApiNode{node: node},
		Name:    name,
		Picture: picture,
		Skills:  []*Skill{},
		Groups:  []*Artist{},
		Styles:  []*Style{},
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

func NewArtistsManager(db *neoism.Database) *ArtistsManager {
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
			OPTIONAL MATCH (a)-[r:HAS_SKILL|MEMBER_OF|CLASSIFIED_IN]->(n)
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
			switch {
			case res.RelType == "HAS_SKILL":
				skill := SkillFromNode(&res.N)
				skill.Id = res.NodeId
				skill.halify()
				artist.Skills = append(artist.Skills, skill)
			case res.RelType == "MEMBER_OF":
				group := ArtistFromNode(&res.N)
				group.Id = res.NodeId
				group.halify()
				artist.Groups = append(artist.Groups, group)
			case res.RelType == "CLASSIFIED_IN":
				style := StyleFromNode(&res.N)
				style.Id = res.NodeId
				style.halify()
				artist.Styles = append(artist.Styles, style)
			}
		}
	}

	artist.halify()

	return artist, nil
}

func (am *ArtistsManager) Create(artistName string, picture string) *Artist {
	node, err := am.db.CreateNode(neoism.Props{
		"name":    artistName,
		"picture": picture,
	})
	if err != nil {
		log.Fatal(err)
	}

	node.AddLabel("Artist")

	return &Artist{
		ApiNode: ApiNode{
			node: node,
			Id:   node.Id(),
		},
		Name:    artistName,
		Picture: picture,
		Skills:  []*Skill{},
	}
}

type ArtistsFilters struct {
	SkillId int
	StyleId int
}

func (am *ArtistsManager) Find(pager *Pager, filters *ArtistsFilters) *ArtistsCollection {
	results := []struct {
		A        neoism.Node
		ArtistId int
		RelType  string
		N        neoism.Node
		NodeId   int
		Total    int
	}{}

	statement := `MATCH (a:Artist), (b:Artist)`
	params    := neoism.Props{
		"offset":  pager.Offset(),
		"limit":   pager.PerPage,
	}

	if filters != nil {
		if filters.SkillId != 0 {
			statement = statement + `
MATCH (s:Skill)<-[HAS_SKILL]-(a)
MATCH (s:Skill)<-[HAS_SKILL]-(b)
WHERE id(s) = {skillId}`
			params["skillId"] = filters.SkillId
		}
	}

	statement = statement + `
WITH a, count(b) AS total
ORDER BY a.name ASC
SKIP {offset} LIMIT {limit}
OPTIONAL MATCH (a)-[r:HAS_SKILL|CLASSIFIED_IN]->(n)
RETURN a, id(a) AS artistId, type(r) AS relType, n, id(n) AS nodeId, total`

	//fmt.Printf("%#v\n", filters)
	//fmt.Printf("%s\n", statement)

	cq := neoism.CypherQuery{
		Statement:  statement,
		Parameters: params,
		Result:     &results,
	}

	am.db.Cypher(&cq)

	artistsById := make(map[int]*Artist)

	artists := Artists{}

	if len(results) > 0 {
		pager.SetTotal(results[0].Total)
		for _, res := range results {
			_, ok := artistsById[res.ArtistId]
			if !ok {
				artistsById[res.ArtistId] = ArtistFromNode(&res.A)
				artistsById[res.ArtistId].Id = res.ArtistId
				artistsById[res.ArtistId].halify()

				artists = append(artists, artistsById[res.ArtistId])
			}

			if res.NodeId != 0 {
				switch {
				case res.RelType == "HAS_SKILL":
					skill := SkillFromNode(&res.N)
					skill.Id = res.NodeId
					skill.halify()
					artistsById[res.ArtistId].hasSkill(skill)
				case res.RelType == "CLASSIFIED_IN":
					style := StyleFromNode(&res.N)
					style.Id = res.NodeId
					style.halify()
					artistsById[res.ArtistId].Styles = append(artistsById[res.ArtistId].Styles, style)
				}
			}
		}
	}

	collection := ArtistsCollection{
		ApiCollection: ApiCollection{Pager: pager},
		Results:       &artists,
	}

	collection.halify()

	return &collection
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
		Result:     &results,
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
