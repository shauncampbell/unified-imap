package memory

import (
	"errors"

	"github.com/emersion/go-imap/backend"
)

const inbox = "INBOX"

// User returns an in memory representation of a user session.
type User struct {
	username  string
	mailboxes map[string]*Mailbox
}

// Username returns the name of the user who owns the mailbox.
func (u *User) Username() string {
	return u.username
}

// ListMailboxes returns a list of mailboxes for that user.
func (u *User) ListMailboxes(subscribed bool) (mailboxes []backend.Mailbox, err error) {
	for _, mailbox := range u.mailboxes {
		if subscribed && !mailbox.Subscribed {
			continue
		}

		mailboxes = append(mailboxes, mailbox)
	}
	return
}

// GetMailbox returns a specified mailbox.
func (u *User) GetMailbox(name string) (mailbox backend.Mailbox, err error) {
	mailbox, ok := u.mailboxes[name]
	if !ok {
		err = errors.New("no such mailbox")
	}
	return
}

// CreateMailbox creates a new mailbox.
func (u *User) CreateMailbox(name string) error {
	if _, ok := u.mailboxes[name]; ok {
		return errors.New("mailbox already exists")
	}

	u.mailboxes[name] = &Mailbox{name: name, user: u}
	return nil
}

// DeleteMailbox deletes a mailbox.
func (u *User) DeleteMailbox(name string) error {
	if name == inbox {
		return errors.New("cannot delete INBOX")
	}
	if _, ok := u.mailboxes[name]; !ok {
		return errors.New("no such mailbox")
	}

	delete(u.mailboxes, name)
	return nil
}

// RenameMailbox renames an existing mailbox.
func (u *User) RenameMailbox(existingName, newName string) error {
	mbox, ok := u.mailboxes[existingName]
	if !ok {
		return errors.New("no such mailbox")
	}

	u.mailboxes[newName] = &Mailbox{
		name:     newName,
		Messages: mbox.Messages,
		user:     u,
	}

	mbox.Messages = nil

	if existingName != inbox {
		delete(u.mailboxes, existingName)
	}

	return nil
}

// Logout logs a user out from the mailbox.
func (u *User) Logout() error {
	return nil
}
