package core

import (
	"fmt"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// Watcher monitors a directory for changes and triggers prompt reloads
type Watcher struct {
	watcher  *fsnotify.Watcher
	prompter *Prompter
	done     chan bool
}

// NewWatcher creates a new file system watcher
func NewWatcher(prompter *Prompter) (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Watcher{
		watcher:  watcher,
		prompter: prompter,
		done:     make(chan bool),
	}, nil
}

// Start begins watching the prompts directory
func (w *Watcher) Start() error {
	err := w.watcher.Add(w.prompter.dir)
	if err != nil {
		return fmt.Errorf("failed to watch directory: %w", err)
	}

	go w.watch()
	return nil
}

// Stop stops the watcher
func (w *Watcher) Stop() {
	close(w.done)
	w.watcher.Close()
}

// watch handles file system events
func (w *Watcher) watch() {
	for {
		select {
		case event, ok := <-w.watcher.Events:
			if !ok {
				return
			}

			// Only care about .md files
			if filepath.Ext(event.Name) != ".md" {
				continue
			}

			// Reload prompts on any relevant event
			if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove|fsnotify.Rename) != 0 {
				if err := w.prompter.LoadPrompts(); err != nil {
					fmt.Printf("Error reloading prompts: %v\n", err)
				} else {
					fmt.Printf("Reloaded prompts after change to: %s\n", event.Name)
				}
			}

		case err, ok := <-w.watcher.Errors:
			if !ok {
				return
			}
			fmt.Printf("Watcher error: %v\n", err)

		case <-w.done:
			return
		}
	}
}