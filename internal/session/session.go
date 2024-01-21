package session

import (
	"errors"

	"github.com/google/uuid"
)

type Session struct {
	SessionID uuid.UUID
	UserID    uuid.UUID
}

var data Sessions
type Sessions []Session

func New() *Sessions {
  return &data
}

func (s *Sessions) Get(sessionID uuid.UUID) (uuid.UUID, error) {
  for _, session := range *s {
    if session.SessionID == sessionID {
      return session.UserID, nil
    }
  }

  return uuid.Nil, errors.New("No session data")
}

func (s *Sessions) Add(newSession Session) {
  // remove any old session data for userID
  for i, session := range *s {
    if session.UserID == newSession.UserID {
      *s = append((*s)[:i], (*s)[i+1:]...)  
    }
  }
  // add new session data
  *s = append(*s, newSession)
}

func (s *Sessions) Remove(sessionID uuid.UUID) {
  for i, session := range *s {
    if sessionID == session.SessionID {
      *s = append((*s)[:i], (*s)[i+1:]...)
    }
  }
}
