package core

import (
	"time"
)

type Account struct {
	Item
	Email    string    `gorethink:"email,omitempty" json:"email,omitempty"`
	Username string    `gorethink:"username,omitempty" json:"username,omitempty"`
	Hash     string    `gorethink:"password,omitempty" json:"password,omitempty"`
	Avatar   Cover     `gorethink:"avatar,omitempty" json:"avatar,omitempty"`
	Kind     string    `gorethink:"kind,omitempty" json:"kind,omitempty"`
	Joined   time.Time `gorethink:"joined,omitempty" json:"joined,omitempty"`
	Visited  time.Time `gorethink:"visited,omitempty" json:"visited,omitempty"`
	Enabled  bool      `gorethink:"enabled,omitempty" json:"-"`
	Verified bool      `gorethink:"verified,omitempty" json:"verified,omitempty"`
}

func NewAccount() *Account {
	a := &Account{}
	a.Created = time.Now()
	return a
}

func (a *Account) ID() string {
	return a.Item.ID()
}

func (a *Account) SetID(id string) error {
	return a.Item.SetID(id)
}

type AccountList map[string]*Account

type Accounts []*Account

func (a Accounts) Length() int {
	return len(a)
}
