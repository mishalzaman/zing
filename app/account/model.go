package account

import (
	"time"

	"github.com/singnurkar/zing/core"
)

type Account struct {
	core.Item
	Email    string `gorethink:"email,omitempty" json:"email,omitempty"`
	Username string `gorethink:"username,omitempty" json:"username,omitempty"`
	Hash     string `gorethink:"password,omitempty" json:"password,omitempty"`
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

type Accounts []Account

func (a Accounts) Length() int {
	return len(a)
}
