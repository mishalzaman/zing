package dat

import db "github.com/dancannon/gorethink"

/**
 * Read/Write to RethinkDB
 */

type Repository interface {
	Session() *db.Session
	Table() string
}

type Query map[string]interface{}

func Count(r Repository) (*db.Cursor, error) {
	return db.Table(r.Table()).Count().Run(r.Session())
}

func Exists(r Repository, field string, value string) (*db.Cursor, error) {
	return db.Table(r.Table()).
		Filter(Query{field: value}).
		Slice(0, 1).
		Pluck(field).
		Run(r.Session())
}

func All(r Repository, offset int, limit int) (*db.Cursor, error) {
	return db.Table(r.Table()).
		OrderBy(db.OrderByOpts{Index: db.Desc("id")}).
		Slice(offset, limit).
		Run(r.Session())
}

func List(r Repository, nameField string, offset int, limit int) (*db.Cursor, error) {
	return db.Table(r.Table()).
		OrderBy(db.OrderByOpts{Index: db.Desc("id")}).
		Slice(offset, limit).
		Pluck("id", nameField).
		Run(r.Session())
}

func Slot(r Repository) (*db.Cursor, error) {
	return db.Table(r.Table()).
		Filter(db.Row.Field("slot").Gt(0)).
		OrderBy(db.Desc("slot")).
		Pluck("id", "name", "slot").
		Run(r.Session())
}

func One(r Repository, id string) (*db.Cursor, error) {
	return db.Table(r.Table()).
		Filter(map[string]string{"id": id}).
		Run(r.Session())
}

func Match(r Repository, field string, value interface{}) (*db.Cursor, error) {
	return db.Table(r.Table()).
		Filter(Query{field: value}).
		Run(r.Session())
}

func Create(r Repository, v interface{}) (db.WriteResponse, error) {
	return db.Table(r.Table()).
		Insert(v).
		RunWrite(r.Session())
}

func Update(r Repository, id string, v interface{}) (db.WriteResponse, error) {
	return db.Table(r.Table()).
		Filter(Query{"id": id}).
		Update(v).
		RunWrite(r.Session())
}

func Purge(r Repository, id string) (*db.Cursor, error) {
	return db.Table(r.Table()).
		Filter(Query{"id": id}).
		Delete().
		Run(r.Session())
}
