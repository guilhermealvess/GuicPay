package entity

import (
	"time"

	"github.com/google/uuid"
)

type AccountType string

const (
	Personal AccountType = "PERSONAL"
	Seller   AccountType = "SELLER"
)

type AccountStatus string

const (
	AccountStatusActive   AccountStatus = "ACTIVE"
	AccountStatusCanceled AccountStatus = "CANCELED"
)

type Account struct {
	ID              uuid.UUID
	AccountType     AccountType
	CustomerName    string
	DocumentNUmber  string
	Email           string
	PasswordEncoded string
	Salt            string
	PhoneNumber     string
	Status          AccountStatus
	CreatedAt       time.Time
	UpdadatedAt     time.Time
	Wallet          Wallet
}

func NewAccount(t AccountType, name, doc, email, pass, phone string) Account {
	now := time.Now().UTC()
	return Account{
		ID:              uuid.New(),
		AccountType:     t,
		CustomerName:    name,
		DocumentNUmber:  doc,
		Email:           email,
		PasswordEncoded: pass + pass,
		PhoneNumber:     phone,
		Status:          AccountStatusActive,
		CreatedAt:       now,
		UpdadatedAt:     now,
		Wallet:          []Transaction{},
	}
}

func (a *Account) Deposit(v Money) (*Transaction, error) {
	if a.Status == AccountStatusCanceled {
		return nil, NewDepositError("account canceled cant deposit", a.ID, Money(v))
	}

	t := factoryDepositTransaction(*a, v)
	a.Wallet = append(a.Wallet, t)

	return &t, nil
}

func (a *Account) Transfer(payee *Account, v Money) (*TransferOutput, error) {
	if a.AccountType == Seller {
		return nil, NewTransferError("account seller cant make transfer", a.ID, Money(v))
	}

	if a.Wallet.Balance() < Money(v) {
		return nil, NewTransferError("insuficient balance", a.ID, Money(v))
	}

	t1, t2 := factoryTransferTransactions(*a, *payee, v)
	a.Wallet = append(a.Wallet, t1)
	payee.Wallet = append(payee.Wallet, t2)

	return &TransferOutput{
		Payer:        &t1,
		Payee:        &t2,
		CorrelatedID: t1.CorrelatedID.UUID,
	}, nil
}

type TransferOutput struct {
	Payer        *Transaction
	Payee        *Transaction
	CorrelatedID uuid.UUID
}
