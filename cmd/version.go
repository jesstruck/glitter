/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

// import (
// 	"fmt"

// 	"github.com/spf13/cobra"

// 	"github.com/morikuni/aec"
// )

// var (
// 	buildTime string
// 	version   string
// )

// const binverFigletStr = `
//       _ _ _   _
//       | (_) | | |
//   __ _| |_| |_| |_ ___ _ __
//  / _' | | | __| __/ _ \ '__|
// | (_| | | | |_| ||  __/ |
//  \__, |_|_|\__|\__\___|_|
//   __/ |
//  |___/
// `

// // versionCmd represents the version command
// var versionCmd = &cobra.Command{
// 	Use:   "version",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Printf("Version:\t%s\n", version)
// 		fmt.Printf("Build time:\t%s\n", buildTime)
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(versionCmd)

// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

// func SetVersionInfo(version, commit, date string) {
// 	rootCmd.Version = fmt.Sprintf("%s \n\nVersion: %s\nBuild info: %s from Git SHA %s)", aec.LightGreenF.Apply(binverFigletStr), version, date, commit)
// }
