package main

import (
	"log"
	"net/http"
	"path/filepath"
)

type Artists []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var port = ":8080"
var artists Artists
var relation Relation

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/artist", artistPage)
	mux.HandleFunc("/search", search)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("ui/fileServer")})
	mux.Handle("/fileServer", http.NotFoundHandler())
	mux.Handle("/fileServer/", http.StripPrefix("/fileServer", fileServer))

	log.Println("Serving@:", "http://localhost"+port)
	log.Fatal(http.ListenAndServe(port, mux))

}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}
	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
