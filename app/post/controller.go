package post

import (
	"net/http"
	"time"

	db "github.com/dancannon/gorethink"
	"github.com/gorilla/mux"

	"github.com/singnurkar/zing/util"
	"github.com/singnurkar/zing/validate"
)

type PostController struct {
	Posts *PostRepository
}

func NewController(session *db.Session) *PostController {
	repo := NewRepository(session)
	return &PostController{Posts: repo}
}

func (c *PostController) List(res http.ResponseWriter, req *http.Request) {
	list, err := c.Posts.List(0, 1000)
	if err != nil {
		msg := "Could not retrieve list of posts"
		util.LogError(msg, err)
		util.SendError(res, msg, http.StatusInternalServerError)
		return
	}

	util.Send(res, util.Payload{Result: list}, http.StatusOK)
}

func (c *PostController) All(res http.ResponseWriter, req *http.Request) {
	posts, err := c.Posts.All(0, 15)
	if err != nil {
		msg := "Could not retrieve posts"
		util.LogError(msg, err)
		util.SendError(res, msg, http.StatusInternalServerError)
	}

	util.Send(res, util.Payload{Result: posts}, http.StatusOK)
}

func (c *PostController) One(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	post, err := c.Posts.One(vars["id"])
	if err != nil {
		msg := "Could not retrieve post"
		util.LogError(msg, err)
		util.SendError(res, msg, http.StatusInternalServerError)
		return
	}

	util.Send(res, util.Payload{Result: post}, http.StatusFound)
}

func (c *PostController) Save(res http.ResponseWriter, req *http.Request) {
	p := NewPost()
	util.DecodeReqBody(req.Body, p)

	v := validate.NewValidator()
	v.NotEmptyString(p.Title)
	v.NoSpaces(p.Title)
	if v.NotValid() {
		msg := "Invalid post name"
		util.LogError(msg, v)
		util.SendError(res, msg, http.StatusBadRequest)
		return
	}

	if err := c.Posts.Save(p); err == nil {
		msg := "Could not save post"
		util.LogError(msg, err)
		util.SendError(res, msg, http.StatusInternalServerError)
		return
	}

	util.Send(res, util.Payload{Result: p.ID()}, http.StatusCreated)
}

func (c *PostController) Update(res http.ResponseWriter, req *http.Request) {
	p := &Post{}
	util.DecodeReqBody(req.Body, p)
	p.Modified = time.Now()

	if err := c.Posts.Update(p); err != nil {
		msg := "Could not update post"
		util.LogError(msg, err)
		util.SendError(res, msg, http.StatusInternalServerError)
		return
	}

	util.Send(res, util.Payload{Result: p.ID()}, http.StatusOK)
}

func (c *PostController) Purge(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if err := c.Posts.Purge(vars["id"]); err != nil {
		msg := "Could not purge post"
		util.LogError(msg, err)
		util.SendError(res, msg, http.StatusInternalServerError)
		return
	}

	util.Send(res, util.Payload{Success: "Deleted post"}, http.StatusOK)

}
