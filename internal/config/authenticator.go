package config

import (
	"github.com/shauncampbell/unified-imap/internal/authenticator/memory"
)

// Authenticator is a struct providing configuration of an authenticator.
type Authenticator struct {
	Type     string        `yaml:"type"`
	InMemory memory.Config `yaml:"in_memory"`
}
