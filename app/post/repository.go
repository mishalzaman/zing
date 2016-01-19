package post

import (
	"log"

	db "github.com/dancannon/gorethink"

	"github.com/singnurkar/zing/core"
	"github.com/singnurkar/zing/dat"
)

type TopicRepository struct {
	session *db.Session
	table   string
}

func (r *TopicRepository) Session() *db.Session { return r.session }
func (r *TopicRepository) Table() string        { return r.table }

type PostRepository struct {
	session *db.Session
	table   string
	topics  TopicRepository
}

func NewRepository(session *db.Session) *PostRepository {
	return &PostRepository{
		session: session,
		table:   "post",
		topics:  TopicRepository{session, "topic_post"},
	}
}

func (r *PostRepository) Session() *db.Session { return r.session }
func (r *PostRepository) Table() string        { return r.table }

func (r *PostRepository) Count() (int, error) {
	resultSet := []int{}

	result, err := dat.Count(r)
	if err != nil {
		log.Printf("Error counting posts", err)
	}

	result.All(&resultSet)
	result.Close()

	return resultSet[0], err
}

func (r *PostRepository) List(offset, limit int) (core.PostList, error) {
	rows, err := dat.List(r, "title", offset, limit)
	if err != nil {
		log.Printf("Error retrieving list of posts %d to %d: %s", offset, limit, err)
	}

	posts := []map[string]string{}
	if err = rows.All(&posts); err != nil {
		log.Printf("Error decoding rows into slice: %s", err)
	}
	rows.Close()

	list := core.PostList{}
	for _, v := range posts {
		list[v["id"]] = v["title"]
	}

	return list, err
}

func (r *PostRepository) All(offset, limit int) (core.Posts, error) {
	rows, err := dat.All(r, offset, limit)
	if err != nil {
		log.Printf("Error retrieving posts %d to %d: %s", offset, limit, err)
	}

	postsCol := core.Posts{}
	if err = rows.All(&postsCol); err != nil {
		log.Printf("Error decoding rows into slice of posts: %s", err)
	}
	rows.Close()

	return postsCol, err
}

func (r *PostRepository) One(id string) (*core.Post, error) {
	cursor, err := dat.One(r, id)
	if err != nil {
		log.Printf("Error retrieving post %s: %s", id, err)
	}
	defer cursor.Close()

	post := &core.Post{}
	cursor.One(post)

	return post, err
}

func (r *PostRepository) Save(post *core.Post) error {
	result, err := dat.Create(r, post)
	if err != nil {
		log.Printf("Error creating new post: %s", err)
	}
	post.SetID(result.GeneratedKeys[0])

	return err
}

func (r *PostRepository) Update(post *core.Post) error {
	_, err := dat.Update(r, post.ID(), post)
	if err != nil {
		log.Printf("Error updating post: %s", err)
	}

	return err
}

func (r *PostRepository) Purge(id string) error {
	result, err := dat.Purge(r, id)
	if err != nil {
		log.Printf("Error removing post: %s", err)
	}
	defer result.Close()

	return err
}
