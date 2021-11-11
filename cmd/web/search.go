package main

import (
	"strings"
)

type Suggestions struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	DatesLocations map[string][]string `json:"relationsSlice"`
}

var suggestions Suggestions

func searchInsideStruct(artists []ArtistsRelation, userInput string) interface{} {

	isResult := true
	var returnVal interface{}

	//for _, member := range r.Members {
	//	if strings.HasPrefix(strings.ToLower(strings.TrimSpace(member)), strings.ToLower(userInput)) {
	//		m[r.ID] = member + " -> member"
	//		members := []string{member}
	//		suggestions.add(members)
	//	}
	//}

	// result
	returnVal, isResult = result(artists, userInput, isResult)

	// suggestion
	if !isResult {
		returnVal = suggestion(artists, userInput)
	}

	//data, _ := json.Marshal(suggestions)
	//fmt.Println(string(data))

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
	}
	return m
}

func (sugg *Suggestions) add(data []string) {
	sugg.Members = data
}
