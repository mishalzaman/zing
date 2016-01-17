package post

import (
	"net/http"
	"time"

	db "github.com/dancannon/gorethink"
	"github.com/gorilla/mux"

	"github.com/singnurkar/zing/util"
)

type PostController struct {
	Posts *PostRepository
}

func NewController(session *db.Session) *PostController {
	repo := NewRepository(session)
	return &PostController{Posts: repo}
}

func (c *PostController) List(res http.ResponseWriter, req *http.Request) {
	if list, err := c.Posts.List(0, 1000); err == nil {
		util.WriteResponse(res, util.Payload{Result: list}, http.StatusOK)
	}
}

func (c *PostController) All(res http.ResponseWriter, req *http.Request) {
	var payload util.Payload
	var resCode int

	if posts, err := c.Posts.All(0, 15); err == nil {
		payload = util.Payload{Result: posts}
		resCode = http.StatusOK
	} else {
		payload = util.Payload{Error: "Could not retrieve posts"}
		resCode = http.StatusInternalServerError
	}

	util.WriteResponse(res, payload, resCode)
}

func (c *PostController) One(res http.ResponseWriter, req *http.Request) {
	var payload util.Payload
	var resCode int

	vars := mux.Vars(req)
	if post, err := c.Posts.One(vars["id"]); err == nil {
		payload = util.Payload{Result: post}
		resCode = http.StatusFound
	} else {
		payload = util.Payload{Error: "Could not retrieve post"}
		resCode = http.StatusInternalServerError
	}

	util.WriteResponse(res, payload, resCode)
}

func (c *PostController) Save(res http.ResponseWriter, req *http.Request) {
	var payload util.Payload
	var resCode int

	p := NewPost()
	util.DecodeReqBody(req.Body, p)
	if err := c.Posts.Save(p); err == nil {
		payload = util.Payload{Result: p.ID()}
		resCode = http.StatusCreated
	} else {
		payload = util.Payload{Result: p.ID()}
		resCode = http.StatusCreated
	}

	util.WriteResponse(res, payload, resCode)

}

func (c *PostController) Update(res http.ResponseWriter, req *http.Request) {
	p := &Post{}
	util.DecodeReqBody(req.Body, p)
	p.Modified = time.Now()
	if err := c.Posts.Save(p); err == nil {
		util.WriteResponse(res, util.Payload{Result: p.ID()}, http.StatusAccepted)
	}
}

func (c *PostController) Purge(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if err := c.Posts.Purge(vars["id"]); err == nil {
		util.WriteResponse(res, util.Payload{Success: "Deleted " + vars["id"]}, http.StatusOK)
	}
}
