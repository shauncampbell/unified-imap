// Package source provides source implementations based on configuration
package source

import (
	"fmt"

	"github.com/shauncampbell/unified-imap/internal/source/memory"
	"github.com/shauncampbell/unified-imap/pkg/source"
)

// ForConfiguration produces a source implementation based on the configuration.
func ForConfiguration(username string, cfg *source.Config) (source.Source, error) {
	if cfg.Type == "in_memory" {
		return memory.New(username, cfg.Settings), nil
	}
	return nil, fmt.Errorf("no source called '%s'", cfg.Type)
}
