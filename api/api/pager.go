package api

import (
	//"fmt"
	"strconv"
	"net/url"
	"math"
)

type Pager struct {
	Total     int `json:"total"`
	PerPage   int `json:"per_page"`
	Page      int `json:"page"`
	PageCount int `json:"page_count"`
}

func NewPager(values url.Values) *Pager {
	urlPage := values.Get("page")
	page := 1
	if urlPage != "" {
		p, err := strconv.Atoi(urlPage)
		if err == nil && p > 0 {
			page = p
		}
	}

	urlPerPage := values.Get("per_page")
	perPage := 10
	if urlPerPage != "" {
		pp, err := strconv.Atoi(urlPerPage)
		if err == nil {
			perPage = pp
		}
	}

	return &Pager{
		PerPage:   perPage,
		Page:      page,
		PageCount: 0,
	}
}

func (p *Pager) SetTotal(total int) {
	p.Total = total
	p.PageCount = int(math.Ceil(float64(total) / float64(p.PerPage)))
}

func (p *Pager) Offset() int {
	return (p.Page - 1) * p.PerPage
}
