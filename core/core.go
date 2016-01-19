package core

import (
	"errors"
	"time"
)

var (
	STATUS_LIVE    = "LIVE"
	STATUS_DELETED = "DELETED"
	STATUS_DRAFT   = "DRAFT"
)

type FlatList map[string]string

type Cover struct {
	Thumb    string `gorethink:"thumb" json:"thumb"`
	Small    string `gorethink:"small" json:"small"`
	Medium   string `gorethink:"medium" json:"medium"`
	Large    string `gorethink:"medium" json:"large"`
	Original string `gorethink:"original" json:"original"`
}

type Files map[string]string

type Content struct {
	Subheading string    `gorethink:"subheading,omitempty" json:"subheading,omitempty"`
	Body       string    `gorethink:"content" json:"content,omitempty"`                  // HTML embed or rendered Markdown
	Raw        string    `gorethink:"raw,omitempty" json:"raw,omitempty"`                // Somewhere to put the Markdown
	Files      Files     `gorethink:"files,omitempty" json:"files,omitempty"`            // Image, audio or video filepath and handles
	ExtID      string    `gorethink:"external_id,omitempty" json:"externalId,omitempty"` // External ID if it's an embed
	Provider   string    `gorethink:"provider,omitempty" json:"provider,omitempty"`      // The embed's provider
	Source     string    `gorethink:"source,omitempty" json:"source,omitempty"`          // Publisher
	Credit     string    `gorethink:"credit,omitempty" json:"credit,omitempty"`          // Artist or author
	Desc       string    `gorethink:"desc,omitempty" json:"desc,omitempty"`
	Kind       string    `gorethink:"kind,omitempty" json:"kind,omitempty"`
	Created    time.Time `gorethink:"created,omitempty" json:"created,omitempty"`
}

type Edit struct {
	editor string    `gorethink:"account_id" json:"editorId"`
	date   time.Time `gorethink:"date" json:"date"`
}

type Meta struct {
	Key   string      `gorethink:"key" json:"key"`
	Value interface{} `gorethink:"value" json:"value"`
}

type Item struct {
	Id       string    `gorethink:"id,omitempty" json:"id,omitempty"`
	Created  time.Time `gorethink:"created,omitempty" json:"created,omitempty"`
	Modified time.Time `gorethink:"modified,omitempty" json:"modified,omitempty"`
	Locked   bool      `gorethink:"locked,omitempty" json:"locked,omitempty"`
}

func (i *Item) ID() string {
	return i.Id
}

func (i *Item) SetID(id string) error {
	var err error
	if len(i.Id) > 0 {
		err = errors.New("ID is already set")
	} else if len(id) == 0 {
		err = errors.New("Given ID is empty")
	} else {
		i.Id = id
	}
	return err
}

type Model interface {
	ID() string
	SetID(id string) error
}

type Collection interface {
	Length()
}
