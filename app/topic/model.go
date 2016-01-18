package topic

import (
	"time"

	"github.com/singnurkar/zing/core"
)

type Topic struct {
	core.Item
	Slug        string      `gorethink:"slug,omitempty" json:"slug"`
	FirstLetter string      `gorethink:"first_letter,omitempty" json:"-"`
	Slot        int64       `gorethink:"slot,omitempty" json:"slot,string"`
	Name        string      `gorethink:"name,omitempty" json:"name"`
	Alt         string      `gorethink:"alt,omitempty" json:"alt"`
	Abbr        string      `gorethink:"abbr,omitempty" json:"abbr"`
	Subscribers int64       `gorethink:"subscriptions,omitempty" json:"subscriptions,string"`
	Posts       int64       `gorethink:"posts,omitempty" json:"nposts,string"`
	Intro       string      `gorethink:"intro,omitempty" json:"intro"`
	Cover       core.Cover  `gorethink:"cover,omitempty" json:"cover,omitempty"`
	Creator     string      `gorethink:"creator_id,omitempty" json:"-"`
	Visible     bool        `gorethink:"visible,omitempty" json:"visible"`
	Status      string      `gorethink:"status,omitempty" json:"status,omitempty"`
	Meta        []core.Meta `gorethink:"meta,omitempty" json:"meta"`
}

func NewTopic() *Topic {
	t := &Topic{}
	t.Status = core.STATUS_DRAFT
	t.Created = time.Now()
	t.Visible = true
	return t
}

func (t *Topic) ID() string {
	return t.Item.ID()
}

func (t *Topic) SetID(id string) error {
	return t.Item.SetID(id)
}

type TopicList map[string]string

type Topics []Topics

func (t Topics) Length() int {
	return len(t)
}

type TopicParent struct {
	Id     string `gorethink:"id,omitempty" json:"-"`
	Topic  string `gorethink:"topic_id,omitempty" json:"topic"`
	Parent string `gorethink:"parent_id,omitempty" json:"parent"`
}

func NewParent(childID, parentID string) *TopicParent {
	return &TopicParent{
		Topic:  childID,
		Parent: parentID,
	}
}

type TopicPost struct {
	Id    string `gorethink:"id,omitempty" json:"-"`
	Topic string `gorethink:"topic_id,omitempty" json:"topic"`
	Post  string `gorethink:"post_id,omitempty" json:"post"`
	Slot  int64  `gorethink:"slot,omitempty" json:"slot,string"`
}

func NewPost(topicID, postID string) *TopicPost {
	return &TopicPost{
		Topic: topicID,
		Post:  postID,
	}
}
