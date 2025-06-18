# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

PromptCaddy is a Go-based prompt management tool that operates in dual modes:
- **MCP Server**: Serves prompts to MCP-compatible IDEs via JSON-RPC 2.0 protocol
- **CLI Tool**: Direct command-line prompt execution with piping support

## Architecture

### Core Components
- **`internal/core/prompter.go`**: Thread-safe prompt management engine with concurrent access controls
- **`internal/core/watcher.go`**: File system monitoring for hot-reload functionality  
- **`internal/server/mcp_handlers.go`**: MCP protocol implementation (tools/list, tools/call)
- **`internal/cli/`**: CLI command implementations for list and call operations
- **`cmd/promptcaddy/main.go`**: Application entry point with Cobra CLI framework

### Prompt File Format
Prompts are Markdown files with YAML frontmatter stored in `prompts/` directory:
```yaml
---
id: "unique.prompt.id"
title: "Display Title"
description: "Brief description"
parameters:
  - name: "param_name"
    description: "Parameter description"
    type: "string"
    required: true
---
Prompt content with {{param_name}} and {{selection}} placeholders.
```

### Concurrency Design
- Uses `sync.RWMutex` for thread-safe prompt storage access
- File watcher runs in separate goroutine with event handling
- MCP server handles concurrent client connections safely

## Development Commands

### Build and Run
```bash
# Build application
go build ./cmd/promptcaddy

# Build with specific output name
go build -o promptcaddy ./cmd/promptcaddy

# Run MCP server (development)
./promptcaddy serve --dir ./prompts

# Test CLI functionality
./promptcaddy list --dir ./prompts
./promptcaddy call <prompt-id> --parameter value --dir ./prompts

# Pipe input to prompts
cat source.go | ./promptcaddy call explain-code --audience "beginner"
```

### Testing
Currently no automated test suite exists. Manual testing involves:
- Running both MCP server and CLI modes
- Testing file watcher with prompt modifications
- Verifying parameter injection and template processing

## Key Implementation Details

### Hot-Reload Architecture
- File watcher monitors `.md` files in prompts directory
- Automatic prompt map updates on file create/modify/delete events
- Changes are immediately available in both server and CLI modes

### Parameter Handling
- CLI dynamically generates flags based on prompt frontmatter parameters
- MCP server accepts parameters via JSON-RPC arguments
- Template interpolation replaces `{{parameter}}` placeholders in prompt content
- Special `{{selection}}` parameter populated from stdin or MCP selection

### MCP Protocol Implementation
- Custom JSON-RPC 2.0 server implementation
- Advertises `tools` capability with list and call operations
- Follows MCP specification for tool discovery and execution
- Handles both sync and async tool calls appropriately