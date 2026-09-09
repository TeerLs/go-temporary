package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

type SessionStore interface {
	Generate(phone string) (string, error)
	Delete(sessionId string) error
	Save(sessionId string, phone string) error
	Get(sessionId string) (string, error)
}

type InMemorySessionStore struct {
	sessions map[string]string
}

func (s *InMemorySessionStore) Generate(phone string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	currentTime := time.Now().UnixNano()

	input := fmt.Sprintf("%x-%d-%s", salt, currentTime, phone)

	hash := sha256.Sum256([]byte(input))

	return hex.EncodeToString(hash[:]), nil
}

func (s *InMemorySessionStore) Save(sessionId string, phone string) error {
	if _, exists := s.sessions[sessionId]; exists {
		return errors.New(ErrSessionIdAlreadyExists)
	}
	s.sessions[sessionId] = phone
	return nil
}

func (s *InMemorySessionStore) Get(sessionId string) (string, error) {
	phone, exists := s.sessions[sessionId]
	if !exists {
		return "", errors.New(ErrSessionIdDoesNotExist)
	}
	return phone, nil
}

func (s *InMemorySessionStore) Delete(sessionId string) error {
	delete(s.sessions, sessionId)
	return nil
}

func NewSessionStore() SessionStore {
	return &InMemorySessionStore{
		sessions: make(map[string]string),
	}
}