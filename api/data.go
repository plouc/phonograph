package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jmcvetta/neoism"
	"gopkg.in/yaml.v2"

	api "./api"
)

type Fixtures struct {
	Year   []int
	Artist map[string]struct {
		Name    string
		Picture string
		Skills  []string
		Styles  []string
		Groups  []string
	}
	Label  map[string]string
	Skill  map[string]string
	Style  map[string]string
	Master map[string]struct {
		Name    string
		Year    int
		Picture string
		Artists []string
		Styles  []string
		Tracks  []struct {
			Name     string
			Duration int
		}
		Releases []struct {
			Year    int
			Label   string
			Country string
		}
	}
}

type Refs struct {
	years    map[int]*api.Year
	artists  map[string]*api.Artist
	labels   map[string]*api.Label
	skills   map[string]*api.Skill
	styles   map[string]*api.Style
	masters  map[string]*api.Master
	releases map[string]*api.Release
}

func main() {
	data, err := ioutil.ReadFile("./sample.yml")
	if err != nil {
		panic(err)
	}

	f := Fixtures{}

	err = yaml.Unmarshal([]byte(data), &f)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	db, err := neoism.Connect("http://localhost:7474/db/data")
	if err != nil {
		log.Fatal(err)
	}

	// cleanup
	cq := neoism.CypherQuery{
		Statement: "MATCH (n) OPTIONAL MATCH (n)-[r]-() DELETE n,r",
	}
	db.Cypher(&cq)

	labels   := api.NewLabelsManager(db)
	artists  := api.NewArtistsManager(db)
	masters  := api.NewMastersManager(db)
	tracks   := api.NewTracksManager(db)
	releases := api.NewReleasesManager(db)
	skills   := api.NewSkillsManager(db)
	styles   := api.NewStylesManager(db)
	calendar := api.NewCalendarManager(db)

	refs := Refs{
		years:    make(map[int]*api.Year),
		labels:   make(map[string]*api.Label),
		artists:  make(map[string]*api.Artist),
		masters:  make(map[string]*api.Master),
		releases: make(map[string]*api.Release),
		skills:   make(map[string]*api.Skill),
		styles:   make(map[string]*api.Style),
	}

	//////////////////
	// Create nodes //
	//////////////////
	fmt.Println("\nYears:")
	for _, y := range f.Year {
		fmt.Printf("> creating year '%d'\n", y)
		refs.years[y] = calendar.CreateYear(y)
	}

	fmt.Println("\nArtists:")
	for k, a := range f.Artist {
		fmt.Printf("> creating artist '%s' (%s)\n", a.Name, a.Picture)
		refs.artists[k] = artists.Create(a.Name, a.Picture)
	}

	fmt.Println("\nLabels:")
	for k, l := range f.Label {
		fmt.Printf("> creating label '%s'\n", l)
		refs.labels[k] = labels.Create(l)
	}

	fmt.Println("\nSkills:")
	for k, s := range f.Skill {
		fmt.Printf("> creating skill '%s'\n", s)
		refs.skills[k] = skills.Create(s)
	}

	fmt.Println("\nStyles:")
	for k, s := range f.Style {
		fmt.Printf("> creating style '%s'\n", s)
		refs.styles[k] = styles.Create(s)
	}

	fmt.Println("\nMasters:")
	for k, m := range f.Master {
		fmt.Printf("> creating master '%s'\n", m.Name)
		master := masters.Create(m.Name, m.Year, m.Picture)
		fmt.Printf("> add master '%s' to year '%d'\n", m.Name, m.Year)
		refs.years[m.Year].AddMaster(master)
		for _, t := range m.Tracks {
			fmt.Printf("> adding track '%s' to '%s'\n", t.Name, m.Name)
			track := tracks.Create(t.Name, t.Duration)
			master.AddTrack(track)
		}
		refs.masters[k] = master
	}

	/////////////////////
	// Build relations //
	/////////////////////
	fmt.Println("\nArtists relations:")
	for k, a := range f.Artist {
		ref := refs.artists[k]
		for _, g := range a.Groups {
			fmt.Printf("> add group '%s' to '%s'\n", g, ref.Name)
			ref.AddMembership(refs.artists[g])
		}
		for _, s := range a.Skills {
			fmt.Printf("> add skill '%s' to '%s'\n", s, ref.Name)
			ref.AddSkill(refs.skills[s])
		}
		for _, s := range a.Styles {
			fmt.Printf("> add style '%s' to '%s'\n", s, ref.Name)
			ref.AddStyle(refs.styles[s])
		}
	}

	fmt.Println("\nMasters relations:")
	for k, m := range f.Master {
		ref := refs.masters[k]
		for _, a := range m.Artists {
			fmt.Printf("> add artist '%s' to '%s'\n", a, ref.Name)
			refs.artists[a].PlayedIn(ref)
		}
		for _, s := range m.Styles {
			fmt.Printf("> add style '%s' to '%s'\n", s, ref.Name)
			ref.AddStyle(refs.styles[s])
		}

		for _, r := range m.Releases {
			fmt.Printf("> add release '%s' (%d)\n", ref.Name, r.Year)
			rel := releases.Create(r.Year, r.Country)
			rel.AddLabel(refs.labels[r.Label])
			ref.AddRelease(rel)
		}
	}
}
