package auth

import (
	"crypto/rand"
	mathrand "math/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
	"time"
)

type SessionStore interface {
	Generate(phone string) (string, error)
	Delete(sessionId string) error
	Save(sessionId string, phone string) error
	Get(sessionId string) (*Session, error)
}

type Session struct {
	Phone string
	Code string
}

func GenerateCode() string {
	return fmt.Sprintf("%06d", mathrand.Intn(1000000))
}

func (s *Session) GenerateAndSetCode() {
	s.Code = GenerateCode()
}


type InMemorySessionStore struct {
	mu       sync.RWMutex
	sessions map[string]Session
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
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.sessions[sessionId]; exists {
		return errors.New(ErrSessionIdAlreadyExists)
	}
	s.sessions[sessionId] = Session{Phone: phone, Code: GenerateCode()}
	return nil
}

func (s *InMemorySessionStore) Get(sessionId string) (*Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, exists := s.sessions[sessionId]
	if !exists {
		return nil, errors.New(ErrSessionIdDoesNotExist)
	}
	return &session, nil
}

func (s *InMemorySessionStore) Delete(sessionId string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.sessions, sessionId)
	return nil
}

func NewSessionStore() SessionStore {
	return &InMemorySessionStore{
		sessions: make(map[string]Session),
	}
}