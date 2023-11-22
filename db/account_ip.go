package db

import (
	"context"
	"fmt"

	"github.com/xackery/nf/raw"
)

func AccountIPList(ctx context.Context, accountID int32) ([]*raw.AccountIP, error) {
	accountIPs := []*raw.AccountIP{}
	rows, err := conn.QueryxContext(ctx, "SELECT * FROM account_ip")
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		accountIP := raw.AccountIP{}
		err = rows.StructScan(&accountIP)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		accountIPs = append(accountIPs, &accountIP)
	}

	return accountIPs, nil
}
