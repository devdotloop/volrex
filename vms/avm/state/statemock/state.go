// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/vms/avm/state (interfaces: State)
//
// Generated by this command:
//
//	mockgen -package=statemock -destination=statemock/state.go -mock_names=State=State . State
//

// Package statemock is a generated GoMock package.
package statemock

import (
	reflect "reflect"
	time "time"

	database "github.com/ava-labs/avalanchego/database"
	ids "github.com/ava-labs/avalanchego/ids"
	block "github.com/ava-labs/avalanchego/vms/avm/block"
	txs "github.com/ava-labs/avalanchego/vms/avm/txs"
	avax "github.com/ava-labs/avalanchego/vms/components/avax"
	gomock "go.uber.org/mock/gomock"
)

// State is a mock of State interface.
type State struct {
	ctrl     *gomock.Controller
	recorder *StateMockRecorder
	isgomock struct{}
}

// StateMockRecorder is the mock recorder for State.
type StateMockRecorder struct {
	mock *State
}

// NewState creates a new mock instance.
func NewState(ctrl *gomock.Controller) *State {
	mock := &State{ctrl: ctrl}
	mock.recorder = &StateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *State) EXPECT() *StateMockRecorder {
	return m.recorder
}

// Abort mocks base method.
func (m *State) Abort() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Abort")
}

// Abort indicates an expected call of Abort.
func (mr *StateMockRecorder) Abort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abort", reflect.TypeOf((*State)(nil).Abort))
}

// AddBlock mocks base method.
func (m *State) AddBlock(block block.Block) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBlock", block)
}

// AddBlock indicates an expected call of AddBlock.
func (mr *StateMockRecorder) AddBlock(block any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBlock", reflect.TypeOf((*State)(nil).AddBlock), block)
}

// AddTx mocks base method.
func (m *State) AddTx(tx *txs.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddTx", tx)
}

// AddTx indicates an expected call of AddTx.
func (mr *StateMockRecorder) AddTx(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTx", reflect.TypeOf((*State)(nil).AddTx), tx)
}

// AddUTXO mocks base method.
func (m *State) AddUTXO(utxo *avax.UTXO) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddUTXO", utxo)
}

// AddUTXO indicates an expected call of AddUTXO.
func (mr *StateMockRecorder) AddUTXO(utxo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUTXO", reflect.TypeOf((*State)(nil).AddUTXO), utxo)
}

// Checksum mocks base method.
func (m *State) Checksum() ids.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checksum")
	ret0, _ := ret[0].(ids.ID)
	return ret0
}

// Checksum indicates an expected call of Checksum.
func (mr *StateMockRecorder) Checksum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checksum", reflect.TypeOf((*State)(nil).Checksum))
}

// Close mocks base method.
func (m *State) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *StateMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*State)(nil).Close))
}

// Commit mocks base method.
func (m *State) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *StateMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*State)(nil).Commit))
}

// CommitBatch mocks base method.
func (m *State) CommitBatch() (database.Batch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitBatch")
	ret0, _ := ret[0].(database.Batch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommitBatch indicates an expected call of CommitBatch.
func (mr *StateMockRecorder) CommitBatch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitBatch", reflect.TypeOf((*State)(nil).CommitBatch))
}

// DeleteUTXO mocks base method.
func (m *State) DeleteUTXO(utxoID ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUTXO", utxoID)
}

// DeleteUTXO indicates an expected call of DeleteUTXO.
func (mr *StateMockRecorder) DeleteUTXO(utxoID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUTXO", reflect.TypeOf((*State)(nil).DeleteUTXO), utxoID)
}

// GetBlock mocks base method.
func (m *State) GetBlock(blkID ids.ID) (block.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", blkID)
	ret0, _ := ret[0].(block.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *StateMockRecorder) GetBlock(blkID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*State)(nil).GetBlock), blkID)
}

// GetBlockIDAtHeight mocks base method.
func (m *State) GetBlockIDAtHeight(height uint64) (ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockIDAtHeight", height)
	ret0, _ := ret[0].(ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockIDAtHeight indicates an expected call of GetBlockIDAtHeight.
func (mr *StateMockRecorder) GetBlockIDAtHeight(height any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockIDAtHeight", reflect.TypeOf((*State)(nil).GetBlockIDAtHeight), height)
}

// GetLastAccepted mocks base method.
func (m *State) GetLastAccepted() ids.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastAccepted")
	ret0, _ := ret[0].(ids.ID)
	return ret0
}

// GetLastAccepted indicates an expected call of GetLastAccepted.
func (mr *StateMockRecorder) GetLastAccepted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastAccepted", reflect.TypeOf((*State)(nil).GetLastAccepted))
}

// GetTimestamp mocks base method.
func (m *State) GetTimestamp() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimestamp")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimestamp indicates an expected call of GetTimestamp.
func (mr *StateMockRecorder) GetTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimestamp", reflect.TypeOf((*State)(nil).GetTimestamp))
}

// GetTx mocks base method.
func (m *State) GetTx(txID ids.ID) (*txs.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTx", txID)
	ret0, _ := ret[0].(*txs.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTx indicates an expected call of GetTx.
func (mr *StateMockRecorder) GetTx(txID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTx", reflect.TypeOf((*State)(nil).GetTx), txID)
}

// GetUTXO mocks base method.
func (m *State) GetUTXO(utxoID ids.ID) (*avax.UTXO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUTXO", utxoID)
	ret0, _ := ret[0].(*avax.UTXO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUTXO indicates an expected call of GetUTXO.
func (mr *StateMockRecorder) GetUTXO(utxoID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUTXO", reflect.TypeOf((*State)(nil).GetUTXO), utxoID)
}

// InitializeChainState mocks base method.
func (m *State) InitializeChainState(stopVertexID ids.ID, genesisTimestamp time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeChainState", stopVertexID, genesisTimestamp)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeChainState indicates an expected call of InitializeChainState.
func (mr *StateMockRecorder) InitializeChainState(stopVertexID, genesisTimestamp any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeChainState", reflect.TypeOf((*State)(nil).InitializeChainState), stopVertexID, genesisTimestamp)
}

// IsInitialized mocks base method.
func (m *State) IsInitialized() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInitialized")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsInitialized indicates an expected call of IsInitialized.
func (mr *StateMockRecorder) IsInitialized() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInitialized", reflect.TypeOf((*State)(nil).IsInitialized))
}

// SetInitialized mocks base method.
func (m *State) SetInitialized() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInitialized")
	ret0, _ := ret[0].(error)
	return ret0
}

// SetInitialized indicates an expected call of SetInitialized.
func (mr *StateMockRecorder) SetInitialized() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInitialized", reflect.TypeOf((*State)(nil).SetInitialized))
}

// SetLastAccepted mocks base method.
func (m *State) SetLastAccepted(blkID ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLastAccepted", blkID)
}

// SetLastAccepted indicates an expected call of SetLastAccepted.
func (mr *StateMockRecorder) SetLastAccepted(blkID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastAccepted", reflect.TypeOf((*State)(nil).SetLastAccepted), blkID)
}

// SetTimestamp mocks base method.
func (m *State) SetTimestamp(t time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTimestamp", t)
}

// SetTimestamp indicates an expected call of SetTimestamp.
func (mr *StateMockRecorder) SetTimestamp(t any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimestamp", reflect.TypeOf((*State)(nil).SetTimestamp), t)
}

// UTXOIDs mocks base method.
func (m *State) UTXOIDs(addr []byte, previous ids.ID, limit int) ([]ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UTXOIDs", addr, previous, limit)
	ret0, _ := ret[0].([]ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UTXOIDs indicates an expected call of UTXOIDs.
func (mr *StateMockRecorder) UTXOIDs(addr, previous, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UTXOIDs", reflect.TypeOf((*State)(nil).UTXOIDs), addr, previous, limit)
}
