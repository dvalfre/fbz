package presenters

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/ess/fbz/pkg/fbz"
)

func PrintCaseDetails(c *fbz.Case) {
	PrintCaseList([]*fbz.Case{c})

	fmt.Printf("\nEvents:\n")

	for _, event := range c.Events {
		fmt.Printf("\nWhen: %s\nWho: %s\nWhat: %s\n", event.CreatedAt, event.Creator, event.Description)

		if len(event.Text) > 0 {
			fmt.Printf("\n%s\n", event.Text)
		} else {
			fmt.Println("\n<no message>")
		}
	}
}

func PrintCaseList(cases []*fbz.Case) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', uint(0))
	fmt.Fprintln(writer, "Case ID\tParent ID\tProject\tArea\tPriority\tAssignee\tEstimate\tStatus\tTitle")
	fmt.Fprintln(writer, "=======\t=========\t=======\t====\t========\t========\t========\t======\t=====")
	for _, c := range cases {
		fmt.Fprintf(
			writer,
			"%d\t%d\t%s\t%s\t%s\t%s\t%d pts\t%s\t%s\n",
			c.ID,
			c.ParentID,
			c.ProjectName,
			c.AreaName,
			c.Priority,
			c.Assignee,
			c.Points,
			c.Status,
			c.Title,
		)
	}

	writer.Flush()
}
