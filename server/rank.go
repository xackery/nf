package server

import (
	"context"
	"fmt"

	"github.com/xackery/nf/db"
	ranksv1 "github.com/xackery/nf/proto/aas/ranks/v1"
	"github.com/xackery/nf/raw"
)

// CreateRank adds a rank to the in-memory store.
func (b *Backend) CreateRank(ctx context.Context, req *ranksv1.CreateRankRequest) (*ranksv1.CreateRankResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	rank := &raw.AARank{}
	rank.Read(req.Rank)

	rank, err := db.AARankCreate(ctx, rank)
	if err != nil {
		return nil, fmt.Errorf("aa rank create: %w", err)
	}

	resp := &ranksv1.CreateRankResponse{
		Rank: rank.Write(),
	}
	return resp, nil
}

// ListRanks lists all ranks in the store.
func (b *Backend) ListRanks(req *ranksv1.ListRanksRequest, srv ranksv1.RankService_ListRanksServer) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rawRanks, err := db.AARankList(ctx, req.Aaid)
	if err != nil {
		return fmt.Errorf("aa rank list: %w", err)
	}

	for _, rawRank := range rawRanks {
		err = srv.Send(&ranksv1.ListRanksResponse{
			Aaid: req.Aaid,
			Rank: rawRank.Write(),
		})
		if err != nil {
			return fmt.Errorf("list rank send: %w", err)
		}
	}

	return nil
}

// GetRank gets a rank from the store.
func (b *Backend) GetRank(ctx context.Context, req *ranksv1.GetRankRequest) (*ranksv1.GetRankResponse, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	rank, err := db.AARank(ctx, req.Aaid, req.Id)
	if err != nil {
		return nil, fmt.Errorf("aa rank: %w", err)
	}

	resp := &ranksv1.GetRankResponse{
		Aaid: req.Aaid,
		Rank: rank.Write(),
	}

	for _, effect := range rank.Effects {
		resp.Rank.Effects = append(resp.Rank.Effects, effect.Write())
	}

	return resp, nil
}

// UpdateRank updates a rank in the store.
func (b *Backend) UpdateRank(ctx context.Context, req *ranksv1.UpdateRankRequest) (*ranksv1.UpdateRankResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	rank := &raw.AARank{}
	rank.Read(req.Rank)

	rank, err := db.AARankUpdate(ctx, rank)
	if err != nil {
		return nil, fmt.Errorf("aa rank update: %w", err)
	}

	resp := &ranksv1.UpdateRankResponse{
		Rank: rank.Write(),
	}

	return resp, nil
}

// DeleteRank deletes a rank from the store.
func (b *Backend) DeleteRank(ctx context.Context, req *ranksv1.DeleteRankRequest) (*ranksv1.DeleteRankResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	err := db.AARankDelete(ctx, req.Aaid, req.Id)
	if err != nil {
		return nil, fmt.Errorf("aa rank delete: %w", err)
	}

	resp := &ranksv1.DeleteRankResponse{
		Aaid: req.Aaid,
		Id:   req.Id,
	}
	return resp, nil
}
