package server

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/martini"
	"net/http"
	"raspimusic/player"
	"raspimusic/song"
)

type RaspiMusicServer struct {
	m               *martini.ClassicMartini
	port            string
	songsCollection song.SongsCollection
	player          *player.MPDPlayer
}

func NewRaspiMusicServer(port int, songsPath string) *RaspiMusicServer {
	var server RaspiMusicServer
	server.player = player.NewPlayer()
	server.m = martini.Classic()
	server.songsCollection = song.SongsCollection(songsPath)
	server.port = fmt.Sprintf(":%d", port)
	server.registerRoutes()
	return &server
}

func (server RaspiMusicServer) registerRoutes() {
	server.m.Get("/songs", server.getSongsList)
	server.m.Post("/songs/:id/play", server.PlaySong)
	server.m.Post("/stop", server.Stop)

}

func (server RaspiMusicServer) Stop() (int, string) {
	server.player.Stop()
	return 200, "Stopped"
}

func (server RaspiMusicServer) PlaySong(params martini.Params) (int, string) {
	id := params["id"]
	path, err := server.songsCollection.SongPath(id)
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	server.player.AddSong(path)
	server.player.Play()
	return 200, "PLAYING.. " + path
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
