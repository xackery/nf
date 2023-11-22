package db

import (
	"context"
	"fmt"

	"github.com/xackery/nf/query"
	"github.com/xackery/nf/raw"
	"github.com/xackery/nf/validate"
)

// AdventureEntryCreate creates a aa ability
func AdventureEntryCreate(ctx context.Context, aa *raw.AdventureEntry) (*raw.AdventureEntry, error) {
	err := validate.Against(aa)
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	resp, err := conn.NamedExecContext(ctx, query.AdventureEntryCreate, aa)
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

// AdventureEntryList lists all aa abilities
func AdventureEntryList(ctx context.Context, adventureID int32) ([]*raw.AdventureEntry, error) {
	adventureEntrys := []*raw.AdventureEntry{}
	err := conn.SelectContext(ctx, &adventureEntrys, query.AdventureEntryList, adventureID)
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	return adventureEntrys, nil
}

// AdventureEntry returns a aa ability by id
func AdventureEntry(ctx context.Context, id int32) (*raw.AdventureEntry, error) {
	adventureEntry := &raw.AdventureEntry{}
	err := conn.GetContext(ctx, adventureEntry, query.AdventureEntry, id)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return adventureEntry, nil
}

// AdventureEntryUpdate updates a aa ability by id
func AdventureEntryUpdate(ctx context.Context, aa *raw.AdventureEntry) error {
	err := validate.Against(aa)
	if err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	resp, err := conn.NamedExecContext(ctx, query.AdventureEntryUpdate, aa)
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

// AdventureEntryDelete deletes a aa ability by id
func AdventureEntryDelete(ctx context.Context, id int32) error {
	resp, err := conn.ExecContext(ctx, query.AdventureEntryDelete, id)
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
