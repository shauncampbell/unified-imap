package config

import (
	"github.com/shauncampbell/unified-imap/internal/credentials/memory"
)

// Credentials provides configuration for a credentials source.
type Credentials struct {
	Type     string        `yaml:"type"`
	InMemory memory.Config `yaml:"in_memory"`
}
