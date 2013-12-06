package song

import (
	"log"
	"os"
	"path/filepath"
)

type SongsCollection string

type Song struct {
	Title string `json:"title"`
}

func (collection SongsCollection) All() []Song {
	return allMP3Files(string(collection))
}

// chiku

func allMP3Files(path string) []Song {
	dir, err := os.Open(path)
	checkErr(err)
	defer dir.Close()
	fi, err := dir.Stat()
	checkErr(err)

	filenames := make([]Song, 0)
	if fi.IsDir() {
		fis, err := dir.Readdir(-1) // -1 means return all the FileInfos
		checkErr(err)
		for _, fileinfo := range fis {
			if !fileinfo.IsDir() && filepath.Ext(fileinfo.Name()) == ".mp3" {
				title := removeExtension(fileinfo.Name())
				filenames = append(filenames, Song{Title: title})
			}
		}
	}
	return filenames
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error %v", err)
	}
}

func removeExtension(filename string) string {
	var extension = filepath.Ext(filename)
	return filename[0 : len(filename)-len(extension)]
}
