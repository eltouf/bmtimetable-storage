//Package to fetch data from https://opendata.bordeaux-metropole.fr/api/datasets
package main

import (
	"net/url"
	"github.com/eltouf/bmtimetable-storage/functions/fetchdatasets"
)

func main() {
	filters := &url.Values{}
	filters.Set("rows", "50")
	filters.Set("refine.keyword", "saeiv")

	fetchdatasets.FetchDatasets(filters)
}
