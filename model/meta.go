package model

import (
	"encoding/json"
	"log"
)

type Meta struct {
	Id       string
	Title    string
	Actor    string
	Producer string
	Series   string
	Age      string
	Sample   string
	Poster   string
	Images   []string
	Label    string
	Genre    string
}

func (meta Meta) Json() string {
	out, err := json.Marshal(meta)
	if err != nil {
		log.Panic(err)
	}
	return string(out)
}
