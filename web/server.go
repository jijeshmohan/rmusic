package web

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/jijeshmohan/raspimusic/player"
	"net/http"
)

type RaspiMusicServer struct {
	m      *martini.ClassicMartini
	port   string
	player *player.MPDPlayer
}

func NewRaspiMusicServer(port int, songsPath string) *RaspiMusicServer {
	var server RaspiMusicServer
	server.player = player.NewPlayer()
	server.m = martini.Classic()
	server.port = fmt.Sprintf(":%d", port)
	server.registerRoutes()
	return &server
}

func (s RaspiMusicServer) registerRoutes() {
	s.m.Get("/songs", s.getSongsList)
	s.m.Post("/songs/:id/play", s.PlaySong)
	s.m.Post("/stop", s.Stop)
	s.m.Post("/next", s.Next)
	s.m.Post("/prev", s.Prev)
}
func (s RaspiMusicServer) Stop() (int, string) {
	err := s.player.Stop()
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	return 200, "Stopped"
}

func (s RaspiMusicServer) Next() (int, string) {
	err := s.player.Next()
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	return 200, "Next Song"
}

func (s RaspiMusicServer) Prev() (int, string) {
	err := s.player.Prev()
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	return 200, "Prev Song"
}

func (s RaspiMusicServer) PlaySong(params martini.Params) (int, string) {
	id := params["id"]
	path, err := s.player.AddSong(id)
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	s.player.Play()
	return 200, "PLAYING.. " + path
}

func (s RaspiMusicServer) getSongsList() (int, string) {
	songs, _ := json.Marshal(s.player.Songs())
	return 200, string(songs)
}

func (s RaspiMusicServer) Run() {
	http.ListenAndServe(s.port, s.m)
}
