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
	if list, err := c.Topics.List(0, 1000); err != nil {
		msg := "Could not retrieve list of Topics"
		util.LogError(msg, err)
		util.SendError(res, msg, http.StatusInternalServerError)
		return
	}

	util.Send(res, util.Payload{Result: list}, http.StatusOK)
}

func (c *TopicController) All(res http.ResponseWriter, req *http.Request) {
	if topics, err := c.Topics.All(0, 15); err != nil {
		msg := "Could not retrieve topics"
		util.LogError(msg, err)
		util.SendError(res, msg, http.StatusInternalServerError)
		return
	}

	util.Send(res, util.Payload{Result: topics}, http.StatusOK)
}

func (c *TopicController) One(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if topic, err := c.Topics.One(vars["id"]); err != nil {
		msg := "Could not retrieve topic"
		util.LogError(msg, err)
		util.SendError(res, msg, http.StatusInternalServerError)
		return
	}

	util.Send(res, util.Payload{Result: topic}, http.StatusFound)
}

func (c *TopicController) Save(res http.ResponseWriter, req *http.Request) {
	t := NewTopic()
	util.DecodeReqBody(req.Body, t)

	v := validate.NewValidator()
	v.NotEmptyString(t.Name)
	v.NoSpaces(t.Name)
	if v.NotValid() {
		msg := "Invalid topic name"
		util.LogError(msg, v)
		util.SendError(res, msg, http.StatusBadRequest)
		return
	}

	if err := c.Topics.Save(t); err != nil {
		msg := "Could not save Topic"
		util.LogError(msg, err)
		util.SendError(res, msg, http.StatusInternalServerError)
		return
	}

	util.Send(res, util.Payload{Result: t.ID()}, http.StatusCreated)
}

func (c *TopicController) Update(res http.ResponseWriter, req *http.Request) {
	t := &Topic{}
	util.DecodeReqBody(req.Body, t)
	t.Modified = time.Now()

	if err := c.Topics.Update(t); err != nil {
		msg := "Could not update Topic"
		util.LogError(msg, err)
		util.SendError(res, msg, http.StatusInternalServerError)
		return
	}

	util.Send(res, util.Payload{Result: t.ID()}, http.StatusOK)
}

func (c *TopicController) Purge(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if err := c.Topics.Purge(vars["id"]); err != nil {
		msg := "Unable to delete Topic"
		util.LogError(msg, err)
		util.SendError(res, msg, http.StatusInternalServerError)
		return
	}

	util.Send(res, util.Payload{Success: "Deleted Topic"}, http.StatusOK)
}
