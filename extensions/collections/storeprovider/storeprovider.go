/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package storeprovider

import (
	proto "github.com/hyperledger/fabric-protos-go/transientstore"
	storeapi "github.com/hyperledger/fabric/extensions/collections/api/store"
)

// NewProviderFactory returns a new private data store provider factory
func NewProviderFactory() *StoreProvider {
	return &StoreProvider{}
}

// StoreProvider is a noop implementation of a private data store provider
type StoreProvider struct {
}

// Initialize initializes the store provider
func (sp *StoreProvider) Initialize() {
	// Noop by default
}

// StoreForChannel returns a noop store
func (sp *StoreProvider) StoreForChannel(channel string) storeapi.Store {
	return &store{}
}

// OpenStore returns a noop store
func (sp *StoreProvider) OpenStore(ledgerID string) (storeapi.Store, error) {
	return &store{}, nil
}

// Close does nothing
func (sp *StoreProvider) Close() {
}

type store struct {
}

func (s *store) Persist(txid string, privateSimulationResultsWithConfig *proto.TxPvtReadWriteSetWithConfigInfo) error {
	return nil
}

func (s *store) Close() {
}
