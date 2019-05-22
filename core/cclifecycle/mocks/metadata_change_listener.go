// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import chaincode "github.com/hyperledger/fabric/common/chaincode"
import mock "github.com/stretchr/testify/mock"

// MetadataChangeListener is an autogenerated mock type for the MetadataChangeListener type
type MetadataChangeListener struct {
	mock.Mock
}

// HandleMetadataUpdate provides a mock function with given fields: channel, chaincodes
func (_m *MetadataChangeListener) HandleMetadataUpdate(channel string, chaincodes chaincode.MetadataSet) {
	_m.Called(channel, chaincodes)
}
