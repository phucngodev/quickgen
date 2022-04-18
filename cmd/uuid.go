/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate uuid code",
	Long:  "Generate uuid code",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uuid called")
		fmt.Println(uuid.NewString())
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
}
