package api

type Links struct {
	Self string `json:"self"`
}

type CollectionLinks struct {
	Links
	Prev string `json:"prev"`
	Next string `json:"next"`
}

type HalCollection struct {
	Links CollectionLinks `json:"_links"`
}

type Embedded struct {

}
