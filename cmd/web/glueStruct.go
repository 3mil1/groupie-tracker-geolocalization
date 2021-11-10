package main

import (
	"server/cmd/pkg"
)

type ArtistsRelation struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	DatesLocations map[string][]string `json:"relationsSlice"`
}

type Relations struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}
}

var artistsRelation []ArtistsRelation
var relations Relations

func glueStruct() {
	pkg.JsonFromAPI("https://groupietrackers.herokuapp.com/api/artists", &artistsRelation)
	pkg.JsonFromAPI("https://groupietrackers.herokuapp.com/api/relation", &relations)

	for _, ar := range artistsRelation {
		for _, r := range relations.Index {
			if r.ID == ar.ID {
				// add relation to artist
				artistsRelation[ar.ID-1].add(r.DatesLocations)
			}
		}
	}

}

func (ar *ArtistsRelation) add(rel map[string][]string) {
	ar.DatesLocations = rel
}
