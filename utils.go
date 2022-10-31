package main

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	if s.IsDir() {
		return true
	} else {
		return false
		// TODO:Determine if the parent folder exists and create a new directory if it does
		// parentPaths := strings.Split(path,"/")
		// parentPaths = parentPaths[:len(parentPaths)-1]
		// path = strings.Join(parentPaths,"/")
	}
}

func defaultDir() string {
	str, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	return str + "/"
}

func convertGitUrlToFolderName(url string) string {
	urls := strings.Split(url, "/")
	name := urls[len(urls)-1:][0]
	name = name[:len(name)-4]
	return name
}

func splitFilename(url string) string {
	urls := strings.Split(url, "/")
	if len(urls) == 0 {
		return url
	}
	name := urls[len(urls)-1:][0]
	return name
}

// Remove the git folder when an error occurs
func rmDir(name string) bool {
	if err := os.RemoveAll(name); err != nil {
		log.Errorln("remove git folder fail")
		return false
	}
	return true
}
