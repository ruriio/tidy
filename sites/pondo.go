package sites

import (
	"encoding/json"
	"fmt"
	"log"
	. "tidy/model"
	. "tidy/selector"
)

func Pondo(id string) Site {
	next := Site{
		Url:       fmt.Sprintf("https://www.1pondo.tv/dyn/dla/json/movie_gallery/%s.json", id),
		UserAgent: MobileUserAgent,
		Json:      true,
		Decor:     PondoDecor{},

		Selector: Selector{
			Images: Query("Rows"),
		},
	}

	site := Site{
		Url:       fmt.Sprintf("https://www.1pondo.tv/dyn/phpauto/movie_details/movie_id/%s.json", id),
		WebUrl:    fmt.Sprintf("https://www.1pondo.tv/movies/%s/", id),
		UserAgent: MobileUserAgent,
		Json:      true,

		Selector: Selector{
			Title:    Query("Title"),
			Actor:    Query("Actor"),
			Poster:   Query("ThumbHigh"),
			Producer: Preset("1Pondo"),
			Sample:   Preset(fmt.Sprintf("http://smovie.1pondo.tv/sample/movies/%s/1080p.mp4", id)),
			Series:   Query("Series"),
			Release:  Query("Release"),
			Duration: Query("Duration"),
			Id:       Query("MovieID"),
			Label:    Preset(""),
			Genre:    Query("UCNAME"),
		},
		Next: &next,
	}

	return site
}

type PondoDecor struct {
	Decor
}

type PondoImage struct {
	Img       string
	Filename  string
	Protected bool
}

func (decor PondoDecor) Decorate(meta *Meta) *Meta {
	origin := meta.Images

	if len(origin) > 0 {
		var images []string
		var pondo PondoImage
		for _, s := range origin {
			err := json.Unmarshal([]byte(s), &pondo)
			if err != nil {
				log.Fatal(err)
			}
			if !pondo.Protected {
				img := pondo.Img
				if len(img) == 0 {
					img = pondo.Filename
				}
				images = append(images, "https://www.1pondo.tv/dyn/dla/images/"+img)
			}
		}
		meta.Images = images
	}
	return meta
}
