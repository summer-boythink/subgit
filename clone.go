package main

import (
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CloneOption struct {
	url string
}

func newCloneCmd() *cobra.Command {
	var opts CloneOption
	cmd := &cobra.Command{
		Use:   "clone <repository>",
		Short: "clone a git repo",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.url = args[0]
			return opts.run()
		},
	}
	return cmd
}

func (c *CloneOption) run() error {
	if wantKeepDir != "" && wantKeepFile != "" {
		log.Errorln("Only one of wantKeepDir and wantKeepFile can be non-empty")
		return nil
	}

	// Check whether storeDir is valid
	if !isDir(storeDir) {
		log.Errorln("storeDir is not a dir")
		return nil
	}

	// git clone
	_, err := exec.Command("git", "clone", c.url).Output()
	if err != nil {
		log.Errorln("maybe git path is invalid or path is exists")
		return nil
	}

	// move folder
	folderName := convertGitUrlToFolderName(c.url)
	// TODO:Formatting Directory
	sourceName := defaultDir() + folderName
	destName := storeDir + folderName

	if storeDir != defaultDir() && wantKeepDir == "" && wantKeepFile == "" {
		err := os.Rename(sourceName, destName)
		if err != nil {
			log.Errorln(sourceName, "not exists, or", destName, "has exists")
			rmDir(sourceName)
			return nil
		}
	}
	log.Infoln(destName, "is created")

	// keep a file or dir
	if wantKeepDir != "" {
		//TODO:check dir

		newSourceDir := sourceName + "/" + wantKeepDir
		destDir := storeDir + wantKeepDir
		if err := os.Rename(newSourceDir, destDir); err != nil {
			log.Errorln("keep dir failed")
		}
		log.Infoln(destDir, "is created")
		log.Warnln(destName, "will be removed")
		rmDir(sourceName)
	}

	if wantKeepFile != "" {
		//TODO:check file

		newSourceFile := sourceName + "/" + wantKeepFile
		destFile := storeDir + filepath.Base(wantKeepFile)
		if err := os.Rename(newSourceFile, destFile); err != nil {
			log.Errorln("keep file failed")
		}
		log.Infoln(destFile, "is created")
		log.Warnln(destName, "will be removed")
		rmDir(sourceName)
	}

	return nil
}
