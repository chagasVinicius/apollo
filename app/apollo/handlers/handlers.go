package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chagasVinicius/apollo/web"
	"github.com/julienschmidt/httprouter"
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

	jsonStr, err := json.Marshal(newCategories)

	fmt.Fprintln(rw, "Categories: ", string(jsonStr))
}


func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
