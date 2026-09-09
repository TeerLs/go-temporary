package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"
)

type SessionStore interface {
	Delete(sessionId string) error
	Save(phone string) (error, string)
	Get(sessionId string) (*Session, error)
}

type Session struct {
	Phone string
	Code  string
}

func GenerateCode() (string, error) {
	number, err := rand.Int(rand.Reader, big.NewInt(10000))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%04d", number.Int64()), nil
}

type InMemorySessionStore struct {
	mu       sync.RWMutex
	sessions map[string]Session
}

func GenerateSessionId(phone string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	currentTime := time.Now().UnixNano()

	input := fmt.Sprintf("%x-%d-%s", salt, currentTime, phone)

	hash := sha256.Sum256([]byte(input))

	return hex.EncodeToString(hash[:]), nil
}

func (s *InMemorySessionStore) Save(phone string) (error, string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	sessionId, err := GenerateSessionId(phone)
	if err != nil {
		return err, ""
	}

	if _, exists := s.sessions[sessionId]; exists {
		return errors.New(ErrSessionIdAlreadyExists), ""
	}
	code, err := GenerateCode()
	if err != nil {
		return err, ""
	}
	s.sessions[sessionId] = Session{Phone: phone, Code: code}
	return nil, sessionId
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
