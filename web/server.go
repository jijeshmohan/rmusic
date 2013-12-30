package web

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/jijeshmohan/raspimusic/player"
	"net/http"
	"time"
)

const (
	duration = 1
)

type RaspiMusicServer struct {
	m         *martini.ClassicMartini
	port      string
	player    *player.MPDPlayer
	stopTimer *time.Timer
}

func NewRaspiMusicServer(port int) *RaspiMusicServer {
	var server RaspiMusicServer
	server.player = player.NewPlayer()
	server.m = martini.Classic()
	server.port = fmt.Sprintf(":%d", port)
	server.stopTimer = time.NewTimer(time.Hour * duration)
	server.registerRoutes()
	server.handleStopTimer()
	return &server
}

func (s *RaspiMusicServer) handleStopTimer() {
	go func() {
		for {
			<-s.stopTimer.C
			s.player.ClearPlayList()
			s.stop()
		}
	}()
}

func (s RaspiMusicServer) registerRoutes() {
	s.m.Get("/songs", s.getSongsList)
	s.m.Get("/playlist", s.playlist)
	s.m.Post("/songs/play", s.playSong)
	s.m.Post("/stop", s.stop)
	s.m.Post("/next", s.next)
	s.m.Post("/prev", s.prev)
}

func (s RaspiMusicServer) playlist() (int, string) {
	attr, err := s.player.PlaylistInfo()
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	songs, err := json.Marshal(attr)
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	return 200, string(songs)
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

func (s *RaspiMusicServer) playSong(w http.ResponseWriter, r *http.Request) (int, string) {
	if s.stopTimer != nil {
		s.stopTimer.Reset(time.Hour * duration)
	}
	path, err := s.player.AddSong(r.FormValue("path"))
	if err != nil {
		return 500, fmt.Sprintf("%v\n", err)
	}
	if !s.player.IsPlaying() {
		s.player.Play()
	}
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

func (s RaspiMusicServer) clear() {
	s.player.ClearPlayList()
}

func (s RaspiMusicServer) Run() {
	http.ListenAndServe(s.port, s.m)
}

func (s RaspiMusicServer) Quit() {
	s.stopTimer.Stop()
	s.stop()
}
