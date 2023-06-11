package session

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

func Get(token string) *Session {
	sessions := GetActiveSessions()
	return (*sessions)[token]
}

func (s *Session) Add() {
	sessions := GetActiveSessions()
	(*sessions)[s.Token] = s
}

func (s *Session) Remove() {
	sessions := GetActiveSessions()
	delete(*sessions, s.Token)
}