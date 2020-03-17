package model

import (
	"log"
	"tidy/util"
)

type Meta struct {
	Id       string            `json:"id"`
	Title    string            `json:"title"`
	Actor    string            `json:"actor"`
	Producer string            `json:"producer,omitempty"`
	Series   string            `json:"series"`
	Release  string            `json:"release"`
	Duration string            `json:"duration"`
	Sample   string            `json:"sample"`
	Poster   string            `json:"poster"`
	Images   []string          `json:"images"`
	Label    string            `json:"label"`
	Genre    []string          `json:"genre"`
	Url      string            `json:"url"`
	Extras   map[string]string `json:"extras"`
}

func (meta Meta) Json() string {
	return string(meta.Byte())
}

func (meta Meta) Byte() []byte {
	out, err := util.JSONMarshal(meta)
	if err != nil {
		log.Panic(err)
	}
	return out
}
