package topic

import (
	"log"

	db "github.com/dancannon/gorethink"

	"github.com/singnurkar/zing/core"
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
		return 0, err
	}
	defer result.Close()

	if err = result.All(&resultSet); err != nil {
		log.Printf("Error decoding result", err)
		return 0, err
	}

	return resultSet[0], err
}

func (r *TopicRepository) List(offset, limit int) (core.FlatList, error) {
	rows, err := dat.List(r, "name", offset, limit)
	if err != nil {
		log.Printf("Error retrieving list of topics %d to %d: %s", offset, limit, err)
		return core.FlatList{}, err
	}
	defer rows.Close()

	topics := []map[string]string{}
	if err = rows.All(&topics); err != nil {
		log.Printf("Error decoding rows into slice: %s", err)
		return core.FlatList{}, err
	}

	list := core.FlatList{}
	for _, v := range topics {
		list[v["id"]] = v["title"]
	}

	return list, err
}

func (r *TopicRepository) All(offset, limit int) (core.Topics, error) {
	rows, err := dat.All(r, offset, limit)
	if err != nil {
		log.Printf("Error retrieving topics %d to %d: %s", offset, limit, err)
		return core.Topics{}, err
	}
	defer rows.Close()

	topicsCol := core.Topics{}
	if err = rows.All(&topicsCol); err != nil {
		log.Printf("Error decoding rows into slice of topics: %s", err)
		return core.Topics{}, err
	}

	return topicsCol, err
}

func (r *TopicRepository) One(id string) (*core.Topic, error) {
	cursor, err := dat.One(r, id)
	if err != nil {
		log.Printf("Error retrieving topic %s: %s", id, err)
		return &core.Topic{}, err
	}
	defer cursor.Close()

	topic := &core.Topic{}
	if err = cursor.One(topic); err != nil {
		log.Printf("Error decoding row into topic")
		return &core.Topic{}, err
	}

	return topic, err
}

func (r *TopicRepository) Save(topic *core.Topic) error {
	result, err := dat.Create(r, topic)
	if err != nil {
		log.Printf("Error creating new topic: %s", err)
	}
	topic.SetID(result.GeneratedKeys[0])

	return err
}

func (r *TopicRepository) Update(topic *core.Topic) error {
	_, err := dat.Update(r, topic.ID(), topic)
	if err != nil {
		log.Printf("Error updating Topic: %s", err)
	}

	return err
}

func (r *TopicRepository) Purge(id string) error {
	result, err := dat.Delete(r, id)
	if err != nil {
		log.Printf("Error removing topic: %s", err)
	}
	defer result.Close()

	return err
}

func (r *TopicRepository) Parents(topicId string) (core.TopicList, error) {
	datas := []struct{ Right *core.Topic }{}
	query := dat.Query{"topic_id": topicId}
	field := "post_id"
	table := "post"

	result, err := dat.Join(r.parents, datas, query, field, table)
	if err != nil {
		log.Printf("Error retrieving parents of topic %s", topicId)
		return core.TopicList{}, err
	}
	defer result.Close()

	if err = result.All(&datas); err != nil {
		log.Printf("Could not read data set: %s", err.Error())
		return core.TopicList{}, err
	}

	parents := core.TopicList{}
	for _, v := range datas {
		parents[v.Right.Id] = v.Right
	}

	return parents, err
}

func (r *TopicRepository) AddParents(topicId string, pending []string) error {
	parents := make([]*core.TopicParent, len(pending))
	for i, v := range pending {
		tParent := core.NewTopicParent(topicId, v)
		parents[i] = tParent
	}

	_, err := dat.Assoc(r.parents, &parents)
	if err != nil {
		log.Printf("Error saving parents of topic %s", topicId)
	}

	return err
}
