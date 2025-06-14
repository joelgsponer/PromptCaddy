package core

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/adrg/frontmatter"
)

// PromptParameter represents a parameter that can be passed to a prompt
type PromptParameter struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
	Required    bool   `yaml:"required"`
}

// Prompt represents a parsed prompt from a markdown file
type Prompt struct {
	ID          string            `yaml:"id"`
	Title       string            `yaml:"title"`
	Description string            `yaml:"description"`
	Version     string            `yaml:"version"`
	Parameters  []PromptParameter `yaml:"parameters"`
	Content     string
	FilePath    string
}

// Prompter manages a collection of prompts
type Prompter struct {
	mu      sync.RWMutex
	prompts map[string]*Prompt
	dir     string
}

// NewPrompter creates a new Prompter instance
func NewPrompter(dir string) *Prompter {
	return &Prompter{
		prompts: make(map[string]*Prompt),
		dir:     dir,
	}
}

// LoadPrompts scans the directory and loads all prompt files
func (p *Prompter) LoadPrompts() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Clear existing prompts
	p.prompts = make(map[string]*Prompt)

	err := filepath.Walk(p.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		prompt, err := p.loadPromptFile(path)
		if err != nil {
			// Log error but continue loading other files
			fmt.Fprintf(os.Stderr, "Warning: failed to load %s: %v\n", path, err)
			return nil
		}

		if prompt.ID == "" {
			fmt.Fprintf(os.Stderr, "Warning: prompt in %s has no ID, skipping\n", path)
			return nil
		}

		p.prompts[prompt.ID] = prompt
		return nil
	})

	return err
}

// loadPromptFile loads a single prompt file
func (p *Prompter) loadPromptFile(path string) (*Prompt, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var prompt Prompt
	content, err := frontmatter.Parse(file, &prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	// Read the remaining content
	contentBytes, err := io.ReadAll(content)
	if err != nil {
		return nil, fmt.Errorf("failed to read content: %w", err)
	}

	prompt.Content = string(contentBytes)
	prompt.FilePath = path

	return &prompt, nil
}

// GetPrompt retrieves a prompt by ID
func (p *Prompter) GetPrompt(id string) (*Prompt, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	prompt, exists := p.prompts[id]
	if !exists {
		return nil, fmt.Errorf("prompt not found: %s", id)
	}

	return prompt, nil
}

// ListPrompts returns all loaded prompts
func (p *Prompter) ListPrompts() []*Prompt {
	p.mu.RLock()
	defer p.mu.RUnlock()

	prompts := make([]*Prompt, 0, len(p.prompts))
	for _, prompt := range p.prompts {
		prompts = append(prompts, prompt)
	}

	return prompts
}

// InterpolatePrompt replaces placeholders in the prompt content with provided values
func (p *Prompter) InterpolatePrompt(prompt *Prompt, values map[string]string) (string, error) {
	content := prompt.Content

	// Replace all parameters
	for _, param := range prompt.Parameters {
		placeholder := fmt.Sprintf("{{%s}}", param.Name)
		value, exists := values[param.Name]
		
		if !exists && param.Required {
			return "", fmt.Errorf("required parameter missing: %s", param.Name)
		}
		
		content = strings.ReplaceAll(content, placeholder, value)
	}

	// Replace special {{selection}} parameter
	if selection, exists := values["selection"]; exists {
		content = strings.ReplaceAll(content, "{{selection}}", selection)
	}

	return content, nil
}