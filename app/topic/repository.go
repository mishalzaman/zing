package topic

import (
	"log"

	db "github.com/dancannon/gorethink"

	"github.com/singnurkar/zing/dat"
)

type ParentRepository struct {
	session *db.Session
	table   string
}

func (r *ParentRepository) Session() *db.Session { return r.session }
func (r *ParentRepository) Table() string        { return r.table }

type PostRepository struct {
	session *db.Session
	table   string
}

func (r *PostRepository) Session() *db.Session { return r.session }
func (r *PostRepository) Table() string        { return r.table }

type TopicRepository struct {
	session *db.Session
	table   string
	parents *ParentRepository
	posts   *PostRepository
}

func NewRepository(session *db.Session) *TopicRepository {
	return &TopicRepository{
		session: session,
		table:   "topic",
		parents: &ParentRepository{session, "topic_parent"},
		posts:   &PostRepository{session, "topic_post"},
	}
}

func (r *TopicRepository) Session() *db.Session { return r.session }
func (r *TopicRepository) Table() string        { return r.table }

func (r *TopicRepository) Count() (int, error) {
	resultSet := []int{}

	result, err := dat.Count(r)
	if err != nil {
		log.Printf("Error counting topics", err)
	}

	result.All(&resultSet)
	result.Close()

	return resultSet[0], err
}

func (r *TopicRepository) List(offset, limit int) (TopicList, error) {
	rows, err := dat.List(r, "name", offset, limit)
	if err != nil {
		log.Printf("Error retrieving list of topics %d to %d: %s", offset, limit, err)
	}

	topics := []map[string]string{}
	if err = rows.All(&topics); err != nil {
		log.Printf("Error decoding rows into slice: %s", err)
	}
	rows.Close()

	list := TopicList{}
	for _, v := range topics {
		list[v["id"]] = v["title"]
	}

	return list, err
}

func (r *TopicRepository) All(offset, limit int) (Topics, error) {
	rows, err := dat.All(r, offset, limit)
	if err != nil {
		log.Printf("Error retrieving topics %d to %d: %s", offset, limit, err)
	}

	topicsCol := Topics{}
	if err = rows.All(&topicsCol); err != nil {
		log.Printf("Error decoding rows into slice of topics: %s", err)
	}
	rows.Close()

	return topicsCol, err
}

func (r *TopicRepository) One(id string) (*Topic, error) {
	cursor, err := dat.One(r, id)
	if err != nil {
		log.Printf("Error retrieving topic %s: %s", id, err)
	}
	defer cursor.Close()

	topic := &Topic{}
	cursor.One(topic)

	return topic, err
}

func (r *TopicRepository) Save(topic *Topic) error {
	result, err := dat.Create(r, topic)
	if err != nil {
		log.Printf("Error creating new topic: %s", err)
	}
	topic.SetID(result.GeneratedKeys[0])

	return err
}

func (r *TopicRepository) Update(topic *Topic) error {
	_, err := dat.Update(r, topic.ID(), topic)
	if err != nil {
		log.Printf("Error updating Topic: %s", err)
	}

	return err
}

func (r *TopicRepository) Purge(id string) error {
	result, err := dat.Purge(r, id)
	if err != nil {
		log.Printf("Error removing topic: %s", err)
	}
	defer result.Close()

	return err
}

func (r TopicRepository) AddParents(topicID string, pendingParents []string) error {
	parents := make([]*TopicParent, len(pendingParents))

	for i, v := range pendingParents {
		tParent := NewParent(topicID, v)
		parents[i] = tParent
	}

	_, err := dat.Assoc(r.parents, &parents)
	if err != nil {
		log.Printf("Error saving parents of topic %s", topicID)
	}

	return err
}
