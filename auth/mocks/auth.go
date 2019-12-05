// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	"context"
	"sync"

	"github.com/mainflux/mainflux/auth"
)

var _ auth.KeyRepository = (*keyRepositoryMock)(nil)

type keyRepositoryMock struct {
	mu   sync.Mutex
	keys map[string]auth.Key
}

// NewKeyRepository creates in-memory user repository
func NewKeyRepository() auth.KeyRepository {
	return &keyRepositoryMock{
		keys: make(map[string]auth.Key),
	}
}

func (krm *keyRepositoryMock) Save(ctx context.Context, key auth.Key) (string, error) {
	krm.mu.Lock()
	defer krm.mu.Unlock()

	if _, ok := krm.keys[key.ID]; ok {
		return "", auth.ErrConflict
	}

	krm.keys[key.ID] = key
	return key.ID, nil
}
func (krm *keyRepositoryMock) Retrieve(ctx context.Context, issuer, id string) (auth.Key, error) {
	krm.mu.Lock()
	defer krm.mu.Unlock()

	if key, ok := krm.keys[id]; ok && key.Issuer == issuer {
		return key, nil
	}

	return auth.Key{}, auth.ErrNotFound
}
func (krm *keyRepositoryMock) Remove(ctx context.Context, issuer, id string) error {
	krm.mu.Lock()
	defer krm.mu.Unlock()
	if key, ok := krm.keys[id]; ok && key.Issuer == issuer {
		delete(krm.keys, id)
	}
	return nil
}