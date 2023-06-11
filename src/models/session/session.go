package session

import (
	"crypto/sha1"
	"encoding/base64"
)

type Session struct {
	Token string
	Email string
}

type Sessions map[string]*Session

var ActiveSessions *Sessions

type SessionInterface interface {
	Get(string) *Session
	Add()
	Remove()
}

func GetActiveSessions() *Sessions {
	if ActiveSessions == nil {
		ActiveSessions = &Sessions{}
	}

	return ActiveSessions
}

func Encode(email string) string {
	h := sha1.New()
	h.Write([]byte(email))

	// Convierto el hash a base64
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func Get(token string) *Session {
	sessions := GetActiveSessions()
	return (*sessions)[token]
}

func (s *Session) Add() string {
	token := Encode(s.Email)

	sessions := GetActiveSessions()
	(*sessions)[token] = s

	return token
}

func (s *Session) Remove() {
	sessions := GetActiveSessions()
	delete(*sessions, s.Token)
}