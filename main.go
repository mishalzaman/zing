package main

import (
	"io/ioutil"
	"log"
	"net/http"

	db "github.com/dancannon/gorethink"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/justinas/nosurf"
	"gopkg.in/yaml.v2"

	"github.com/singnurkar/zing/app/account"
	"github.com/singnurkar/zing/app/post"
	"github.com/singnurkar/zing/app/topic"
	"github.com/singnurkar/zing/auth"
)

type DatabaseConfig struct {
	Host    string `yaml:"host"`
	Name    string `yaml:"name"`
	MaxOpen int    `yaml:"max_open"`
	MaxIdle int    `yaml:"max_idle"`
}

type Config struct {
	Host     string          `yaml:"host"`
	Prefix   string          `yaml:"prefix"`
	Database *DatabaseConfig `yaml:"database"`
}

var config *Config

func connect(config *DatabaseConfig) *db.Session {
	opts := db.ConnectOpts{
		Address:  config.Host,
		Database: config.Name,
		MaxOpen:  config.MaxOpen,
	}

	session, err := db.Connect(opts)
	if err != nil {
		log.Fatalf("Failed connectiong to Database: %s", err)
	}

	return session
}

func handle(fn func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(fn)
}

func init() {
	rc, _ := ioutil.ReadFile("config.yml")
	err := yaml.Unmarshal(rc, &config)
	if err != nil {
		log.Fatalf("Error unmarshaling config.yml: %s", err)
	}
}

func main() {
	dbconn := connect(config.Database)

	basicauth := auth.NewBasic(dbconn)

	chain := alice.New(basicauth.Basic, nosurf.NewPure)

	router := mux.NewRouter()
	router.StrictSlash(true)

	api := router.PathPrefix("/v1").Subrouter()

	accounts := account.NewController(dbconn)
	posts := post.NewController(dbconn)
	topics := topic.NewController(dbconn)

	api.Handle("/accounts", chain.Then(handle(accounts.All))).Methods("GET")
	api.Handle("/account", chain.Then(handle(accounts.Save))).Methods("POST")
	api.Handle("/account/{id}", chain.Then(handle(accounts.One))).Methods("GET")
	api.Handle("/account/{id}", chain.Then(handle(accounts.Update))).Methods("PUT")
	api.Handle("/account/{id}", chain.Then(handle(accounts.Purge))).Methods("DELETE")

	api.Handle("/posts", chain.Then(handle(posts.All))).Methods("GET")
	api.Handle("/posts/list", chain.Then(handle(posts.List))).Methods("GET")
	api.Handle("/post", chain.Then(handle(posts.Save))).Methods("POST")
	api.Handle("/post/{id}", chain.Then(handle(posts.One))).Methods("GET")
	api.Handle("/post/{id}", chain.Then(handle(posts.Update))).Methods("PUT")
	api.Handle("/post/{id}", chain.Then(handle(posts.Purge))).Methods("DELETE")

	api.Handle("/topics", chain.Then(handle(topics.All))).Methods("GET")
	api.Handle("/topics/list", chain.Then(handle(topics.List))).Methods("GET")
	api.Handle("/topic", chain.Then(handle(topics.Save))).Methods("POST")
	api.Handle("/topic/{id}", chain.Then(handle(topics.One))).Methods("GET")
	api.Handle("/topic/{id}", chain.Then(handle(topics.Update))).Methods("PUT")
	api.Handle("/topic/{id}", chain.Then(handle(topics.Purge))).Methods("DELETE")

	api.Handle("/topic/{id}/posts", chain.Then(handle(topics.Posts))).Methods("GET")

	api.Handle("/topic/{id}/parents", chain.Then(handle(topics.Parents))).Methods("GET")
	api.Handle("/topic/{id}/parents", chain.Then(handle(topics.AddParents))).Methods("PUT")
	api.Handle("/topic/{id}/parents", chain.Then(handle(topics.RemoveParents))).Methods("DELETE")

	log.Fatal(http.ListenAndServe(config.Host, router))
}
