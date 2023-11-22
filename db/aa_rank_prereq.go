package db

import (
	"context"
	"fmt"

	"github.com/xackery/nf/raw"
)

// AARankPrereqList returns a list of ranks by aa id
func AARankPrereqList(ctx context.Context, rankID int32) ([]*raw.AARankPrereq, error) {
	aaRankPrereqs := []*raw.AARankPrereq{}
	rows, err := conn.QueryxContext(ctx, "SELECT * FROM aa_rank_prereqs WHERE rank_id = ?", rankID)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		aaRankPrereq := raw.AARankPrereq{}
		err = rows.StructScan(&aaRankPrereq)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		aaRankPrereqs = append(aaRankPrereqs, &aaRankPrereq)
	}

	return aaRankPrereqs, nil
}
