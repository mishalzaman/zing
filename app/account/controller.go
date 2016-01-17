package account

import (
	"net/http"
	"time"

	db "github.com/dancannon/gorethink"
	"github.com/gorilla/mux"

	"github.com/singnurkar/zing/util"
)

type AccountController struct {
	Accounts *AccountRepository
}

func NewController(session *db.Session) *AccountController {
	repo := NewRepository(session)
	return &AccountController{repo}
}

func (c *AccountController) All(res http.ResponseWriter, req *http.Request) {
	if accounts, err := c.Accounts.All(0, 15); err == nil {
		util.WriteResponse(res, util.Payload{Result: accounts}, http.StatusOK)
	}
}

func (c *AccountController) One(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if account, err := c.Accounts.One(vars["id"]); err == nil {
		util.WriteResponse(res, util.Payload{Result: account}, http.StatusOK)
	}
}

func (c *AccountController) Save(res http.ResponseWriter, req *http.Request) {
	a := NewAccount()
	util.DecodeReqBody(req.Body, a)
	if err := c.Accounts.Save(a); err == nil {
		util.WriteResponse(res, util.Payload{Result: a.ID()}, http.StatusCreated)
	}
}

func (c *AccountController) Update(res http.ResponseWriter, req *http.Request) {
	a := &Account{}
	util.DecodeReqBody(req.Body, a)
	a.Modified = time.Now()
	if err := c.Accounts.Save(a); err == nil {
		util.WriteResponse(res, util.Payload{Result: a.ID()}, http.StatusAccepted)
	}
}

func (c *AccountController) Purge(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if err := c.Accounts.Purge(vars["id"]); err == nil {
		util.WriteResponse(res, util.Payload{Success: "Deleted " + vars["id"]}, http.StatusOK)
	}
}
