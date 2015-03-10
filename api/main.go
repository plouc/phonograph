package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	api "./api"

	"encoding/json"

	"github.com/jmcvetta/neoism"

	"github.com/gorilla/mux"
)

func main() {

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
	releases := api.NewReleasesManager(db)
	skills   := api.NewSkillsManager(db)

	skill_vocals := skills.Create("vocals")
	skill_bass   := skills.Create("bass")
	skill_piano  := skills.Create("piano")
	skill_drums  := skills.Create("drums")

	blue_note    := labels.Create("Blue Note")

	art      := artists.Create("Art Blakey")
	night_in := releases.Create("Night in Tunisia")
	art.AddSkill(skill_drums).PlayedIn(night_in)
	night_in.ProducedBy(blue_note)

	astrud         := artists.Create("Astrud Gilberto")
	best_of_astrud := releases.Create("The very best of Astrud Gilberto")
	verve          := labels.Create("Verve")
	astrud.AddSkill(skill_vocals).PlayedIn(best_of_astrud)
	best_of_astrud.ProducedBy(verve)

	duke         := artists.Create("Duke Ellington")
	max          := artists.Create("Max Roach")
	charles      := artists.Create("Charles Mingus")
	money_jungle := releases.Create("Money Jungle")
	duke.PlayedIn(money_jungle).AddSkill(skill_piano).AddSkill(skill_vocals)
	max.PlayedIn(money_jungle).AddSkill(skill_drums)
	charles.PlayedIn(money_jungle).AddSkill(skill_bass)
	money_jungle.ProducedBy(blue_note)


	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome!")
	})

	router.HandleFunc("/artists", func (w http.ResponseWriter, r *http.Request) {
		a := artists.Find()

		b, err := json.MarshalIndent(a, "", "  ")
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Fprintln(w, string(b))
	})

	router.HandleFunc("/artists/{artistId}", func (w http.ResponseWriter, r * http.Request) {
		vars := mux.Vars(r)
		artistId, err := strconv.Atoi(vars["artistId"])
		if err != nil {
			log.Fatal(err)
		}

		artist, err := artists.FindById(artistId)
		if err != nil {
			log.Panic(err)
		}

		b, err := json.MarshalIndent(artist, "", "  ")
		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Fprintln(w, string(b))
	})

	router.HandleFunc("/labels", func (w http.ResponseWriter, r *http.Request) {
		l := labels.Find()

		b, err := json.MarshalIndent(l, "", "  ")
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Fprintln(w, string(b))
	})


	log.Fatal(http.ListenAndServe(":2000", router))
}
