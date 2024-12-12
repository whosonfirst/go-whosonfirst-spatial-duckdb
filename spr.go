package duckdb

import (
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
)

type SPRResults struct {
	spr.StandardPlacesResults `json:",omitempty"`
	SPRResults []spr.StandardPlacesResult `json:"results"`
}

func (r *SPRResults) Results() []spr.StandardPlacesResult {
	return r.SPRResults
}

func NewSPRResults(results []spr.StandardPlacesResult) spr.StandardPlacesResults {

	return &SPRResults{
		SPRResults: results,
	}
}

