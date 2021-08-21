// Package unified provides a backend for go-imap which allows multiple users from unified sources.
package unified

import (
	"errors"

	"github.com/shauncampbell/unified-imap/pkg/credentials"

	imap "github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend"
	"github.com/shauncampbell/unified-imap/pkg/authenticator"
)

// Unified creates a unified backend for go-imap.
type Unified struct {
	auth  authenticator.Authenticator
	store credentials.Store
	backend.Backend
}

// Login logs a user in with the specified username and password.
func (u *Unified) Login(connInfo *imap.ConnInfo, username, password string) (backend.User, error) {
	if u.auth.Authenticate(username, password) {
		return u.store.GetBackendForUser(username)
	}
	return nil, errors.New("bad username or password")
}

// New creates a new unified backend.
func New(auth authenticator.Authenticator, creds credentials.Store) backend.Backend {
	return &Unified{auth: auth, store: creds}
}
