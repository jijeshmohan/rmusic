package server

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/martini"
	"net/http"
	"raspimusic/song"
)

type RaspiMusicServer struct {
	m               *martini.ClassicMartini
	port            string
	songsCollection song.SongsCollection
}

func NewRaspiMusicServer(port int, songsPath string) *RaspiMusicServer {
	var server RaspiMusicServer

	server.m = martini.Classic()
	server.songsCollection = song.SongsCollection(songsPath)
	server.port = fmt.Sprintf(":%d", port)
	server.registerRoutes()
	return &server
}

func (server RaspiMusicServer) registerRoutes() {
	server.m.Get("/songs", server.getSongsList)
}

func (server RaspiMusicServer) getSongsList() (int, string) {
	songs, err := json.Marshal(server.songsCollection.All())
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	return 200, string(songs)
}
func (server RaspiMusicServer) Run() {
	http.ListenAndServe(server.port, server.m)
}
