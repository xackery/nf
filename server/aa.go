package server

import (
	"context"
	"fmt"

	"github.com/xackery/nf/db"
	aasv1 "github.com/xackery/nf/proto/aas/v1"
	"github.com/xackery/nf/raw"
)

// CreateAa adds a aa to the in-memory store.
func (b *Backend) CreateAa(ctx context.Context, req *aasv1.CreateAaRequest) (*aasv1.CreateAaResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if req.Aa == nil {
		return nil, fmt.Errorf("aa is nil")
	}
	aa := &raw.AAAbility{}
	aa.Read(req.Aa)

	aa, err := db.AAAbilityCreate(ctx, aa)
	if err != nil {
		return nil, fmt.Errorf("aa ability create: %w", err)
	}

	resp := &aasv1.CreateAaResponse{
		Aa: aa.Write(),
	}

	return resp, nil
}

// ListAas lists all aas in the store.
func (b *Backend) ListAas(_ *aasv1.ListAasRequest, srv aasv1.AaService_ListAasServer) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, aa := range b.aas {
		err := srv.Send(&aasv1.ListAasResponse{Aa: aa})
		if err != nil {
			return err
		}
	}

	return nil
}

// GetAa gets a aa from the store.
func (b *Backend) GetAa(ctx context.Context, req *aasv1.GetAaRequest) (*aasv1.GetAaResponse, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	aa, err := db.AAAbility(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("aa ability: %w", err)
	}

	resp := &aasv1.GetAaResponse{
		Aa: aa.Write(),
	}

	return resp, nil
}

// UpdateAa updates a aa in the store.
func (b *Backend) UpdateAa(ctx context.Context, req *aasv1.UpdateAaRequest) (*aasv1.UpdateAaResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if req.Aa == nil {
		return nil, fmt.Errorf("aa is nil")
	}
	aa := &raw.AAAbility{}
	aa.Read(req.Aa)
	aa.ID = req.Id

	err := db.AAAbilityUpdate(ctx, aa)
	if err != nil {
		return nil, fmt.Errorf("aa ability update: %w", err)
	}
	resp := &aasv1.UpdateAaResponse{
		Aa: aa.Write(),
	}

	return resp, nil
}

// DeleteAa deletes a aa from the store.
func (b *Backend) DeleteAa(ctx context.Context, req *aasv1.DeleteAaRequest) (*aasv1.DeleteAaResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	err := db.AAAbilityDelete(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("aa ability delete: %w", err)
	}

	resp := &aasv1.DeleteAaResponse{}
	return resp, nil
}
