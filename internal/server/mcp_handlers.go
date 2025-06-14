package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/promptcaddy/promptcaddy/internal/core"
)

// MCPServer implements the Model Context Protocol server
type MCPServer struct {
	prompter *core.Prompter
	watcher  *core.Watcher
	reader   *bufio.Reader
	writer   io.Writer
}

// NewMCPServer creates a new MCP server instance
func NewMCPServer(promptsDir string) (*MCPServer, error) {
	prompter := core.NewPrompter(promptsDir)
	
	// Load initial prompts
	if err := prompter.LoadPrompts(); err != nil {
		return nil, fmt.Errorf("failed to load prompts: %w", err)
	}

	// Set up file watcher
	watcher, err := core.NewWatcher(prompter)
	if err != nil {
		return nil, fmt.Errorf("failed to create watcher: %w", err)
	}

	return &MCPServer{
		prompter: prompter,
		watcher:  watcher,
		reader:   bufio.NewReader(os.Stdin),
		writer:   os.Stdout,
	}, nil
}

// MCPRequest represents an incoming MCP request
type MCPRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

// MCPResponse represents an outgoing MCP response
type MCPResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Result  interface{}     `json:"result,omitempty"`
	Error   *MCPError       `json:"error,omitempty"`
}

// MCPError represents an MCP error
type MCPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ToolInfo represents information about a tool
type ToolInfo struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	InputSchema map[string]interface{} `json:"inputSchema"`
}

// Start begins the MCP server
func (s *MCPServer) Start() error {
	// Start file watcher
	if err := s.watcher.Start(); err != nil {
		return err
	}
	defer s.watcher.Stop()

	// Send initialize response
	s.sendInitialize()

	// Main message loop
	for {
		line, err := s.reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		var req MCPRequest
		if err := json.Unmarshal(line, &req); err != nil {
			continue // Skip malformed messages
		}

		s.handleRequest(&req)
	}
}

// sendInitialize sends the server initialization info
func (s *MCPServer) sendInitialize() {
	init := map[string]interface{}{
		"protocolVersion": "2024-11-05",
		"capabilities": map[string]interface{}{
			"tools": map[string]interface{}{},
		},
		"serverInfo": map[string]interface{}{
			"name":    "promptcaddy",
			"version": "1.0.0",
		},
	}

	response := MCPResponse{
		JSONRPC: "2.0",
		ID:      json.RawMessage("0"),
		Result:  init,
	}

	s.sendResponse(&response)
}

// handleRequest processes an incoming MCP request
func (s *MCPServer) handleRequest(req *MCPRequest) {
	switch req.Method {
	case "initialize":
		s.handleInitialize(req)
	case "tools/list":
		s.handleListTools(req)
	case "tools/call":
		s.handleCallTool(req)
	default:
		s.sendError(req.ID, -32601, "Method not found")
	}
}

// handleInitialize handles the initialize request
func (s *MCPServer) handleInitialize(req *MCPRequest) {
	result := map[string]interface{}{
		"protocolVersion": "2024-11-05",
		"capabilities": map[string]interface{}{
			"tools": map[string]interface{}{},
		},
		"serverInfo": map[string]interface{}{
			"name":    "promptcaddy",
			"version": "1.0.0",
		},
	}

	s.sendResponse(&MCPResponse{
		JSONRPC: "2.0",
		ID:      req.ID,
		Result:  result,
	})
}

// handleListTools handles the tools/list request
func (s *MCPServer) handleListTools(req *MCPRequest) {
	prompts := s.prompter.ListPrompts()
	tools := make([]ToolInfo, 0, len(prompts))

	for _, prompt := range prompts {
		// Build input schema
		properties := make(map[string]interface{})
		required := make([]string, 0)

		for _, param := range prompt.Parameters {
			properties[param.Name] = map[string]interface{}{
				"type":        param.Type,
				"description": param.Description,
			}
			if param.Required {
				required = append(required, param.Name)
			}
		}

		// Add selection parameter
		properties["selection"] = map[string]interface{}{
			"type":        "string",
			"description": "The selected text or code to process",
		}

		inputSchema := map[string]interface{}{
			"type":       "object",
			"properties": properties,
			"required":   required,
		}

		tools = append(tools, ToolInfo{
			Name:        prompt.ID,
			Description: prompt.Description,
			InputSchema: inputSchema,
		})
	}

	result := map[string]interface{}{
		"tools": tools,
	}

	s.sendResponse(&MCPResponse{
		JSONRPC: "2.0",
		ID:      req.ID,
		Result:  result,
	})
}

// CallToolParams represents parameters for tools/call
type CallToolParams struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

// handleCallTool handles the tools/call request
func (s *MCPServer) handleCallTool(req *MCPRequest) {
	var params CallToolParams
	if err := json.Unmarshal(req.Params, &params); err != nil {
		s.sendError(req.ID, -32602, "Invalid params")
		return
	}

	prompt, err := s.prompter.GetPrompt(params.Name)
	if err != nil {
		s.sendError(req.ID, -32602, "Tool not found")
		return
	}

	// Convert arguments to string map
	values := make(map[string]string)
	for k, v := range params.Arguments {
		values[k] = fmt.Sprintf("%v", v)
	}

	// Interpolate the prompt
	result, err := s.prompter.InterpolatePrompt(prompt, values)
	if err != nil {
		s.sendError(req.ID, -32602, err.Error())
		return
	}

	s.sendResponse(&MCPResponse{
		JSONRPC: "2.0",
		ID:      req.ID,
		Result: map[string]interface{}{
			"content": []map[string]string{
				{
					"type": "text",
					"text": result,
				},
			},
		},
	})
}

// sendResponse sends a response to the client
func (s *MCPServer) sendResponse(resp *MCPResponse) {
	data, _ := json.Marshal(resp)
	fmt.Fprintf(s.writer, "%s\n", data)
}

// sendError sends an error response
func (s *MCPServer) sendError(id json.RawMessage, code int, message string) {
	s.sendResponse(&MCPResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error: &MCPError{
			Code:    code,
			Message: message,
		},
	})
}