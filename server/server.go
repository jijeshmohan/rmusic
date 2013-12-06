package server

import (
	"fmt"
	"github.com/codegangsta/martini"
	"log"
	"net/http"
	"raspimusic/song"
	"strings"
)

type RaspiMusicServer struct {
	m    *martini.ClassicMartini
	port string
}

func NewRaspiMusicServer(port int) *RaspiMusicServer {
	var server RaspiMusicServer

	server.m = martini.Classic()
	server.port = fmt.Sprintf(":%d", port)
	server.registerRoutes()
	return &server
}

func (server RaspiMusicServer) registerRoutes() {
	server.m.Get("/songs", server.getSongsList)
}

func (server RaspiMusicServer) getSongsList() string {
	log.Println("Songs list")
	return strings.Join(song.All(), ", ")
}

func (server RaspiMusicServer) Run() {
	http.ListenAndServe(server.port, server.m)
}
