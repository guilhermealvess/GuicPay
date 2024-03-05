// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/guilhermealvess/guicpay/domain/entity"
	gateway "github.com/guilhermealvess/guicpay/domain/gateway"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// AccountRepository is an autogenerated mock type for the AccountRepository type
type AccountRepository struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: ctx, account
func (_m *AccountRepository) CreateAccount(ctx context.Context, account entity.Account) error {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Account) error); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAccount provides a mock function with given fields: ctx, accountID
func (_m *AccountRepository) FindAccount(ctx context.Context, accountID uuid.UUID) (*entity.Account, error) {
	ret := _m.Called(ctx, accountID)

	if len(ret) == 0 {
		panic("no return value specified for FindAccount")
	}

	var r0 *entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*entity.Account, error)); ok {
		return rf(ctx, accountID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entity.Account); ok {
		r0 = rf(ctx, accountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountByEmail provides a mock function with given fields: ctx, email
func (_m *AccountRepository) FindAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for FindAccountByEmail")
	}

	var r0 *entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Account, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Account); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountByIDs provides a mock function with given fields: ctx, ids
func (_m *AccountRepository) FindAccountByIDs(ctx context.Context, ids ...uuid.UUID) (map[uuid.UUID]*entity.Account, error) {
	_va := make([]interface{}, len(ids))
	for _i := range ids {
		_va[_i] = ids[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindAccountByIDs")
	}

	var r0 map[uuid.UUID]*entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...uuid.UUID) (map[uuid.UUID]*entity.Account, error)); ok {
		return rf(ctx, ids...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...uuid.UUID) map[uuid.UUID]*entity.Account); ok {
		r0 = rf(ctx, ids...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uuid.UUID]*entity.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...uuid.UUID) error); ok {
		r1 = rf(ctx, ids...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: ctx
func (_m *AccountRepository) FindAll(ctx context.Context) ([]*entity.Account, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []*entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*entity.Account, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.Account); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindResumeAccount provides a mock function with given fields: ctx, email
func (_m *AccountRepository) FindResumeAccount(ctx context.Context, email string) (*entity.ResumeAccount, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for FindResumeAccount")
	}

	var r0 *entity.ResumeAccount
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.ResumeAccount, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.ResumeAccount); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.ResumeAccount)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTransaction provides a mock function with given fields: ctx
func (_m *AccountRepository) NewTransaction(ctx context.Context) (gateway.Tx, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for NewTransaction")
	}

	var r0 gateway.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (gateway.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) gateway.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gateway.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAtomicTransactions provides a mock function with given fields: ctx, transactions
func (_m *AccountRepository) SaveAtomicTransactions(ctx context.Context, transactions ...entity.Transaction) error {
	_va := make([]interface{}, len(transactions))
	for _i := range transactions {
		_va[_i] = transactions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SaveAtomicTransactions")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...entity.Transaction) error); ok {
		r0 = rf(ctx, transactions...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSnapshotTransactions provides a mock function with given fields: ctx, snapshotID, transactionIDs
func (_m *AccountRepository) SetSnapshotTransactions(ctx context.Context, snapshotID uuid.UUID, transactionIDs uuid.UUIDs) error {
	ret := _m.Called(ctx, snapshotID, transactionIDs)

	if len(ret) == 0 {
		panic("no return value specified for SetSnapshotTransactions")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUIDs) error); ok {
		r0 = rf(ctx, snapshotID, transactionIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAccountRepository creates a new instance of AccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountRepository {
	mock := &AccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
