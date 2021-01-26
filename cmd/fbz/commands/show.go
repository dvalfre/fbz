package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	//"github.com/ess/fbz/pkg/fbz"
	"github.com/ess/fbz/pkg/fbz/http"
)

var showCmd = &cobra.Command{
	Use: "show <Case ID>",

	Short: "Show details of a case",

	Long: `Show details of a case

Given a case ID, show the information for the case and its various events.`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Usage: fbz show <Case ID>")
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

		c, err := cases.Get(caseID)
		if err != nil {
			return err
		}

		fmt.Printf("(%d) %s\n", c.ID, c.Title)

		for _, event := range c.Events {
			fmt.Printf("\nWhen: %s\nWho: %s\nWhat: %s\n", event.CreatedAt, event.Creator, event.Description)

			if len(event.Text) > 0 {
				fmt.Printf("\n%s\n", event.Text)
			} else {
				fmt.Println("\n<no message>")
			}
		}

		return nil
	},

	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	RootCmd.AddCommand(showCmd)
}
