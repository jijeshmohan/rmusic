package song

type SongsCollection string

type Song struct {
	Title string `json:"title"`
}

func (collection SongsCollection) All() []Song {
	return []Song{
		{Title: "Song1"},
		{Title: "Song2"},
	}
}

// chiku
