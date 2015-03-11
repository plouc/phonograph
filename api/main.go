package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/rs/cors"
	"github.com/jmcvetta/neoism"
	"github.com/gorilla/mux"

	api "./api"
)

func main() {
	db, err := neoism.Connect("http://localhost:7474/db/data")
	if err != nil {
		log.Fatal(err)
	}

	labels   := api.NewLabelsManager(db)
	artists  := api.NewArtistsManager(db)
	masters  := api.NewMastersManager(db)
	//releases := api.NewReleasesManager(db)
	//skills   := api.NewSkillsManager(db)
	//styles   := api.NewStylesManager(db)

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

		w.Header().Set("Content-Type", "application/json")
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

		j, err := json.MarshalIndent(artist, "", "  ")
		if err != nil {
			fmt.Println("error:", err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(j))
	})

	router.HandleFunc("/artists/{artistId}/similars", func (w http.ResponseWriter, r * http.Request) {
		vars := mux.Vars(r)
		artistId, err := strconv.Atoi(vars["artistId"])
		if err != nil {
			log.Fatal(err)
		}

		artist, err := artists.FindById(artistId)
		if err != nil {
			log.Panic(err)
		}

		similars := artists.Similars(artist)

		j, err := json.MarshalIndent(similars, "", "  ")
		if err != nil {
			fmt.Println("error:", err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(j))
	})

	router.HandleFunc("/artists/{artistId}/masters", func (w http.ResponseWriter, r * http.Request) {
		vars := mux.Vars(r)
		artistId, err := strconv.Atoi(vars["artistId"])
		if err != nil {
			log.Fatal(err)
		}

		artist, err := artists.FindById(artistId)
		if err != nil {
			log.Panic(err)
		}

		artistMasters := masters.PlayedBy(artist)

		j, err := json.MarshalIndent(artistMasters, "", "  ")
		if err != nil {
			fmt.Println("error:", err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(j))
	})

	router.HandleFunc("/masters", func (w http.ResponseWriter, r *http.Request) {
		m := masters.Find()

		b, err := json.MarshalIndent(m, "", "  ")
		if err != nil {
			fmt.Println("error:", err)
		}

		w.Header().Set("Content-Type", "application/json")
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

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":2000", handler))
}
