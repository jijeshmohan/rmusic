package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"raspimusic/web"
)

var (
	port      = flag.Int("port", 8080, "port of the webserver")
	songsPath = flag.String("path", "/home/pi/musics", "Music files path")
)

func main() {
	log.Println("Starting Raspimusic")
	flag.Parse()

	server := web.NewRaspiMusicServer(*port, *songsPath)
	server.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
}
