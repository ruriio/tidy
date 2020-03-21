package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func mkdir(name string) {
	if !exists(name) {
		err := os.MkdirAll(name, os.ModePerm)
		check(err)
	}
}

func mkdirParent(dest string) {
	dir := path.Dir(path.Clean(dest))
	err := os.MkdirAll(dir, 0777)
	check(err)
}

func Move(from string, to string) string {

	dest := to

	if !exists(from) {
		return to
	}

	if isFile(from) && strings.HasSuffix(to, "/") {
		name := path.Base(from)
		dest = to + name
	}

	if IsDirectory(from) && !exists(dest) {
		// rename dir name directly
		mkdirParent(dest)
		err := os.Rename(from, dest)
		check(err)
		return dest
	} else {
		if IsDirectory(from) {
			// prevent dir be moved to same name sub dir
			dest = strings.TrimSuffix(dest, "/")
		}

		dir := path.Dir(dest)
		mkdir(dir)

		name := path.Base(dest)
		file := path.Join(dir, name)

		var ext string

		if isFile(from) {
			ext = path.Ext(name)
		}

		base := strings.TrimSuffix(name, ext)

		count := 1
		for exists(file) {
			name = fmt.Sprintf("%s-%d%s", base, count, ext)
			file = path.Join(dir, name)
			count++
		}

		err := os.Rename(from, file)
		check(err)

		if isFile(file) {
			return path.Dir(file)
		} else {
			return file
		}
	}
}

func Write(file string, data []byte) {
	dir := path.Dir(file)
	mkdir(dir)

	err := ioutil.WriteFile(file, data, 0777)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func exists(name string) bool {
	_, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return err == nil
}

func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func isFile(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.Mode().IsRegular()
}
