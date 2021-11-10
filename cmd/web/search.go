package main

import "strings"

func searchInsideStruct(artists []ArtistsRelation, userInput string) map[int]string {
	m := make(map[int]string)

	for _, r := range artists {

		for _, member := range r.Members {
			if strings.HasPrefix(strings.ToLower(strings.TrimSpace(member)), strings.ToLower(userInput)) {
				m[r.ID] = member
			}
		}

		if strings.HasPrefix(strings.ToLower(strings.TrimSpace(r.Name)), strings.ToLower(userInput)) {
			m[r.ID] = r.Name
		}

	}
	return m
}
