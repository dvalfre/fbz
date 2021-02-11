package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	//"github.com/ess/fbz/pkg/fbz"
	"github.com/ess/fbz/pkg/fbz/http"

	"github.com/ess/fbz/cmd/fbz/presenters"
)

var startCmd = &cobra.Command{
	Use: "start <Case ID>",

	Short: "Start working on a case",

	Long: `Start working on a case

Given a case ID, mark that case as "In Progress".`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Usage: fbz start <Case ID>")
		}

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

		caseID, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		c, err := cases.Start(caseID)
		if err != nil {
			return err
		}

		presenters.PrintCaseDetails(c)

		return nil
	},

	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	RootCmd.AddCommand(startCmd)
}
