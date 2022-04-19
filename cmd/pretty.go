/*
MIT License

Copyright © 2022 Phuc ngovanphuc@hotmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
)

var saveToFile bool

// prettyCmd represents the pretty command
var jsonPrettyCmd = &cobra.Command{
	Use:   "pretty",
	Short: "Format json",
	Run: func(cmd *cobra.Command, args []string) {
		err := formatJson(args)
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	jsonPrettyCmd.Flags().BoolVar(&saveToFile, "w", false, "save result to file")
	jsonCmd.AddCommand(jsonPrettyCmd)
}

func formatJson(args []string) error {
	if len(args) < 1 {
		return errors.New("Enter input json file")
	}

	data, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}

	formatedData := pretty.Pretty(data)
	if saveToFile {
		return os.WriteFile(fmt.Sprintf("%s.fomarted.json", args[0]), formatedData, 0644)
	}

	fmt.Println(string(formatedData))
	return nil
}
