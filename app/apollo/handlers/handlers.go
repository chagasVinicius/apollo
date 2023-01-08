package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chagasVinicius/apollo/web"
	"github.com/julienschmidt/httprouter"
	"github.com/zmb3/spotify/v2"
)

func AddCategory(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newCategories map[string][]string
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newCategories)
	if err != nil {
		panic(err)
	}

	category := newCategories["categories"][0]

	playlists := web.CategoryPlaylists(category)

	playlistID := spotify.ID(playlists.Playlists.Playlists[0].ID)

	playlistsItems := web.GetPlaylistItems(playlistID)

	item := playlistsItems.Items[0]

	jsonStr, err := json.Marshal(item)

	fmt.Fprintln(rw, "Categories: ", string(jsonStr))
}


func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
