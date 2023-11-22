package db

import (
	"context"
	"fmt"

	"github.com/xackery/nf/query"
	"github.com/xackery/nf/raw"
	"github.com/xackery/nf/validate"
)

// AccountCreate creates a aa ability
func AccountCreate(ctx context.Context, aa *raw.Account) (*raw.Account, error) {
	err := validate.Against(aa)
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	resp, err := conn.NamedExecContext(ctx, query.AccountCreate, aa)
	if err != nil {
		return nil, fmt.Errorf("insert: %w", err)
	}
	id, err := resp.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("last insert id: %w", err)
	}
	aa.ID = int(id)
	return aa, nil
}

// AccountList lists all aa abilities
func AccountList(ctx context.Context) ([]*raw.Account, error) {
	accounts := []*raw.Account{}
	err := conn.SelectContext(ctx, &accounts, query.AccountList)
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	return accounts, nil
}

// Account returns a aa ability by id
func Account(ctx context.Context, id int32) (*raw.Account, error) {
	account := &raw.Account{}
	err := conn.GetContext(ctx, account, query.Account, id)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return account, nil
}

// AccountUpdate updates a aa ability by id
func AccountUpdate(ctx context.Context, aa *raw.Account) error {
	err := validate.Against(aa)
	if err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	resp, err := conn.NamedExecContext(ctx, query.AccountUpdate, aa)
	if err != nil {
		return fmt.Errorf("update aa_ability: %w", err)
	}
	rows, err := resp.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if rows < 1 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

// AccountDelete deletes a aa ability by id
func AccountDelete(ctx context.Context, id int32) error {
	resp, err := conn.ExecContext(ctx, query.AccountDelete, id)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	rows, err := resp.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if rows < 1 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}
