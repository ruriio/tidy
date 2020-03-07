package model

import (
	"encoding/json"
	"log"
)

type Meta struct {
	Id       string   `json:"id"`
	Title    string   `json:"title"`
	Actor    string   `json:"actor"`
	Producer string   `json:"producer,omitempty"`
	Series   string   `json:"series"`
	Release  string   `json:"release"`
	Duration string   `json:"duration"`
	Sample   string   `json:"sample"`
	Poster   string   `json:"poster"`
	Images   []string `json:"images"`
	Label    string   `json:"label"`
	Genre    []string `json:"genre"`
	Url      string   `json:"url"`
}

func (meta Meta) Json() string {
	out, err := json.MarshalIndent(meta, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	return string(out)
}
