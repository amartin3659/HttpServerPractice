package session

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNew(t *testing.T) {
  session := New()
  if reflect.TypeOf(session).String() != "*session.Sessions" {
    t.Error("Returned unexpected type", reflect.TypeOf(session).String())
  }
}

func TestGet(t *testing.T) {
  // ok
  // create new sessions slice
  sessions := New()
  // add session
  sid := uuid.New()
  uid := uuid.New()
  s := Session{
    SessionID: sid,
    UserID: uid,
  }
  sessions.Add(s)
  userID, err := sessions.Get(sid)
  if err != nil {
    t.Error("Could not find session")
  }
  if userID != uid {
    t.Error("User ids did not match")
  }
  // no session
  // create new sessions slice
  sessions = New()
  // add session
  sid = uuid.New()
  uid = uuid.New()
  s = Session{
    SessionID: sid,
    UserID: uid,
  }
  sessions.Add(s)
  userID, err = sessions.Get(uuid.New())
  if err == nil {
    t.Error("Expected to not find session")
  }
}

func TestAdd(t *testing.T) {
  // create sessions slice
  sessions := New()
  sid := uuid.New()
  uid := uuid.New()
  s := Session{
    SessionID: sid,
    UserID: uid,
  }
  sessions.Add(s)
  // add new session with same userid
  newSID := uuid.New()
  ns := Session{
    SessionID: newSID,
    UserID: uid,
  }
  sessions.Add(ns)
  userID, err := sessions.Get(newSID)
  if err != nil {
    t.Error("Expected there to be a valid session")
  }
  if userID != uid {
    t.Error("Expected new session id to match user id")
  }
}

func TestRemove(t *testing.T) {
  // create sessions slice
  sessions := New()
  sid := uuid.New()
  uid := uuid.New()
  s := Session{
    SessionID: sid,
    UserID: uid,
  }
  sessions.Add(s)
  sessions.Remove(sid)
  // check if session was removed
  _, err := sessions.Get(sid)
  if err == nil {
    t.Error("Expected to not get a session")
  }
}
