package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/promptcaddy/promptcaddy/internal/core"
	"github.com/spf13/cobra"
)

// NewCallCmd creates the call command
func NewCallCmd(promptsDir *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "call <prompt-id> [flags]",
		Short: "Execute a prompt with the given parameters",
		Long: `Execute a prompt by its ID, providing parameters via flags.
Input from stdin will be used as the {{selection}} parameter.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			promptID := args[0]
			
			prompter := core.NewPrompter(*promptsDir)
			if err := prompter.LoadPrompts(); err != nil {
				return fmt.Errorf("failed to load prompts: %w", err)
			}

			prompt, err := prompter.GetPrompt(promptID)
			if err != nil {
				return err
			}

			// Collect parameter values from flags
			values := make(map[string]string)
			
			// Read stdin for {{selection}}
			if !isTerminal(os.Stdin) {
				input, err := io.ReadAll(os.Stdin)
				if err != nil {
					return fmt.Errorf("failed to read stdin: %w", err)
				}
				values["selection"] = string(input)
			}

			// Get values from flags
			for _, param := range prompt.Parameters {
				value, err := cmd.Flags().GetString(param.Name)
				if err == nil && value != "" {
					values[param.Name] = value
				}
			}

			// Interpolate and output
			result, err := prompter.InterpolatePrompt(prompt, values)
			if err != nil {
				return err
			}

			fmt.Print(result)
			return nil
		},
	}

	// This will be populated dynamically before execution
	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return nil
		}

		promptID := args[0]
		prompter := core.NewPrompter(*promptsDir)
		
		if err := prompter.LoadPrompts(); err != nil {
			return fmt.Errorf("failed to load prompts: %w", err)
		}

		prompt, err := prompter.GetPrompt(promptID)
		if err != nil {
			return err
		}

		// Add flags for each parameter
		for _, param := range prompt.Parameters {
			cmd.Flags().String(param.Name, "", param.Description)
			if param.Required {
				cmd.MarkFlagRequired(param.Name)
			}
		}

		return nil
	}

	return cmd
}

// isTerminal checks if a file descriptor is a terminal
func isTerminal(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}