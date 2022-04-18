/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"crypto/md5"
	"fmt"

	"github.com/spf13/cobra"
)

// md5Cmd represents the md5 command
var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "Generate md5 for given string",
	Long:  "Generate md5 for given string",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Input string in quote")
		}
		data := []byte(args[0])
		fmt.Printf("%x", md5.Sum(data))
	},
}

func init() {
	rootCmd.AddCommand(md5Cmd)
}
