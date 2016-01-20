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
	"github.com/singnurkar/zing/window"
)

type DatabaseConfig struct {
	Host    string `yaml:"host"`
	Name    string `yaml:"name"`
	MaxOpen int    `yaml:"max_open"`
	MaxIdle int    `yaml:"max_idle"`
}

type Config struct {
	Dev      bool            `yaml:"dev"`
	Title    string          `yaml:"title"`
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
	// basicauth := auth.NewBasic(dbconn)
	chain := alice.New(nosurf.NewPure)
	router := mux.NewRouter()
	router.StrictSlash(true)

	routes := map[string]string{
		"root":          "/",
		"assets":        "/assets",
		"accounts":      "/accounts",
		"new_account":   "/account",
		"account":       "/account/{id}",
		"posts":         "/posts",
		"count_posts":   "/posts/count",
		"list_posts":    "/posts/list",
		"new_post":      "/post",
		"post":          "/post/{id}",
		"topics":        "/topics",
		"count_topics":  "/topics/count",
		"list_topics":   "/topics/list",
		"new_topic":     "/topic",
		"topic":         "/topic/{id}",
		"topic_parents": "/topic/{id}/parents",
		"topic_posts":   "/topic/{id}/posts",
	}

	w := window.NewController(&window.Options{Title: config.Title, Dev: config.Dev})
	router.Handle(routes["root"], chain.Then(handle(w.Render))).Methods("GET")

	fs := http.FileServer(http.Dir("window/dist/"))
	router.PathPrefix(routes["assets"]).Handler(http.StripPrefix("/assets/", fs))

	api := router.PathPrefix("/v1").Subrouter()
	accounts := account.NewController(dbconn)
	posts := post.NewController(dbconn)
	topics := topic.NewController(dbconn)

	api.Handle(routes["accounts"], chain.Then(handle(accounts.All))).Methods("GET")
	api.Handle(routes["new_accounts"], chain.Then(handle(accounts.Save))).Methods("POST")
	api.Handle(routes["account"], chain.Then(handle(accounts.One))).Methods("GET")
	api.Handle(routes["account"], chain.Then(handle(accounts.Update))).Methods("PUT")
	api.Handle(routes["account"], chain.Then(handle(accounts.Purge))).Methods("DELETE")

	api.Handle(routes["posts"], chain.Then(handle(posts.All))).Methods("GET")
	api.Handle(routes["list_posts"], chain.Then(handle(posts.List))).Methods("GET")
	api.Handle(routes["new_post"], chain.Then(handle(posts.Save))).Methods("POST")
	api.Handle(routes["post"], chain.Then(handle(posts.One))).Methods("GET")
	api.Handle(routes["post"], chain.Then(handle(posts.Update))).Methods("PUT")
	api.Handle(routes["post"], chain.Then(handle(posts.Purge))).Methods("DELETE")

	api.Handle(routes["topics"], chain.Then(handle(topics.All))).Methods("GET")
	api.Handle(routes["list_topic"], chain.Then(handle(topics.List))).Methods("GET")
	api.Handle(routes["new_topic"], chain.Then(handle(topics.Save))).Methods("POST")
	api.Handle(routes["topic"], chain.Then(handle(topics.One))).Methods("GET")
	api.Handle(routes["topic"], chain.Then(handle(topics.Update))).Methods("PUT")
	api.Handle(routes["topic"], chain.Then(handle(topics.Purge))).Methods("DELETE")
	api.Handle(routes["topic_parents"], chain.Then(handle(topics.Parents))).Methods("GET")
	api.Handle(routes["topic_parents"], chain.Then(handle(topics.AddParents))).Methods("PUT")
	api.Handle(routes["topic_parents"], chain.Then(handle(topics.RemoveParents))).Methods("DELETE")
	api.Handle(routes["topic_posts"], chain.Then(handle(topics.Posts))).Methods("GET")

	log.Fatal(http.ListenAndServe(config.Host, router))
}
