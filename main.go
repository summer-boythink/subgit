package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	storeDir     string
	wantKeepDir  string
	wantKeepFile string
	logLevel     string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "subgit <subcommand> [options]",
		Short: "subgit is a gadget for cloning a folder in a git repository",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			l, err := log.ParseLevel(logLevel)
			if err != nil {
				return err
			}
			log.SetLevel(l)
			return nil
		},
	}

	rootCmd.PersistentFlags().StringVarP(&storeDir, "storedir", "", defaultDir(), "set dir to store folder(The default value is the current folder)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log", "info", `Set the logging level ("debug"|"info"|"warn"|"error"|"fatal")`)
	// TODO:Is there any api to provide? The user can only select one of the two flags
	rootCmd.PersistentFlags().StringVarP(&wantKeepDir, "keepdir", "d", "", "Select a subfolder to leave behind the git clone")
	rootCmd.PersistentFlags().StringVarP(&wantKeepFile, "keepfile", "f", "", "Select a file to leave behind the git clone")
	rootCmd.AddCommand(newCloneCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
