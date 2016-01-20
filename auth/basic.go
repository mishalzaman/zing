package auth

import (
	"encoding/base64"
	"encoding/json"
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
			reject(res, "No credentials given")
			return
		}

		passport := strings.TrimSpace(req.Header["Authorization"][0])
		if !strings.HasPrefix(passport, "Basic") {
			reject(res, "Unrecognized auth scheme")
			return
		}

		credentials := strings.Split(passport, " ")
		fingerprint, err := base64.StdEncoding.DecodeString(credentials[1])
		if err != nil {
			reject(res, err.Error())
			return
		}

		client, err := a.identify(strings.Split(string(fingerprint), ":"))
		if err != nil {
			reject(res, err.Error())
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

func reject(res http.ResponseWriter, msg string) {
	// http.Error(res, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	content, err := json.Marshal(map[string]string{"error": msg})
	if err == nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		res.Header().Set("Content-Type", "application/json")
		res.Write(content)
	}
}
