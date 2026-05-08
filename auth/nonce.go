package auth


import (	
	"time"

	"sync"
)


type NonceItem struct {
	Nonce string
	Expiration time.Time
}

type NonceStore struct {
	store map[string]NonceItem
	mu sync.RWMutex
}
func NewNonceStore() *NonceStore {
	return &NonceStore{
		store: make(map[string]NonceItem),
	}
}
var nonceStoreInstance = NewNonceStore();

func SetNonce(address string, nonce string, expiration time.Time) {
	nonceStoreInstance.mu.Lock()
	defer nonceStoreInstance.mu.Unlock()
	nonceStoreInstance.store[address] = NonceItem{
		Nonce: nonce,
		Expiration: expiration,
	}
}

func GetNonce(address string) (string, bool) {
	nonceStoreInstance.mu.RLock()
	defer nonceStoreInstance.mu.RUnlock()
	item, exists := nonceStoreInstance.store[address]
	if !exists || time.Now().After(item.Expiration) {
		return "", false
	}
	return item.Nonce, true
}

func DeleteNonce(address string) {
	nonceStoreInstance.mu.Lock()
	defer nonceStoreInstance.mu.Unlock()
	delete(nonceStoreInstance.store, address)
}