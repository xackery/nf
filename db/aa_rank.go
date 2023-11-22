package db

import (
	"context"
	"fmt"

	"github.com/xackery/nf/query"
	"github.com/xackery/nf/raw"
)

// AARankCreate creates a aa rank
func AARankCreate(ctx context.Context, aaRank *raw.AARank) (*raw.AARank, error) {
	tx, err := conn.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	resp, err := tx.NamedExecContext(ctx, query.AARankCreate, aaRank)
	if err != nil {
		return nil, fmt.Errorf("insert: %w", err)
	}
	id, err := resp.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("last insert id: %w", err)
	}
	aaRank.ID = int32(id)

	for _, prereq := range aaRank.Prereqs {
		prereq.RankID = aaRank.ID
		_, err = tx.NamedExecContext(ctx, query.AARankPrereqCreate, prereq)
		if err != nil {
			return nil, fmt.Errorf("prereq insert: %w", err)
		}
	}

	for _, effect := range aaRank.Effects {
		effect.RankID = aaRank.ID
		_, err = tx.NamedExecContext(ctx, query.AARankEffectCreate, effect)
		if err != nil {
			return nil, fmt.Errorf("effect insert: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("commit: %w", err)
	}

	return aaRank, nil
}

// AARank returns a rank by rank id
func AARank(ctx context.Context, aaID int32, rankID int32) (*raw.AARank, error) {
	aaRank := &raw.AARank{}
	err := conn.GetContext(ctx, aaRank, query.AARank, rankID)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	aaRankEffects, err := AARankEffectList(ctx, rankID)
	if err != nil {
		return nil, fmt.Errorf("aa rank effect list: %w", err)
	}
	aaRank.Effects = aaRankEffects

	aaRankPrereqs, err := AARankPrereqList(ctx, rankID)
	if err != nil {
		return nil, fmt.Errorf("aa rank prereq list: %w", err)
	}
	aaRank.Prereqs = aaRankPrereqs

	return aaRank, nil
}

// AARanks returns a list of ranks by aa id
func AARankList(ctx context.Context, aaID int32) ([]*raw.AARank, error) {
	aa, err := AAAbility(ctx, aaID)
	if err != nil {
		return nil, fmt.Errorf("aa ability: %w", err)
	}

	aaRanks := []*raw.AARank{}
	aaRank, err := AARank(ctx, aaID, aa.FirstRankID)
	if err != nil {
		return nil, fmt.Errorf("aa rank firstrankID: %w", err)
	}
	aaRankEffects, err := AARankEffectList(ctx, aaRank.ID)
	if err != nil {
		return nil, fmt.Errorf("aa rank effect list: %w", err)
	}
	aaRank.Effects = aaRankEffects
	aaRankPrereqs, err := AARankPrereqList(ctx, aaRank.ID)
	if err != nil {
		return nil, fmt.Errorf("aa rank prereq list: %w", err)
	}
	aaRank.Prereqs = aaRankPrereqs

	aaRanks = append(aaRanks, aaRank)

	for aaRank.NextID > 0 {
		aaRank, err = AARank(ctx, aaID, aaRank.NextID)
		if err != nil {
			return nil, fmt.Errorf("aa rank nextID: %w", err)
		}
		aaRankEffects, err := AARankEffectList(ctx, aaRank.ID)
		if err != nil {
			return nil, fmt.Errorf("aa rank effect list: %w", err)
		}
		aaRank.Effects = aaRankEffects

		aaRankPrereqs, err := AARankPrereqList(ctx, aaRank.ID)
		if err != nil {
			return nil, fmt.Errorf("aa rank prereq list: %w", err)
		}
		aaRank.Prereqs = aaRankPrereqs

		aaRanks = append(aaRanks, aaRank)
	}

	return aaRanks, nil
}

// AARankUpdate updates a aa rank
func AARankUpdate(ctx context.Context, aaRank *raw.AARank) (*raw.AARank, error) {
	tx, err := conn.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.NamedExecContext(ctx, query.AARankUpdate, aaRank)
	if err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}

	for _, prereq := range aaRank.Prereqs {
		prereq.RankID = aaRank.ID
		_, err = tx.NamedExecContext(ctx, query.AARankPrereqUpdate, prereq)
		if err != nil {
			return nil, fmt.Errorf("prereq update: %w", err)
		}
	}

	for _, effect := range aaRank.Effects {
		effect.RankID = aaRank.ID
		_, err = tx.NamedExecContext(ctx, query.AARankEffectUpdate, effect)
		if err != nil {
			return nil, fmt.Errorf("effect update: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("commit: %w", err)
	}

	return aaRank, nil
}

// AARankDelete deletes a aa rank
func AARankDelete(ctx context.Context, aaID int32, rankID int32) error {
	tx, err := conn.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, query.AARankDelete, rankID)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	_, err = tx.ExecContext(ctx, query.AARankPrereqDelete, rankID)
	if err != nil {
		return fmt.Errorf("prereq delete: %w", err)
	}

	_, err = tx.ExecContext(ctx, query.AARankEffectDelete, rankID)
	if err != nil {
		return fmt.Errorf("effect delete: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit: %w", err)
	}

	return nil
}
