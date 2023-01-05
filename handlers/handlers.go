package handlers

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func MusicClassifyHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	musicName := p.ByName("name")
	fmt.Fprintln(rw, "the music is", musicName)
}
