package player

import (
	"code.google.com/p/gompd/mpd"
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
)

type MPDPlayer struct {
	client   *mpd.Client
	songsMap map[string]string
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
	player.songsMap = make(map[string]string)
	player.collectSongs()
	return &player
}

func getSHA(filename string) string {
	h := sha1.New()
	h.Write([]byte(filename))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (p *MPDPlayer) collectSongs() {
	songs, _ := p.client.GetFiles()
	for _, path := range songs {
		p.songsMap[getSHA(path)] = path
	}
}

func (p *MPDPlayer) AddSong(id string) (string, error) {
	path, present := p.songsMap[id]
	if !present {
		return "", errors.New("Unable to find song")
	}
	err := p.client.Add(path)
	return path, err
}

func (p *MPDPlayer) Play() error {
	return p.client.Play(-1)
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

func (p *MPDPlayer) Songs() map[string]string {
	return p.songsMap
}
