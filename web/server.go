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

func NewRaspiMusicServer(port int) *RaspiMusicServer {
	var server RaspiMusicServer
	server.player = player.NewPlayer()
	server.m = martini.Classic()
	server.port = fmt.Sprintf(":%d", port)
	server.registerRoutes()
	return &server
}

func (s RaspiMusicServer) registerRoutes() {
	s.m.Get("/songs", s.getSongsList)
	s.m.Post("/songs/play", s.playSong)
	s.m.Post("/stop", s.stop)
	s.m.Post("/next", s.next)
	s.m.Post("/prev", s.prev)
}

func (s RaspiMusicServer) stop() (int, string) {
	err := s.player.Stop()
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	return 200, "Stopped"
}

func (s RaspiMusicServer) next() (int, string) {
	err := s.player.Next()
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	return 200, "Next Song"
}

func (s RaspiMusicServer) prev() (int, string) {
	err := s.player.Prev()
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	return 200, "Prev Song"
}

func (s RaspiMusicServer) playSong(w http.ResponseWriter, r *http.Request) (int, string) {
	path, err := s.player.AddSong(r.FormValue("path"))
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	s.player.Play()
	return 200, "PLAYING.. " + path
}

func (s RaspiMusicServer) getSongsList() (int, string) {
	attr, err := s.player.Songs()
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	songs, err := json.Marshal(attr)
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	return 200, string(songs)
}

func (s RaspiMusicServer) Run() {
	http.ListenAndServe(s.port, s.m)
}
