package server

import (
	"context"
	"fmt"

	"github.com/xackery/nf/db"
	accountsv1 "github.com/xackery/nf/proto/accounts/v1"
	"github.com/xackery/nf/raw"
)

// CreateAccount adds a account to the in-memory store.
func (b *Backend) CreateAccount(ctx context.Context, req *accountsv1.CreateAccountRequest) (*accountsv1.CreateAccountResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if req.Account == nil {
		return nil, fmt.Errorf("account is nil")
	}
	account := &raw.Account{}
	account.Read(req.Account)

	account, err := db.AccountCreate(ctx, account)
	if err != nil {
		return nil, fmt.Errorf("account ability create: %w", err)
	}

	resp := &accountsv1.CreateAccountResponse{
		Account: account.Write(),
	}

	return resp, nil
}

// ListAccounts lists all accounts in the store.
func (b *Backend) ListAccounts(req *accountsv1.ListAccountsRequest, srv accountsv1.AccountService_ListAccountsServer) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rawAccounts, err := db.AccountList(ctx)
	if err != nil {
		return fmt.Errorf("account list: %w", err)
	}

	for _, rawAccount := range rawAccounts {
		err = srv.Send(&accountsv1.ListAccountsResponse{
			Account: rawAccount.Write(),
		})
		if err != nil {
			return fmt.Errorf("list account send: %w", err)
		}
	}

	return nil
}

// GetAccount gets a account from the store.
func (b *Backend) GetAccount(ctx context.Context, req *accountsv1.GetAccountRequest) (*accountsv1.GetAccountResponse, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	account, err := db.Account(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("account ability: %w", err)
	}

	resp := &accountsv1.GetAccountResponse{
		Account: account.Write(),
	}

	return resp, nil
}

// UpdateAccount updates a account in the store.
func (b *Backend) UpdateAccount(ctx context.Context, req *accountsv1.UpdateAccountRequest) (*accountsv1.UpdateAccountResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if req.Account == nil {
		return nil, fmt.Errorf("account is nil")
	}
	account := &raw.Account{}
	account.Read(req.Account)
	account.ID = int(req.Id)

	err := db.AccountUpdate(ctx, account)
	if err != nil {
		return nil, fmt.Errorf("account ability update: %w", err)
	}
	resp := &accountsv1.UpdateAccountResponse{
		Account: account.Write(),
	}

	return resp, nil
}

// DeleteAccount deletes a account from the store.
func (b *Backend) DeleteAccount(ctx context.Context, req *accountsv1.DeleteAccountRequest) (*accountsv1.DeleteAccountResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	err := db.AccountDelete(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("account ability delete: %w", err)
	}

	resp := &accountsv1.DeleteAccountResponse{}
	return resp, nil
}
