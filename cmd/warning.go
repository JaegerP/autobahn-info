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
	"autobahn-info/lib"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	roadId string
)

const (
	urlTemplate = "https://verkehr.autobahn.de/o/autobahn/%s/services/warning"
)

// warningCmd represents the warning command
var warningCmd = &cobra.Command{
	Use:   "warning",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if roadId == "" {
			fmt.Println("Road Id is required. Exiting.")
			os.Exit(1)
		}
		result := lib.Warning{}
		request := restClient.
			NewRequest().
			SetHeader("Accept", "application/json").
			SetResult(&result)
		_, err := request.Execute("GET", fmt.Sprintf(urlTemplate, roadId))
		if err != nil {
			return
		}
		fmt.Println("Found ", len(result.Warning), " warnings:")
		fmt.Println()
		for _, w := range result.Warning {
			fmt.Println(w.Title, ":")
			fmt.Println(w.Description[2], w.Description[3], w.Description[4])
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(warningCmd)
	warningCmd.Flags().StringVarP(&roadId, "road-id", "i", "", "Road Id to check (required)")
}
