/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"time"

	"github.com/carlmjohnson/versioninfo"
	"github.com/techchapter/glitter/cmd"
)

func main() {
	cmd.SetVersionInfo(versioninfo.Version, versioninfo.Revision, versioninfo.LastCommit.Format(time.RFC3339))
	cmd.Execute()
}
