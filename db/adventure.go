package db

import (
	"context"
	"fmt"

	"github.com/xackery/nf/query"
	"github.com/xackery/nf/raw"
	"github.com/xackery/nf/validate"
)

// AdventureCreate creates a aa ability
func AdventureCreate(ctx context.Context, aa *raw.Adventure) (*raw.Adventure, error) {
	err := validate.Against(aa)
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	resp, err := conn.NamedExecContext(ctx, query.AdventureCreate, aa)
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

// AdventureList lists all aa abilities
func AdventureList(ctx context.Context) ([]*raw.Adventure, error) {
	adventures := []*raw.Adventure{}
	err := conn.SelectContext(ctx, &adventures, query.AdventureList)
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	return adventures, nil
}

// Adventure returns a aa ability by id
func Adventure(ctx context.Context, id int32) (*raw.Adventure, error) {
	adventure := &raw.Adventure{}
	err := conn.GetContext(ctx, adventure, query.Adventure, id)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	adventure.Entries, err = AdventureEntryList(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("adventure entry list: %w", err)
	}
	return adventure, nil
}

// AdventureUpdate updates a aa ability by id
func AdventureUpdate(ctx context.Context, aa *raw.Adventure) error {
	err := validate.Against(aa)
	if err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	resp, err := conn.NamedExecContext(ctx, query.AdventureUpdate, aa)
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

// AdventureDelete deletes a aa ability by id
func AdventureDelete(ctx context.Context, id int32) error {
	resp, err := conn.ExecContext(ctx, query.AdventureDelete, id)
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
