// Package credentials provides interfaces for working with credential stores.
package credentials

import (
	"github.com/emersion/go-imap/backend"
)

// Store is an interface for interrogating a credentials store and returning a user session.
type Store interface {
	GetBackendForUser(username string) (backend.User, error)
}
