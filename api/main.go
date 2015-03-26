package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmcvetta/neoism"
	"github.com/rs/cors"

	api "./api"
)

func main() {
	db, err := neoism.Connect("http://localhost:7474/db/data")
	if err != nil {
		log.Fatal(err)
	}

	labels  := api.NewLabelsManager(db)
	artists := api.NewArtistsManager(db)
	masters := api.NewMastersManager(db)
	//releases := api.NewReleasesManager(db)
	skills   := api.NewSkillsManager(db)
	styles   := api.NewStylesManager(db)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome!")
	})

	router.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		a := artists.Find(api.NewPager(r.URL.Query()), nil)

		api.JsonResponse(w, a)
	})

	router.HandleFunc("/artists/{artistId}", func(w http.ResponseWriter, r *http.Request) {
		artistId, err := strconv.Atoi(mux.Vars(r)["artistId"])
		if err != nil {
			log.Fatal(err)
		}

		artist, err := artists.FindById(artistId)
		if err != nil {
			log.Panic(err)
		}

		api.JsonResponse(w, artist)
	})

	router.HandleFunc("/artists/{artistId}/similars", func(w http.ResponseWriter, r *http.Request) {
		artistId, err := strconv.Atoi(mux.Vars(r)["artistId"])
		if err != nil {
			log.Fatal(err)
		}

		artist, err := artists.FindById(artistId)
		if err != nil {
			log.Panic(err)
		}

		similars := artists.Similars(artist)

		api.JsonResponse(w, similars)
	})

	router.HandleFunc("/artists/{artistId}/masters", func(w http.ResponseWriter, r *http.Request) {
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

		api.JsonResponse(w, artistMasters)
	})

	router.HandleFunc("/masters", func(w http.ResponseWriter, r *http.Request) {
		m := masters.Find(api.NewPager(r.URL.Query()))

		api.JsonResponse(w, m)
	})

	router.HandleFunc("/masters/{masterId}", func(w http.ResponseWriter, r *http.Request) {
		masterId, err := strconv.Atoi(mux.Vars(r)["masterId"])
		if err != nil {
			log.Fatal(err)
		}

		master, err := masters.FindById(masterId)
		if err != nil {
			log.Panic(err)
		}

		api.JsonResponse(w, master)
	})

	router.HandleFunc("/labels", func(w http.ResponseWriter, r *http.Request) {
		l := labels.Find()

		api.JsonResponse(w, l)
	})

	router.HandleFunc("/skills", func(w http.ResponseWriter, r *http.Request) {
		s := skills.Find(api.NewPager(r.URL.Query()))

		api.JsonResponse(w, s)
	})

	router.HandleFunc("/skills/{skillId}", func(w http.ResponseWriter, r *http.Request) {
		skillId, err := strconv.Atoi(mux.Vars(r)["skillId"])
		if err != nil {
			log.Fatal(err)
		}

		skill, err := skills.FindById(skillId)
		if err != nil {
			log.Panic(err)
		}

		api.JsonResponse(w, skill)
	})

	router.HandleFunc("/skills/{skillId}/artists", func(w http.ResponseWriter, r *http.Request) {
			skillId, err := strconv.Atoi(mux.Vars(r)["skillId"])
			if err != nil {
				log.Fatal(err)
			}

			a := artists.Find(api.NewPager(r.URL.Query()), &api.ArtistsFilters{
				SkillId: skillId,
			})

			api.JsonResponse(w, a)
	})

	router.HandleFunc("/labels/{labelId}", func(w http.ResponseWriter, r *http.Request) {
		labelId, err := strconv.Atoi(mux.Vars(r)["labelId"])
		if err != nil {
			log.Fatal(err)
		}

		label, err := labels.FindById(labelId)
		if err != nil {
			log.Panic(err)
		}

		api.JsonResponse(w, label)
	})

	router.HandleFunc("/styles", func(w http.ResponseWriter, r *http.Request) {
		s := styles.Find(api.NewPager(r.URL.Query()))

		api.JsonResponse(w, s)
	})

	router.HandleFunc("/styles/{styleId}", func(w http.ResponseWriter, r *http.Request) {
		styleId, err := strconv.Atoi(mux.Vars(r)["styleId"])
		if err != nil {
			log.Fatal(err)
		}

		style, err := styles.FindById(styleId)
		if err != nil {
			log.Panic(err)
		}

		api.JsonResponse(w, style)
	})

	router.HandleFunc("/styles/{styleId}/artists", func(w http.ResponseWriter, r *http.Request) {
		styleId, err := strconv.Atoi(mux.Vars(r)["styleId"])
		if err != nil {
			log.Fatal(err)
		}

		a := artists.Find(api.NewPager(r.URL.Query()), &api.ArtistsFilters{
			StyleId: styleId,
		})

		api.JsonResponse(w, a)
	})

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":2000", handler))
}
