package api

import (
	//"fmt"
	"log"

	"github.com/jmcvetta/neoism"
)

type Year struct {
	ApiNode
	Year int `json:"year"`
}

func (y *Year) AddMaster(m *Master) {

}

type CalendarManager struct {
	db *neoism.Database
}

func NewCalendarManager(db *neoism.Database) *CalendarManager {
	return &CalendarManager{
		db: db,
	}
}

func (cm *CalendarManager) CreateYear(year int) *Year {
	node, err := cm.db.CreateNode(neoism.Props{"year": year})
	if err != nil {
		log.Fatal(err)
	}

	node.AddLabel("Year")

	return &Year{
		ApiNode: ApiNode{
			node: node,
			Id:   node.Id(),
		},
		Year:   year,
	}
}
