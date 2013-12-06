package song

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type SongsCollection string

type Song struct {
	Title string `json:"title"`
	Id    string `json:"id"`
	path  string
}

func (collection SongsCollection) All() []Song {
	return allMP3Files(string(collection))
}

func (collection SongsCollection) SongPath(id string) (string, error) {
	songs := allMP3Files(string(collection))
	for _, song := range songs {
		if song.Id == id {
			return song.path, nil
		}
	}
	return "", errors.New("unable to file song")
}

// chiku

func allMP3Files(path string) []Song {
	dir, err := os.Open(path)
	checkErr(err)
	defer dir.Close()
	fi, err := dir.Stat()
	checkErr(err)

	songs := make([]Song, 0)
	if fi.IsDir() {
		fis, err := dir.Readdir(-1) // -1 means return all the FileInfos
		checkErr(err)
		for _, fileinfo := range fis {
			if !fileinfo.IsDir() && filepath.Ext(fileinfo.Name()) == ".mp3" {
				title := removeExtension(fileinfo.Name())
				fileSha := getSha(fileinfo.Name())
				songs = append(songs, Song{Title: title, Id: fileSha, path: fileinfo.Name()})
			}
		}
	}
	return songs
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error %v", err)
	}
}
func getSha(filename string) string {
	h := sha1.New()
	h.Write([]byte(filename))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)

}
func removeExtension(filename string) string {
	var extension = filepath.Ext(filename)
	return filename[0 : len(filename)-len(extension)]
}
