package core

import "time"

type Topic struct {
	Item
	Slug        string `gorethink:"slug,omitempty" json:"slug"`
	FirstLetter string `gorethink:"first_letter,omitempty" json:"-"`
	Slot        int64  `gorethink:"slot,omitempty" json:"slot,string"`
	Name        string `gorethink:"name,omitempty" json:"name"`
	Alt         string `gorethink:"alt,omitempty" json:"alt"`
	Abbr        string `gorethink:"abbr,omitempty" json:"abbr"`
	Subscribers int64  `gorethink:"subscriptions,omitempty" json:"subscriptions,string"`
	Intro       string `gorethink:"intro,omitempty" json:"intro"`
	Cover       Cover  `gorethink:"cover,omitempty" json:"cover,omitempty"`
	Parents     Topics `gorethink:"-" json:"parents,omitempty"`
	Posts       Posts  `gorethink:"-" json:"parents,omitempty"`
	Creator     string `gorethink:"creator_id,omitempty" json:"-"`
	Visible     bool   `gorethink:"visible,omitempty" json:"visible"`
	Status      string `gorethink:"status,omitempty" json:"status,omitempty"`
	Meta        []Meta `gorethink:"meta,omitempty" json:"meta"`
}

func NewTopic() *Topic {
	t := &Topic{}
	t.Status = STATUS_DRAFT
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

type TopicList map[string]Topic

type Topics []Topics

func (t Topics) Length() int {
	return len(t)
}

type TopicParent struct {
	Id       string `gorethink:"id,omitempty" json:"-"`
	TopicId  string `gorethink:"topic_id,omitempty" json:"topicId"`
	ParentId string `gorethink:"parent_id,omitempty" json:"parentId"`
}

type TopicParents []TopicParent

func NewTopicParent(topicId, parentId string) *TopicParent {
	return &TopicParent{
		TopicId:  topicId,
		ParentId: parentId,
	}
}

type TopicPost struct {
	Id      string `gorethink:"id,omitempty" json:"-"`
	TopicId string `gorethink:"topic_id,omitempty" json:"topicId"`
	PostId  string `gorethink:"parent_id,omitempty" json:"parentId"`
	Slot    int64  `gorethink:"slot,omitempty" json:"slot,string"`
}

func NewTopicPost(topicId, postId string) *TopicPost {
	return &TopicPost{
		TopicId: topicId,
		PostId:  postId,
	}
}
