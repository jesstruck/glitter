/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/morikuni/aec"

	"github.com/spf13/cobra"
)

var flags Flags
var (
	// Version will be the version tag if the binary is built with "go install url/tool@version".
	// If the binary is built some other way, it will be "(devel)".
	version = "unknown"
	// Revision is taken from the vcs.revision tag in Go 1.18+.
	revision = "unknown"
	// LastCommit is taken from the vcs.time tag in Go 1.18+.
	lastCommit time.Time
	// DirtyBuild is taken from the vcs.modified tag in Go 1.18+.
	dirtyBuild = true
)

const glitterFigletStr = `
      _ _ _   _            
      | (_) | | |           
  __ _| |_| |_| |_ ___ _ __ 
 / _' | | | __| __/ _ \ '__|
| (_| | | | |_| ||  __/ |   
 \__, |_|_|\__|\__\___|_|   
  __/ |                     
 |___/  
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "glitter",
	Version: version,

	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	info()
	rootCmd.SetVersionTemplate(fmt.Sprintf("%s \nVersion: {{.Version}}\nTime: %s\nSHA: %s\nDirty: %t\n\n",
		aec.LightGreenF.Apply(glitterFigletStr),
		&lastCommit,
		revision,
		dirtyBuild))
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.glitter.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVarP(&flags.Token, "token", "t", "", "Your personal Github token")
	rootCmd.PersistentFlags().StringVarP(&flags.Host, "host", "g", "", "If you have a custom Github host")
	rootCmd.PersistentFlags().StringVarP(&flags.Organisation, "organisaton", "o", "", "The Github Organisaton")
}

func info() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	if info.Main.Version != "" {
		version = info.Main.Version
	}
	for _, kv := range info.Settings {
		if kv.Value == "" {
			continue
		}
		switch kv.Key {
		case "vcs.revision":
			revision = kv.Value
		case "vcs.time":
			lastCommit, _ = time.Parse(time.RFC3339, kv.Value)
		case "vcs.modified":
			dirtyBuild = kv.Value == "true"
		}
	}

}
