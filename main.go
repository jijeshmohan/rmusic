package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"raspimusic/server"
)

var (
	port = flag.Int("port", 8080, "port of the webserver")
)

func main() {
	log.Println("Starting Raspimusic")
	flag.Parse()

	server.NewRaspiMusicServer(*port).Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
}
