package main

import (
	"log"

	"github.com/jmcvetta/neoism"

	api "./api"
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
	masters  := api.NewMastersManager(db)
	releases := api.NewReleasesManager(db)
	skills   := api.NewSkillsManager(db)
	styles   := api.NewStylesManager(db)

	style_jazz     := styles.Create("jazz")
	style_grunge   := styles.Create("grunge")
	style_country  := styles.Create("country")
	style_rock     := styles.Create("rock")
	style_hiphop   := styles.Create("hip hop")
	style_alt_rock := styles.Create("alternative rock")

	skill_vocals := skills.Create("vocals")
	skill_bass   := skills.Create("bass")
	skill_piano  := skills.Create("piano")
	skill_guitar := skills.Create("guitar")
	skill_drums  := skills.Create("drums")

	label_blue_note := labels.Create("Blue Note")
	label_verve     := labels.Create("Verve")

	artists.Create("Wes Montgomery").AddSkill(skill_guitar).AddStyle(style_jazz)
	artists.Create("Alton Elis").AddSkill(skill_vocals)
	artists.Create("Toots and The Maytals")
	artists.Create("Nina Simone").AddSkill(skill_vocals).AddStyle(style_jazz)
	artists.Create("Nat King Cole").AddStyle(style_jazz)
	artists.Create("Bon Iver").AddSkill(skill_vocals).AddSkill(skill_guitar)
	artists.Create("Beirut")
	artists.Create("The Roots").AddStyle(style_hiphop)
	artists.Create("Fats Waller").AddSkill(skill_vocals).AddSkill(skill_piano).AddStyle(style_jazz)
	artists.Create("Louis Jordan").AddSkill(skill_vocals).AddStyle(style_jazz)
	artists.Create("Coleman Hawkins").AddStyle(style_jazz)
	the_smiths := artists.Create("The Smiths").AddStyle(style_rock)
	artists.Create("Morissey").AddSkill(skill_vocals).AddMembership(the_smiths).AddSkill(skill_vocals).AddStyle(style_rock)
	artists.Create("Serge Gainsbourg").AddSkill(skill_vocals).AddSkill(skill_piano)
	artists.Create("Joy Division").AddStyle(style_rock)
	artists.Create("Arcade Fire").AddStyle(style_rock)
	artists.Create("John Fahey").AddSkill(skill_guitar).AddStyle(style_country)

	art := artists.Create("Art Blakey").AddStyle(style_jazz)
	jazz_messengers := artists.Create("Art Blakey & The Jazz Messengers")
	art_afro_cub := artists.Create("Art Blakey And His Afro Cuban Boys")
	art.AddMembership(jazz_messengers).AddMembership(art_afro_cub)
	night_in := masters.Create("Night in Tunisia")
	night_in_rel := releases.Create("")
	night_in_rel.ProducedBy(label_blue_note)
	night_in.AddRelease(night_in_rel)
	art.AddSkill(skill_drums).PlayedIn(night_in)

	astrud := artists.Create("Astrud Gilberto").AddStyle(style_jazz)
	best_of_astrud := masters.Create("The very best of Astrud Gilberto")
	astrud.AddSkill(skill_vocals).PlayedIn(best_of_astrud)
	best_of_astrud_rel := releases.Create("")
	best_of_astrud_rel.ProducedBy(label_verve)
	best_of_astrud.AddRelease(best_of_astrud_rel)

	duke := artists.Create("Duke Ellington").AddStyle(style_jazz)
	max := artists.Create("Max Roach").AddStyle(style_jazz)
	charles := artists.Create("Charles Mingus").AddStyle(style_jazz)
	money_jungle := masters.Create("Money Jungle")
	money_jungle_rel := releases.Create("")
	money_jungle_rel.ProducedBy(label_blue_note)
	money_jungle.AddRelease(money_jungle_rel)
	duke.PlayedIn(money_jungle).AddSkill(skill_piano).AddSkill(skill_vocals)
	max.PlayedIn(money_jungle).AddSkill(skill_drums)
	charles.PlayedIn(money_jungle).AddSkill(skill_bass)

	lithium := masters.Create("lithium")
	in_utero := masters.Create("in utero")
	bleach := masters.Create("bleach")
	nirvana := artists.Create("Nirvana")
	nirvana.AddStyle(style_grunge).AddStyle(style_alt_rock)
	kurt := artists.Create("Kurt Cobain")
	kurt.AddMembership(nirvana)
	kurt.AddStyle(style_grunge).AddStyle(style_alt_rock)
	kurt.PlayedIn(lithium).PlayedIn(in_utero).PlayedIn(bleach)
	kurt.AddSkill(skill_guitar).AddSkill(skill_vocals)
}
