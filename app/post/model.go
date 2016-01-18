package post

import (
	"time"

	"github.com/singnurkar/zing/app/topic"
	"github.com/singnurkar/zing/core"
)

type Post struct {
	core.Item
	Slug        string         `gorethink:"slug,omitempty" json:"slug,omitempty"`
	Title       string         `gorethink:"title,omitempty" json:"title,omitempty"`
	FirstLetter string         `gorethink:"first_letter,omitempty" json:"-"`
	Alt         string         `gorethink:"alt,omitempty" json:"alt,omitempty"`
	Abbr        string         `gorethink:"abbr,omitempty" json:"abbr,omitempty"`
	Summary     string         `gorethink:"summary,omitempty" json:"summary,omitempty"`
	Author      string         `gorethink:"author_id,omitempty" json:"author,omitempty"`
	Cover       core.Cover     `gorethink:"cover,omitempty" json:"cover,omitempty"`
	Content     []core.Content `gorethink:"content,omitempty" json:"content"`
	TopicIds    string         `gorethink:"topic_ids,omitempty" json:"-"`
	Topics      []topic.Topic  `gorethink:"-" json:"topics,omitempty"`
	Upvotes     int64          `gorethink:"-" json:"upvotes,omitempty,string"`
	Downvotes   int64          `gorethink:"-" json:"upvotes,omitempty,string"`
	Published   time.Time      `gorethink:"published,omitempty" json:"published,omitempty"`
	Edits       []core.Edit    `gorethink:"edits,omitempty" json:"edits,omitempty"`
	Visible     bool           `gorethink:"visible,omitempty" json:"visible"`
	Status      string         `gorethink:"status,omitempty" json:"status,omitempty"`
	Meta        []core.Meta    `gorethink:"meta,omitempty" json:"meta"`
}

func NewPost() *Post {
	p := &Post{}
	p.Status = core.STATUS_DRAFT
	p.Created = time.Now()
	return p
}

func (p *Post) ID() string {
	return p.Item.ID()
}

func (p *Post) SetID(id string) error {
	return p.Item.SetID(id)
}

type PostList map[string]string

type Posts []Post

func (p Posts) Length() int {
	return len(p)
}
