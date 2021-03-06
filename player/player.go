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
	conn.Update("")
	conn.Clear()
	player.client = conn
	return &player
}

func (p *MPDPlayer) IsPlaying() bool {
	state, _ := p.client.Status()
	return state["state"] == "play"
}

func (p *MPDPlayer) PlaylistInfo() ([]mpd.Attrs, error) {
	return p.client.PlaylistInfo(-1, -1)
}

func (p *MPDPlayer) AddSong(path string) (string, error) {
	err := p.client.Add(path)
	return path, err
}

func (p *MPDPlayer) Play() error {
	return p.client.Play(-1)
}

func (p *MPDPlayer) ClearPlayList() error {
	return p.client.Clear()
}

func (p *MPDPlayer) Next() error {
	return p.client.Next()
}

func (p *MPDPlayer) Prev() error {
	return p.client.Previous()
}

func (p *MPDPlayer) Stop() error {
	return p.client.Stop()
}

func (p *MPDPlayer) Close() error {
	return p.client.Close()
}

func (p *MPDPlayer) Songs() ([]mpd.Attrs, error) {
	return p.client.ListAllInfo("/")
}
