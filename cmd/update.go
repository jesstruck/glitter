/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/google/go-github/v66/github"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("update called")
		client := NewClient(flags.Token, flags.Host)
		orgs, _, err := client.Repositories.ListByOrg(context.Background(), flags.Organisation, nil)
		if err != nil {
			log.Println(err)
		}

		for _, org := range orgs {
			log.Println(fmt.Sprintf("Working on %s: %s", *org.Name, *org.SSHURL))
			// Check if the path is a directory
			path, err := os.Stat(*org.Name)
			if err != nil {
				log.Println(err)
			}
			if stat, err := os.Stat(path.Name()); err == nil && stat.IsDir() {
				// git stash inside the directory
				stash(*org.Name)
				// git pull inside the directory
				r, err := git.PlainOpen(*org.Name)
				cobra.CheckErr(err)
				w, err := r.Worktree()
				cobra.CheckErr(err)
				err = w.Pull(&git.PullOptions{
					RemoteName: "origin",
					Progress:   os.Stdout,
				})
				if err == git.NoErrAlreadyUpToDate {
					log.Println("Already up to date")
					continue
				}
				if err == git.ErrFastForwardMergeNotPossible {
					log.Println("Fast forward merge not possible")
					continue
				}

				cobra.CheckErr(err)

				// Print the latest commit that was just pulled
				ref, err := r.Head()
				cobra.CheckErr(err)
				commit, err := r.CommitObject(ref.Hash())
				cobra.CheckErr(err)

				log.Println(commit)

			}
			clone(*org.SSHURL)

		}

		log.Printf("\n\nDone\n\n")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func NewClient(token, host string) *github.Client {
	ctx := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tokenClient := oauth2.NewClient(ctx, tokenService)

	url := fmt.Sprintf("https://%s/api/v3", host)

	client, err := github.NewClient(tokenClient).WithEnterpriseURLs(url, url)
	if err != nil {
		panic(err)
	}
	return client
}

func stash(name string) {
	log.Println(fmt.Sprintf("Stashing %s", name))
	//Get the current directory
	rootDir := os.Getenv("PWD")
	// Change the directory to the repo folder
	err := os.Chdir(name)
	if err != nil {
		log.Println(err)
	}

	// Stash' eventual changes in the repository
	err = exec.Command("git", "stash", "--message", fmt.Sprintf("Stashed at %s for update", time.Now().Format("2006-01-02 15:04:05"))).Run()
	if err != nil {
		log.Println(err)
	}

	// Changes the directory back to the root
	err = os.Chdir(rootDir)
	if err != nil {
		log.Println(err)
	}
}

func clone(sshUrl string) {
	log.Println("Cloning")
	authMethod, err := ssh.DefaultAuthBuilder("Keymaster")
	if err != nil {
		log.Println(err)
	}

	_, err = git.PlainClone(".", false, &git.CloneOptions{
		URL:      sshUrl,
		Progress: os.Stdout,
		Auth:     authMethod,
	})
	if err != nil {
		log.Println(err)
	}
}
