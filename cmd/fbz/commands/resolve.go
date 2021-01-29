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

var resolveContent string
var rejectCase bool

var resolveCmd = &cobra.Command{
	Use: "resolve <Case ID>",

	Short: "Resolve a case",

	Long: `Resolve a case

Given a case ID, resolve that case.

If you do not provide a comment message on the command line (or if it's hard to
reduce your message to a one-liner), omiting that option will drop you to a "vi"
instance for you to write your message.`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Usage: fbz resolve <Case ID>")
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

		if len(commentContent) == 0 {
			if err = getCommentContent(); err != nil {
				return fmt.Errorf("you must provide a comment message")
			}
		}

		c, err := cases.Resolve(caseID, rejectCase, commentContent)
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
	resolveCmd.Flags().StringVarP(&commentContent, "message", "m", "",
		"Optionally express your resolution comment message on the command line.")

	resolveCmd.Flags().BoolVarP(&rejectCase, "reject", "r", false,
		`Optionally reject the case as "Wont Fix"`)

	RootCmd.AddCommand(resolveCmd)
}
