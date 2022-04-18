/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "base64 sub command",
	Long:  "base64 sub command",
}

func init() {
	rootCmd.AddCommand(base64Cmd)
}
