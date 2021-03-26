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

var reparentCmd = &cobra.Command{
	Use: "reparent <Case ID> <Parent ID>",

	Short: "Change the parent case for a case",

	Long: `Change the parent case for a case

Given a case ID and a parent case ID, update the case with said parent.`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("Usage: fbz reparent <Case ID> <Parent ID>")
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

		parentID, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		c, err := cases.Reparent(caseID, parentID)
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
	RootCmd.AddCommand(reparentCmd)
}
