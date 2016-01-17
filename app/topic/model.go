package topic

import (
	"time"

	"github.com/singnurkar/zing/core"
)

type TopicList map[string]string

type Subscription struct {
	Item    string    `gorethink:"item_id,omitempty" json:"itemId"`
	Account string    `gorethink:"account_id,omitempty" json:"accountId"`
	Date    time.Time `gorethink:"date" json:"date"`
}

type Topic struct {
	Id            string      `gorethink:"id,omitempty" json:"id,omitempty"`
	Slug          string      `gorethink:"slug,omitempty" json:"slug"`
	FirstLetter   string      `gorethink:"first_letter,omitempty" json:"-"`
	Slot          int64       `gorethink:"slot,omitempty" json:"slot,string"`
	Name          string      `gorethink:"name,omitempty" json:"name"`
	Alt           string      `gorethink:"alt,omitempty" json:"alt"`
	Abbr          string      `gorethink:"abbr,omitempty" json:"abbr"`
	Subscriptions int64       `gorethink:"subscriptions,omitempty" json:"subscriptions,string"`
	Posts         int64       `gorethink:"posts,omitempty" json:"nposts,string"`
	Intro         string      `gorethink:"intro,omitempty" json:"intro"`
	Cover         core.Cover  `gorethink:"cover,omitempty" json:"cover,omitempty"`
	Created       time.Time   `gorethink:"created,omitempty" json:"created"`
	Modified      time.Time   `gorethink:"modified,omitempty" json:"modified"`
	Creator       string      `gorethink:"creator_id,omitempty" json:"-"`
	Visible       bool        `gorethink:"visible,omitempty" json:"visible"`
	Status        string      `gorethink:"status,omitempty" json:"status,omitempty"`
	Locked        bool        `gorethink:"locked,omitempty" json:"locked"`
	Meta          []core.Meta `gorethink:"meta,omitempty" json:"meta"`
}

type Topics []Topic

type TopicParent struct {
	Id     string `gorethink:"id,omitempty" json:"-"`
	Topic  string `gorethink:"topic_id,omitempty" json:"topic"`
	Parent string `gorethink:"parent_id,omitempty" json:"parent"`
}

type TopicPost struct {
	Id    string `gorethink:"id,omitempty" json:"-"`
	Topic string `gorethink:"topic_id,omitempty" json:"topic"`
	Post  string `gorethink:"post_id,omitempty" json:"post"`
	Slot  int64  `gorethink:"slot,omitempty" json:"slot,string"`
}

func NewTopic() *Topic {
	return &Topic{
		Visible: true,
		Status:  core.STATUS_DRAFT,
		Created: time.Now(),
	}
}

func (t *Topic) Subscribe(accountID string) *Subscription {
	return NewSubscription(t.Id, accountID)
}

func (t *Topic) Unsubscribe(accountID string) {

}

func NewTopicParent(childID, parentID string) *TopicParent {
	return &TopicParent{
		Topic:  childID,
		Parent: parentID,
	}
}

func NewTopicPost(topicID, postID string) *TopicPost {
	return &TopicPost{
		Topic: topicID,
		Post:  postID,
	}
}

func NewSubscription(itemID, accountID string) *Subscription {
	return &Subscription{
		Item:    itemID,
		Account: accountID,
		Date:    time.Now(),
	}
}
