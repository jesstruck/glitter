/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"time"

	// "github.com/joho/godotenv"
	"github.com/morikuni/aec"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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

const (
	// The name of our config file, without the file extension because viper supports many different config file languages.
	defaultConfigFilename = "glitter"

	// The environment variable prefix of all environment variables bound to our command line flags.
	// For example, --number is bound to CONFLUENCE_NUMBER.
	envPrefix = "GLITTER"

	// Replace hyphenated flag names with camelCase in the config file
	replaceHyphenWithCamelCase = true
)

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
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initializeConfig(cmd)
	},
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
	rootCmd.PersistentFlags().StringVarP(&flags.Organisation, "organisation", "o", "", "The Github Organisaton")

	rootCmd.MarkPersistentFlagRequired("token")
	rootCmd.MarkPersistentFlagRequired("host")
	rootCmd.MarkPersistentFlagRequired("organisation")
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

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()
	// _ = godotenv.Load() //Why if it's never used?

	// Set the base name of the config file, without the file extension.
	v.SetConfigName(defaultConfigFilename)

	// Set as many paths as you like where viper should look for the
	// config file. We are only looking in the current working directory.
	v.AddConfigPath(".")

	// Attempt to read the config file, gracefully ignoring errors
	// caused by a config file not being found. Return an error
	// if we cannot parse the config file.
	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	// When we bind flags to environment variables expect that the
	// environment variables are prefixed, e.g. a flag like --number
	// binds to an environment variable STING_NUMBER. This helps
	// avoid conflicts.
	v.SetEnvPrefix(envPrefix)

	// Environment variables can't have dashes in them, so bind them to their equivalent
	// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// Bind to environment variables
	// Works great for simple config names, but needs help for names
	// like --favorite-color which we fix in the bindFlags function
	v.AutomaticEnv()

	// Bind the current command's flags to viper
	bindFlags(cmd, v)

	return nil
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := f.Name
		// If using camelCase in the config file, replace hyphens with a camelCased string.
		// Since viper does case-insensitive comparisons, we don't need to bother fixing the case, and only need to remove the hyphens.
		if replaceHyphenWithCamelCase {
			configName = strings.ReplaceAll(f.Name, "-", "")
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
