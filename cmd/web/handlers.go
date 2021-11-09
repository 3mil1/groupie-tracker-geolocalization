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

	pkg.JsonFromAPI("https://groupietrackers.herokuapp.com/api/artists", &artists)

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	for key, value := range r.Form {
		fmt.Println(key, value)
	}

	err = templates.ExecuteTemplate(w, "index", artists)
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

	for i := range relation.DatesLocations {
		delete(relation.DatesLocations, i)
	}

	pkg.JsonFromAPI(artists[id-1].Relations, &relation)

	changeKeyInMap(relation.DatesLocations)

	artistWRelation := []interface{}{
		artists[id-1],
		relation,
	}

	err = templates.ExecuteTemplate(w, "artist", artistWRelation)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func changeKeyInMap(dataMap map[string][]string) {
	space := regexp.MustCompile(`\s+`)

	for i, e := range dataMap {
		city := strings.ReplaceAll(strings.ReplaceAll(i, "_", " "), "-", " - ")
		str := strings.Title(space.ReplaceAllString(city, " "))
		delete(dataMap, i)
		dataMap[str] = e
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	s := make(map[int]string)

	pkg.JsonFromAPI("https://groupietrackers.herokuapp.com/api/artists", &artists)

	switch r.Method {
	case "GET":
		a := r.URL.Query().Get("a")

		for _, r := range artists {
			if strings.HasPrefix(strings.ToLower(r.Name), strings.ToLower(a)) {
				s[r.ID] = r.Name
			}
		}
	}
	json.NewEncoder(w).Encode(s)
}
