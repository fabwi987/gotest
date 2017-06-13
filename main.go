package main

import (
	"log"
	"net/http"

	"github.com/fabwi987/gotest/models"
)

type Env struct {
	db models.Datastore
}

func main() {

	db, err := models.NewDB("root:trustno1@/test?parseTime=true")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	http.HandleFunc("/meets", env.meetsIndex)
	http.ListenAndServe(":7000", nil)

}

func (env *Env) meetsIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	w.Header().Set("Content type", "text/xml")
	w.Write([]byte("Hello wolrd!"))

}
