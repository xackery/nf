package db

import (
	"context"
	"fmt"

	"github.com/xackery/nf/raw"
)

// AARankEffectList returns a list of ranks by aa id
func AARankEffectList(ctx context.Context, rankID int32) ([]*raw.AARankEffect, error) {
	aaRankEffects := []*raw.AARankEffect{}
	rows, err := conn.QueryxContext(ctx, "SELECT * FROM aa_rank_effects WHERE rank_id = ?", rankID)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		aaRankEffect := raw.AARankEffect{}
		err = rows.StructScan(&aaRankEffect)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		aaRankEffects = append(aaRankEffects, &aaRankEffect)
	}

	return aaRankEffects, nil
}
