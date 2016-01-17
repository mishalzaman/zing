package auth

import (
	"log"

	db "github.com/dancannon/gorethink"

	"github.com/singnurkar/zing/dat"
)

type Client struct {
	Id     string
	Name   string
	Email  string
	URL    string
	Phone  string
	Key    string
	Secret string
}

type ClientRepository struct {
	session *db.Session
	table   string
}

func NewClientRepository(session *db.Session) *ClientRepository {
	return &ClientRepository{session: session, table: "client"}
}

func (r *ClientRepository) Session() *db.Session { return r.session }
func (r *ClientRepository) Table() string        { return r.table }

func (r *ClientRepository) Find(key string) (*Client, error) {
	cursor, err := dat.Match(r, "key", key)
	if err != nil {
		log.Printf("Error retrieving client %s: %s", key, err)
	}
	defer cursor.Close()

	client := &Client{}
	cursor.One(client)

	return client, err
}
