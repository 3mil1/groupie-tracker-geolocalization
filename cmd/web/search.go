package main

import (
	"strconv"
	"strings"
)

func searchInsideStruct(artists []ArtistsRelation, userInput string) interface{} {

	isResult := true
	var returnVal interface{}

	// result
	returnVal, isResult = result(artists, userInput, isResult)

	// suggestion
	if !isResult {
		returnVal = suggestion(artists, userInput)
	}

	return returnVal
}

func result(artists []ArtistsRelation, userInput string, isResult bool) (map[int]string, bool) {
	m := make(map[int]string)

	for _, r := range artists {
		sliceOfName := strings.Fields(r.Name)
		for _, name := range sliceOfName {
			if strings.HasPrefix(strings.ToLower(strings.TrimSpace(name)), strings.ToLower(userInput)) {
				m[r.ID] = r.Name
			}
		}

		for location := range r.DatesLocations {
			correctLocations := strings.ReplaceAll(strings.ReplaceAll(location, "_", " "), "-", " - ")
			if correctLocations == userInput {
				m[r.ID] = r.Name
			}
		}

		for _, member := range r.Members {
			if member == userInput {
				m[r.ID] = r.Name
			}
		}

		if strconv.Itoa(r.CreationDate) == userInput {
			m[r.ID] = r.Name
		}

		if r.FirstAlbum == userInput {
			m[r.ID] = r.Name
		}

	}

	if len(m) == 0 {
		isResult = false
	}

	return m, isResult
}

func suggestion(artists []ArtistsRelation, userInput string) []string {
	var m []string

	for _, r := range artists {
		for location := range r.DatesLocations {
			correctLocations := strings.ReplaceAll(strings.ReplaceAll(location, "_", " "), "-", " - ")
			sliceOfLocations := strings.Fields(correctLocations)

			for _, correct := range sliceOfLocations {
				if strings.HasPrefix(strings.ToLower(strings.TrimSpace(correct)), strings.ToLower(userInput)) {
					m = append(m, correctLocations)
				}
			}
		}

		for _, member := range r.Members {
			sliceOfName := strings.Fields(member)
			for _, name := range sliceOfName {
				if strings.HasPrefix(strings.ToLower(strings.TrimSpace(name)), strings.ToLower(userInput)) {
					m = append(m, member)
				}
			}
		}

		if strings.HasPrefix(strings.ToLower(strings.TrimSpace(strconv.Itoa(r.CreationDate))), strings.ToLower(userInput)) {
			m = append(m, strconv.Itoa(r.CreationDate))
		}

		if strings.HasPrefix(strings.ToLower(strings.TrimSpace(r.FirstAlbum)), strings.ToLower(userInput)) {
			m = append(m, r.FirstAlbum)
		}

	}
	return m
}
