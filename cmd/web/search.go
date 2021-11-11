package main

import (
	"strconv"
	"strings"
)

func searchInsideStruct(artists []ArtistsRelation, userInput string) []interface{} {

	var returnVal []interface{}
	var returnValRES map[int]string
	var returnValSUG []string

	// result
	returnValRES = result(artists, userInput)

	// suggestion
	//if !isResult {
	returnValSUG = suggestion(artists, userInput)
	//}

	returnVal = append(returnVal, returnValRES)
	returnVal = append(returnVal, returnValSUG)

	return returnVal
}

func result(artists []ArtistsRelation, userInput string) map[int]string {
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
	return m
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
