package topic

import (
	"log"

	db "github.com/dancannon/gorethink"
	"github.com/singnurkar/zing/dat"
)

type TopicRepository struct {
	session *db.Session
	table   string
}

func NewRepository(session *db.Session) *TopicRepository {
	return &TopicRepository{session: session, table: "topic"}
}

func (r *TopicRepository) Session() *db.Session { return r.session }
func (r *TopicRepository) Table() string        { return r.table }

func (r *TopicRepository) Count() (int, error) {
	resultSet := []int{}

	result, err := dat.Count(r)
	if err != nil {
		log.Printf("Error counting Topics", err)
	}

	result.All(&resultSet)
	result.Close()

	return resultSet[0], err
}

func (r *TopicRepository) List(offset, limit int) (TopicList, error) {
	rows, err := dat.List(r, "name", offset, limit)
	if err != nil {
		log.Printf("Error retrieving list of Topics %d to %d: %s", offset, limit, err)
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
