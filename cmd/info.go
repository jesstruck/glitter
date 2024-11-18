/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		entries, err := os.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}

		for _, entry := range entries {
			_, err := git.PlainOpen(entry.Name())
			if err != nil {
				//This is not a git repo
				continue
			}

			stash, err := getStashInfo(entry.Name())
			cobra.CheckErr(err)
			if stash == nil || len(stash) == 0 {
				continue
			}
			// This is a Git repo
			log.Printf("---------------------------------------------------------------------------------------------------\nRepo: ./%s\n", entry.Name())

			fmt.Println(string(stash))
			fmt.Println("")
		}
	},
}

func init() {
	stashCmd.AddCommand(infoCmd)
}

func getStashInfo(folder string) ([]byte, error) {
	// Cache current directory
	rootDir := os.Getenv("PWD")

	// Change the directory to the repo folder
	err := os.Chdir(folder)
	cobra.CheckErr(err)

	out, err := exec.Command("git", "stash", "list").Output()

	// Changes the directory back to the root
	err = os.Chdir(rootDir)
	cobra.CheckErr(err)

	return out, err
}
