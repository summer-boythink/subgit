package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	storeDir = defaultDir()

	c := &CloneOption{
		url: "https://github.com/summer-boythink/router-toy.git",
	}
	c.run()
	// TODO
}
