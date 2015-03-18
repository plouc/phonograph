package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmcvetta/neoism"
)

type Links struct {
	Self string `json:"self"`
}

type ApiCollectionLinks struct {
	Links
	Prev string `json:"prev"`
	Next string `json:"next"`
}

type ApiCollection struct {
	Links ApiCollectionLinks `json:"_links"`
	Pager *Pager             `json:"pager"`
}

type ApiNode struct {
	node  *neoism.Node
	Id    int   `json:"id"`
	Links Links `json:"_links"`
}

func JsonResponse(w http.ResponseWriter, data interface{}) {
	m, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(m))
}
