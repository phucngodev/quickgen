/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

// encodeCmd represents the encode command
var base64EncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode base64",
	Long:  "Encode base64",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Input string in quote")
		}
		data := []byte(args[0])
		fmt.Println(base64.StdEncoding.EncodeToString([]byte(data)))
	},
}

func init() {
	base64Cmd.AddCommand(base64EncodeCmd)
}
