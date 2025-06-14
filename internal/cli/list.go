package cli

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/promptcaddy/promptcaddy/internal/core"
	"github.com/spf13/cobra"
)

// NewListCmd creates the list command
func NewListCmd(promptsDir *string) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all available prompts",
		Long:  `List all prompts found in the prompts directory with their ID, version, and title.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			prompter := core.NewPrompter(*promptsDir)
			
			if err := prompter.LoadPrompts(); err != nil {
				return fmt.Errorf("failed to load prompts: %w", err)
			}

			prompts := prompter.ListPrompts()
			if len(prompts) == 0 {
				fmt.Println("No prompts found in", *promptsDir)
				return nil
			}

			// Create a tabwriter for formatted output
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tVERSION\tTITLE")
			fmt.Fprintln(w, "──\t───────\t─────")

			for _, prompt := range prompts {
				fmt.Fprintf(w, "%s\t%s\t%s\n", prompt.ID, prompt.Version, prompt.Title)
			}

			return w.Flush()
		},
	}
}