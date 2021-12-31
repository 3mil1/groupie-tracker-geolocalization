package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"server/cmd/pkg"
	"strconv"
	"strings"
)

var templates = template.Must(template.ParseGlob("./ui/html/*"))

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	//for key, value := range r.Form {
	//	fmt.Println(key, value)
	//}

	err = templates.ExecuteTemplate(w, "index", &artistsRelation)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func artistPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 || id > 52 {
		http.NotFound(w, r)
		return
	}

	//googleMapSlice(artistsRelation[id-1].DatesLocations)

	changeKeyInMap(artistsRelation[id-1].DatesLocations)

	err = templates.ExecuteTemplate(w, "artist", artistsRelation[id-1])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

// from los_angeles-usa to Los Angeles - Usa:
func changeKeyInMap(dataMap map[string][]string) {

	space := regexp.MustCompile(`\s+`)

	for i, e := range dataMap {
		city := strings.ReplaceAll(strings.ReplaceAll(i, "_", " "), "-", " - ")
		str := strings.Title(space.ReplaceAllString(city, " "))
		delete(dataMap, i)
		dataMap[str] = e
	}
}

func googleMapSlice(dataMap map[string][]string) []interface{} {
	var SGMap []interface{}

	for i := range dataMap {
		var GMap Map
		pkg.JsonFromAPI("https://maps.googleapis.com/maps/api/geocode/json?address="+i+"&key=AIzaSyC4IMt9zKp20yjBhUZBMhA0qlqcQjau0dY&callback", &GMap)
		fmt.Println(GMap)
		SGMap = append(SGMap, GMap.AddressComponents[0].Geometry.Location)
	}

	return SGMap
}

func search(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		userInput := r.URL.Query().Get("a")
		//fmt.Println(searchInsideStruct(artistsRelation, userInput))
		json.NewEncoder(w).Encode(searchInsideStruct(artistsRelation, userInput))
	}
}

type Map struct {
	AddressComponents []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}
