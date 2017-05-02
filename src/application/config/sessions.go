package config

import (
	"github.com/gorilla/sessions"
	"github.com/gorilla/securecookie"
)

var appKey = []byte("bIPh/yqFb2BDJgqz6QCTtILESy8CBfQwBFD/+YGBU2Tl6H0miIdzSw==")


func Sessions() *sessions.CookieStore {
	cs := &sessions.CookieStore{
		Codecs: securecookie.CodecsFromPairs(appKey),
		Options: &sessions.Options{
			Path:   "/",
			MaxAge: 86400 * 7,
		},
	}

	cs.MaxAge(cs.Options.MaxAge)
	return cs
}