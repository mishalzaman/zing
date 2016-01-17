package auth

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	db "github.com/dancannon/gorethink"
	"github.com/gorilla/context"
)

type Basic struct {
	session *db.Session
}

func NewBasic(session *db.Session) *Basic {
	return &Basic{session: session}
}

func (a *Basic) Basic(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if len(req.Header["Authorization"]) == 0 {
			reject(res)
			return
		}

		passport := strings.TrimSpace(req.Header["Authorization"][0])
		if !strings.HasPrefix(passport, "Basic") {
			reject(res)
			return
		}

		credentials := strings.Split(passport, " ")
		fingerprint, err := base64.StdEncoding.DecodeString(credentials[1])
		if err != nil {
			reject(res)
			return
		}

		client, err := a.identify(strings.Split(string(fingerprint), ":"))
		if err != nil {
			reject(res)
			return
		}

		context.Set(req, "client", client)
		h.ServeHTTP(res, req)
	})
}

func (a *Basic) identify(identity []string) (*Client, error) {
	clientRepo := NewClientRepository(a.session)
	client, err := clientRepo.Find(identity[0])
	if err == nil && identity[1] == client.Secret {
		return client, nil
	}
	return &Client{}, errors.New("Cannot identify client")
}

func reject(res http.ResponseWriter) {
	res.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
	http.Error(res, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}
