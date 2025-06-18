# PromptCaddy

A lightweight prompt management tool with MCP server and CLI capabilities.

## Overview

PromptCaddy is a Go-based tool for managing, serving, and executing prompts from Markdown files. It can run as a Model Context Protocol (MCP) server or be used directly from the command line.

## Features

- MCP server for serving prompts to MCP-compatible hosts
- CLI for listing and calling prompts directly
- Markdown-based prompt storage with frontmatter metadata
- File watching for automatic prompt updates
- Structured prompt management and organization

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/PromptCaddy.git
cd PromptCaddy

# Build the application
go build ./cmd/promptcaddy
```

## Usage

### MCP Server Mode
```bash
# Start the MCP server
./promptcaddy serve --dir ./prompts
```

### CLI Mode
```bash
# List available prompts
./promptcaddy list --dir ./prompts

# Call a specific prompt
./promptcaddy call <prompt-name> --dir ./prompts
```

## Project Structure

```
PromptCaddy/
├── cmd/
│   └── promptcaddy/    # Main application entry point
├── internal/
│   ├── cli/            # CLI commands (list, call)
│   ├── core/           # Core functionality (prompter, watcher)
│   └── server/         # MCP server implementation
├── prompts/            # Prompt files in Markdown format
├── go.mod              # Go module definition
└── README.md           # This file
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.