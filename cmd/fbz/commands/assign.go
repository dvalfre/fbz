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

var assignCmd = &cobra.Command{
	Use: "assign <Case ID> <Person Name>",

	Short: "Assign a case",

	Long: `Assign a case

Given a case ID and a person's name, assign the case to that person.`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("Usage: fbz assign <Case ID> <Person Name>")
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

		person := args[1]

		c, err := cases.Assign(caseID, person)
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
	RootCmd.AddCommand(assignCmd)
}
