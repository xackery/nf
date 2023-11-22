package db

import (
	"context"
	"fmt"

	"github.com/xackery/nf/query"
	"github.com/xackery/nf/raw"
	"github.com/xackery/nf/validate"
)

// AAAbilityCreate creates a aa ability
func AAAbilityCreate(ctx context.Context, aa *raw.AAAbility) (*raw.AAAbility, error) {
	err := validate.Against(aa)
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	resp, err := conn.NamedExecContext(ctx, query.AAAbilityCreate, aa)
	if err != nil {
		return nil, fmt.Errorf("insert: %w", err)
	}
	id, err := resp.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("last insert id: %w", err)
	}
	aa.ID = int32(id)
	return aa, nil
}

// AAAbility returns a aa ability by id
func AAAbility(ctx context.Context, id int32) (*raw.AAAbility, error) {
	aaAbility := &raw.AAAbility{}
	err := conn.GetContext(ctx, aaAbility, query.AAAbility, id)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return aaAbility, nil
}

// AAAbilityUpdate updates a aa ability by id
func AAAbilityUpdate(ctx context.Context, aa *raw.AAAbility) error {
	err := validate.Against(aa)
	if err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	resp, err := conn.NamedExecContext(ctx, query.AAAbilityUpdate, aa)
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

// AAAbilityDelete deletes a aa ability by id
func AAAbilityDelete(ctx context.Context, id int32) error {
	resp, err := conn.ExecContext(ctx, query.AAAbilityDelete, id)
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
