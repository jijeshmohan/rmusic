package player

import (
	"code.google.com/p/gompd/mpd"
	"log"
)

type MPDPlayer struct {
	client *mpd.Client
}

func NewPlayer() *MPDPlayer {
	var player MPDPlayer
	conn, err := mpd.Dial("tcp", "localhost:6600")
	if err != nil {
		log.Fatalln(err)
	}
	player.client = conn
	return &player
}

func (p *MPDPlayer) AddSong(path string) {

	err := p.client.Add(path)
	if err != nil {
		log.Println("Unable to add songs")
	}
}
func (p *MPDPlayer) Play() {
	err := p.client.Play(-1)
	if err != nil {
		log.Println("Unable to Play songs")
	}
}
func (p *MPDPlayer) Next() {
	err := p.client.Next()
	if err != nil {
		log.Println("Unable to Play Next song")
	}
}

func (p *MPDPlayer) Prev() {
	err := p.client.Previous()
	if err != nil {
		log.Println("Unable to Play Previous song")
	}
}

func (p *MPDPlayer) Stop() {
	err := p.client.Stop()
	if err != nil {
		log.Println("Unable to Stop playing")
	}
}

func (p *MPDPlayer) Close() {
	p.client.Close()
}
