Product Requirements Document: "PromptCaddy" MCP Server & CLI Tool
Document Version: 1.1

Status: Revised Draft for Review

Author: Product Management

Date: June 14, 2025

1. Vision & Mission
Vision: To create the most frictionless, developer-friendly way to manage, use, and serve a version-controlled library of prompts, both via the Model Context Protocol (MCP) and directly on the command line.

Mission: We will build a lightweight, standalone binary using Go, named PromptCaddy. It will operate in two modes: an MCP server that dynamically serves prompts from Markdown files, and a CLI tool that allows users to list and execute these prompts directly. This empowers any developer to extend their AI toolkit with custom, shareable, and trackable prompts with near-zero configuration.

2. Target Audience & User Personas
Persona: "Alex," the AI-Powered Developer

Role: Software Engineer using both an MCP-compatible host (e.g., Cursor) and the command line.

Needs: A unified system for managing reusable prompts for tasks like generating tests, writing documentation, or creating boilerplate code.

Goals:

Integrate custom prompts directly into their IDE's AI assistant via MCP.

Use the same prompts in their terminal for scripting or quick tasks (e.g., cat main.go | promptcaddy call refactor-go-code > main_refactored.go).

Keep prompts in their project's Git repository.

Add or edit a prompt with the simplicity of editing a Markdown file, with changes reflected instantly.

3. Core Principles & Goals
Simplicity is Paramount: The user will download a single, self-contained binary. No runtime installation (like Node.js) is required.

Version Control is Native: Using .md files in a directory makes Git integration inherent.

Dual Utility (Server & CLI): The tool must function seamlessly as both a background MCP server and an interactive CLI client.

Performance is a Feature: Go provides a fast, lightweight binary with minimal resource consumption. Instant hot-reloading for file changes is a critical requirement.

4. Detailed Technical Implementation
This section outlines the precise technical execution plan for the Go implementation. Adherence to these specifications is mandatory.

4.1. Core Technology Stack
Language: Go. Non-negotiable for creating a performant, cross-platform, single-file binary.

Build Toolchain: Standard Go modules.

4.2. Key External Libraries & Justification
MCP Server Framework: We will develop a lightweight, internal MCP handler conforming to the spec, as a mature Go SDK is not yet available. This avoids external dependencies for the core protocol logic.

Markdown Parsing: github.com/adrg/frontmatter - A robust, well-maintained library for parsing YAML frontmatter and content from files.

File System Watching: github.com/fsnotify/fsnotify - The de-facto standard for efficient, cross-platform file system notifications in Go. Essential for the hot-reloading feature.

Command-Line Interface: github.com/spf13/cobra - The premier library for building powerful and modern CLI applications in Go. It will manage our commands, subcommands, and flags.

4.3. Directory & File Structure
The project repository will be structured as follows:

promptcaddy/
├── cmd/                          # Main application entrypoints
│   └── promptcaddy/
│       └── main.go
├── internal/
│   ├── cli/                      # Logic for the CLI tool commands
│   │   ├── call.go
│   │   └── list.go
│   ├── core/                     # Core logic shared by server and CLI
│   │   ├── prompter.go           # The main engine for finding/parsing prompts
│   │   └── watcher.go            # File watcher implementation
│   └── server/                   # Logic for the MCP server mode
│       └── mcp_handlers.go
├── prompts/                      # Default directory for user-defined prompts
│   ├── generate-tests.md
│   └── explain-code.md
├── go.mod
├── go.sum
└── README.md

4.4. The .md Prompt File Specification
This specification is unchanged and remains the core user-facing artifact. It is language-agnostic.

Format: A YAML frontmatter block followed by the prompt template in Markdown.

Example:

---
id: "testgen.generate-unit-tests.v1"
title: "Generate Unit Tests"
description: "Generates unit tests for the selected code using the specified framework."
version: "1.0.2"
parameters:
  - name: "framework"
    description: "The testing framework to use (e.g., Jest, Pytest, Go testing)."
    type: "string"
    required: true
---
Your task is to act as an expert Software Developer in Test. Write a comprehensive suite of unit tests for the following code selection, using the **{{framework}}** framework. The code to test is:

{{selection}}



4.5. Server & Core Logic Flow
Initialization (main.go):

Use cobra to define the root command and two subcommands: serve and call.

Define a persistent flag --dir to specify the prompts directory, defaulting to ./prompts.

Core Prompt Engine (prompter.go):

This component will be responsible for scanning the directory, parsing .md files using frontmatter, validating the structure, and storing the resulting prompt objects in a thread-safe in-memory map (sync.Map).

It will expose methods like GetPrompt(id string) and ListPrompts().

File Watching (watcher.go):

Uses fsnotify to monitor the prompts directory.

On any create, write, or remove event, it will trigger a full rescan and update of the prompt map in the core engine. This ensures atomicity and simplicity.

MCP Server Mode (serve command, mcp_handlers.go):

The serve command will start the file watcher and initialize the MCP server listening on stdio.

listTools handler: Calls prompter.ListPrompts() and returns the result.

callTool handler: Calls prompter.GetPrompt(), performs the variable interpolation, and returns the final prompt string.

4.6. CLI Tool Implementation
The CLI provides direct access to the prompt engine. It will read files from the --dir path on each invocation.

list subcommand:

Command: promptcaddy list

Action: Scans the directory, parses all prompts, and prints a table to stdout with the ID, Version, and Title of each available prompt.

call subcommand:

Command: promptcaddy call <prompt-id> [flags]

Action:

Finds the prompt by its id.

Dynamically adds flags to the call command based on the parameters defined in the prompt's frontmatter. For example, if a prompt has a framework parameter, the CLI will accept a --framework flag.

Reads input from stdin to populate the special {{selection}} parameter.

Performs the interpolation.

Prints the final, rendered prompt to stdout.

Example Usage: cat main.go | promptcaddy call testgen.generate-unit-tests.v1 --framework "Go testing"

5. Testing Strategy (QA Plan)
Unit Tests (Go testing package):

Exhaustively test the prompt parsing and validation logic in prompter.go using testing.T and table-driven tests.

Test the call and list CLI command logic by mocking stdout and args.

Integration Tests:

Write tests that create a temporary directory with sample .md files.

The test will execute the compiled promptcaddy binary as a subprocess.

It will test both serve (by interacting with it via stdio) and call/list modes, asserting that the stdout matches expectations.

It will modify files in the temp directory while the server is running to confirm hot-reloading works.

6. Coding Style & Linting
Formatter: gofmt is mandatory. Code must be formatted before any commit.

Linter: golangci-lint with a provided .golangci.yml configuration file to enforce a strict set of community-standard style guides and best practices.

Comments: All exported functions, types, and complex logic blocks must have clear, idiomatic GoDoc comments.

7. Issue & Project Management
This section remains unchanged.

8. Future Roadmap (Post-V1)
V1.1:

Add a promptcaddy init command to create a sample prompts directory.

Support for more parameter types (number, boolean) with validation in the CLI.

V2.0:

Implement Streamable HTTP transport for the serve command.

Add a "meta-prompt" or "prompt chaining" feature in the CLI.

This revised plan leverages the strengths of Go to better meet our core principles. The scope has expanded slightly with the CLI feature, but the value delivered to the end-user is significantly higher. Let's proceed with this direction.




WORKFLOW
- First initalize the git repo.
- Add a .gitignore.
- Create a dev branch.
- Checkout the dev branch.
- Start to implement the plan. 
- Commit often the title of the commits should read like a CHANGELOG.
- The body should be detailed and cover the key topics adressed.
- Create and checkout new branches as you see fit.
- Merge them back into dev once done.
