package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ess/fbz/pkg/fbz/http"
)

var listQuery string

var listCmd = &cobra.Command{
	Use: "list",

	Short: "List cases",

	Long: `List cases

List out all cases in the instance. This can also be filtered by project and area.`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(viper.GetString("token")) == 0 {
			return fmt.Errorf(
				`This operation requires a FogBugz API token..
				
This can be generated on the User Options page on FogBugz and should be listed
as token: in ~/.fbz.yml`,
			)
		}

		if len(viper.GetString("url")) == 0 {
			return fmt.Errorf(
				`This operation requires your FogBugz URL.

This should be listed as url: in ~/.fbz.yml`,
			)
		}

		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		d, err := http.NewDriver(
			viper.GetString("url"),
			viper.GetString("token"),
		)

		if err != nil {
			return fmt.Errorf("could not set up API client")
		}

		cases := http.NewCaseService(d)

		for _, c := range cases.All(listQuery) {
			fmt.Printf(
				"%d - (%s) - %d pts - [%s] - %s\n",
				c.ID,
				c.Priority,
				c.Points,
				c.Status,
				c.Title,
			)
		}

		return nil
	},

	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	listCmd.Flags().StringVarP(&listQuery, "query", "q", "status:open",
		"A custom query for the list search..")

	RootCmd.AddCommand(listCmd)
}
