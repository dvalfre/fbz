package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	//"github.com/ess/fbz/pkg/fbz"
	"github.com/ess/fbz/pkg/fbz/http"
)

var caseCategory string

var createCmd = &cobra.Command{
	Use: "create <Project Name> <Area Name> <Title>",

	Short: "Create a new case",

	Long: `Create a new case

Given a project name, area name, a title, and an optional comment,
create a new case.

If you do not provide a comment message on the command line (or if it's hard to
reduce your message to a one-liner), omiting that option will drop you to a "vi"
instance for you to write your message.`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return fmt.Errorf("Usage: fbz create <Project Name> <Area Name> <Title>")
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
		project := args[0]
		if len(project) == 0 {
			return fmt.Errorf("you must provide a project name")
		}

		area := args[1]
		if len(area) == 0 {
			return fmt.Errorf("you must provide an area name")
		}

		title := args[2]
		if len(title) == 0 {
			return fmt.Errorf("you must provide a title")
		}

		d, err := http.NewDriver(
			viper.GetString("url"),
			viper.GetString("token"),
		)

		if err != nil {
			return fmt.Errorf("could not set up API client")
		}

		cases := http.NewCaseService(d)

		if len(commentContent) == 0 {
			if err = getCommentContent(); err != nil {
				return fmt.Errorf("you must provide a comment message")
			}
		}

		c, err := cases.Create(project, area, title, caseCategory, commentContent)
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
	createCmd.Flags().StringVarP(&commentContent, "message", "m", "",
		"Optionally express your comment message on the command line.")

	createCmd.Flags().StringVarP(&caseCategory, "category", "c", "Task",
		"The category for the new case.")

	RootCmd.AddCommand(createCmd)
}
