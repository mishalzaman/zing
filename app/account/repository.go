package account

import (
	"log"

	db "github.com/dancannon/gorethink"

	"github.com/singnurkar/zing/dat"
)

type AccountRepository struct {
	session *db.Session
	table   string
}

func NewRepository(session *db.Session) *AccountRepository {
	return &AccountRepository{session: session, table: "account"}
}

func (r *AccountRepository) Session() *db.Session { return r.session }
func (r *AccountRepository) Table() string        { return r.table }

func (r *AccountRepository) All(offset, limit int) (Accounts, error) {
	rows, err := dat.All(r, offset, limit)
	if err != nil {
		log.Printf("Error retrieving Accounts %d to %d: %s", offset, limit, err)
	}

	accountsCol := Accounts{}
	if err = rows.All(&accountsCol); err != nil {
		log.Printf("Error decoding rows into slice of Accounts: %s", err)
	}
	rows.Close()

	return accountsCol, err
}

func (r *AccountRepository) One(id string) (*Account, error) {
	cursor, err := dat.One(r, id)
	if err != nil {
		log.Printf("Error retrieving Account %s: %s", id, err)
	}
	defer cursor.Close()

	account := &Account{}
	cursor.One(account)

	return account, err
}

func (r *AccountRepository) Find(username string) (*Account, error) {
	cursor, err := dat.Match(r, "username", username)
	if err != nil {
		log.Printf("Error retrieving account %s: %s", username, err)
	}
	defer cursor.Close()

	account := &Account{}
	cursor.One(account)

	return account, err
}

func (r *AccountRepository) Save(account *Account) error {
	result, err := dat.Create(r, account)
	if err != nil {
		log.Printf("Error creating new Account: %s", err)
	}
	account.SetID(result.GeneratedKeys[0])

	return err
}

func (r *AccountRepository) Update(account *Account) error {
	_, err := dat.Update(r, account.ID(), account)
	if err != nil {
		log.Printf("Error updating Account: %s", err)
	}

	return err
}

func (r *AccountRepository) Purge(id string) error {
	result, err := dat.Purge(r, id)
	if err != nil {
		log.Printf("Error removing Account: %s", err)
	}
	defer result.Close()

	return err
}
