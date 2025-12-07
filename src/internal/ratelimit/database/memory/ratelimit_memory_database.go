package memory

import (
	"sync"
	"time"
)

type DB struct {
	tokens      map[string]int
	refillTimes map[string]time.Time
	tokensMu    sync.RWMutex
	refillMu    sync.RWMutex
}

func NewDB() *DB {
	return &DB{
		tokens:      make(map[string]int),
		refillTimes: make(map[string]time.Time),
	}
}

func (db *DB) GetTokens(client string) (int, error) {
	db.tokensMu.RLock()
	defer db.tokensMu.RUnlock()

	tokens, exists := db.tokens[client]
	if !exists {
		return 0, nil
	}
	return tokens, nil
}

func (db *DB) SetTokens(client string, tokens int) error {
	db.tokensMu.Lock()
	defer db.tokensMu.Unlock()

	db.tokens[client] = tokens
	return nil
}

func (db *DB) GetLastRefill(client string) (time.Time, error) {
	db.refillMu.RLock()
	defer db.refillMu.RUnlock()

	t, exists := db.refillTimes[client]
	if !exists {
		return time.Now().Add((-5) * time.Minute), nil
	}
	return t, nil
}

func (db *DB) SetLastRefill(client string, t time.Time) error {
	db.refillMu.Lock()
	defer db.refillMu.Unlock()

	db.refillTimes[client] = t
	return nil
}
