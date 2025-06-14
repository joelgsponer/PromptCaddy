package main

import (
	"fmt"
	"os"

	"github.com/promptcaddy/promptcaddy/internal/cli"
	"github.com/promptcaddy/promptcaddy/internal/server"
	"github.com/spf13/cobra"
)

var (
	promptsDir string
	rootCmd    = &cobra.Command{
		Use:   "promptcaddy",
		Short: "A lightweight prompt management tool with MCP server and CLI capabilities",
		Long: `PromptCaddy is a tool for managing, serving, and executing prompts from Markdown files.
It can run as an MCP server or be used directly from the command line.`,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&promptsDir, "dir", "./prompts", "Directory containing prompt files")
	
	// Add subcommands
	rootCmd.AddCommand(newServeCmd())
	rootCmd.AddCommand(cli.NewListCmd(&promptsDir))
	rootCmd.AddCommand(cli.NewCallCmd(&promptsDir))
}

func newServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the MCP server",
		Long:  `Start the Model Context Protocol server to serve prompts to MCP-compatible hosts.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			srv, err := server.NewMCPServer(promptsDir)
			if err != nil {
				return err
			}
			
			fmt.Fprintln(os.Stderr, "Starting PromptCaddy MCP server...")
			return srv.Start()
		},
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}