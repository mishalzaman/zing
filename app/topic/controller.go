package topic

import (
	"net/http"
	"time"

	db "github.com/dancannon/gorethink"
	"github.com/gorilla/mux"

	"github.com/singnurkar/zing/util"
	"github.com/singnurkar/zing/validate"
)

type TopicController struct {
	Topics *TopicRepository
}

func NewController(session *db.Session) *TopicController {
	repo := NewRepository(session)
	return &TopicController{Topics: repo}
}

func (c *TopicController) List(res http.ResponseWriter, req *http.Request) {
	if list, err := c.Topics.List(0, 1000); err == nil {
		util.WriteResponse(res, util.Payload{Result: list}, http.StatusOK)
	}
}

func (c *TopicController) All(res http.ResponseWriter, req *http.Request) {
	var payload util.Payload
	var resCode int

	if topics, err := c.Topics.All(0, 15); err == nil {
		payload = util.Payload{Result: topics}
		resCode = http.StatusOK
	} else {
		payload = util.Payload{Error: "Could not retrieve topics"}
		resCode = http.StatusInternalServerError
	}

	util.WriteResponse(res, payload, resCode)
}

func (c *TopicController) One(res http.ResponseWriter, req *http.Request) {
	var payload util.Payload
	var resCode int

	vars := mux.Vars(req)
	if topic, err := c.Topics.One(vars["id"]); err == nil {
		payload = util.Payload{Result: topic}
		resCode = http.StatusFound
	} else {
		payload = util.Payload{Error: "Could not retrieve topic"}
		resCode = http.StatusInternalServerError
	}

	util.WriteResponse(res, payload, resCode)
}

func (c *TopicController) Save(res http.ResponseWriter, req *http.Request) {
	var payload util.Payload
	var resCode int
	var err error

	t := NewTopic()
	util.DecodeReqBody(req.Body, t)

	v := validate.NewValidator()
	err = v.NotEmptyString(t.Name)
	if err != nil {
		payload = util.Payload{Error: err.Error()}
		resCode = http.StatusBadRequest
	}

	if err := c.Topics.Save(t); err == nil {
		payload = util.Payload{Result: t.ID()}
		resCode = http.StatusCreated
	} else {
		payload = util.Payload{Error: "Could not save Topic"}
		resCode = http.StatusCreated
	}

	util.WriteResponse(res, payload, resCode)

}

func (c *TopicController) Update(res http.ResponseWriter, req *http.Request) {
	p := &Topic{}
	util.DecodeReqBody(req.Body, p)
	p.Modified = time.Now()
	if err := c.Topics.Save(p); err == nil {
		util.WriteResponse(res, util.Payload{Result: p.ID()}, http.StatusAccepted)
	}
}

func (c *TopicController) Purge(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if err := c.Topics.Purge(vars["id"]); err == nil {
		util.WriteResponse(res, util.Payload{Success: "Deleted " + vars["id"]}, http.StatusOK)
	}
}
