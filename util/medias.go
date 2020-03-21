package util

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

func DownloadMedias(dir string, poster string, sample string, images []string) {

	if len(poster) > 0 {
		download(path.Join(dir, "poster.jpg"), poster)
	}

	if len(sample) > 0 {
		if strings.HasSuffix(sample, ".m3u8") {
			downloadM3u8(path.Join(dir, "sample.mp4"), sample)
		} else {
			download(path.Join(dir, "sample.mp4"), sample)
		}
	}

	for i, url := range images {
		mkdir(path.Join(dir, "images"))
		download(path.Join(dir, "images", fmt.Sprintf("%d.jpg", i)), url)
	}

}

func download(filepath string, url string) {
	out, err := os.Create(filepath)
	if out != nil {
		defer out.Close()
	}

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(out, resp.Body)
}

func downloadM3u8(filepath string, url string) {
	cmd := exec.Command("ffmpeg", "-i", url, "-bsf:a", "aac_adtstoasc",
		"-vcodec", "copy", "-c", "copy", "-crf", "50", filepath)
	//cmd.Stderr = os.Stderr
	//cmd.Stdout = os.Stdout

	log.Println(cmd.String())

	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
