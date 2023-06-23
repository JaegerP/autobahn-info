/*
 * ----------------------------------------------------------------------------
 * "THE BEER-WARE LICENSE" (Revision 42):
 * Philipp Jaeger <p@pj4e.de> wrote this file.  As long as you retain this
 * notice you can do whatever you want with this stuff. If we meet some day,
 * and you think this stuff is worth it, you can buy me a beer in return.
 *                                                               Philipp Jaeger
 * ----------------------------------------------------------------------------
 */
package cmd

import (
	"autobahn/lib"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var (
	restClient *resty.Client = resty.New()
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		result := lib.Roads{}
		request := restClient.
			NewRequest().
			SetHeader("Accept", "application/json").
			SetResult(&result)
		_, err := request.Execute("GET", "https://verkehr.autobahn.de/o/autobahn")
		if err != nil {
			return
		}
		fmt.Println("Found ", len(result.Roads), " roads:")
		for _, r := range result.Roads {
			fmt.Println(r)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
