package server

import (
	"context"
	"fmt"

	"github.com/xackery/nf/db"
	adventuresv1 "github.com/xackery/nf/proto/adventures/v1"
	"github.com/xackery/nf/raw"
)

// CreateAdventure adds a adventure to the in-memory store.
func (b *Backend) CreateAdventure(ctx context.Context, req *adventuresv1.CreateAdventureRequest) (*adventuresv1.CreateAdventureResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if req.Adventure == nil {
		return nil, fmt.Errorf("adventure is nil")
	}
	adventure := &raw.Adventure{}
	adventure.Read(req.Adventure)

	adventure, err := db.AdventureCreate(ctx, adventure)
	if err != nil {
		return nil, fmt.Errorf("adventure ability create: %w", err)
	}

	resp := &adventuresv1.CreateAdventureResponse{
		Adventure: adventure.Write(),
	}

	return resp, nil
}

// ListAdventures lists all adventures in the store.
func (b *Backend) ListAdventures(req *adventuresv1.ListAdventuresRequest, srv adventuresv1.AdventureService_ListAdventuresServer) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rawAdventures, err := db.AdventureList(ctx)
	if err != nil {
		return fmt.Errorf("adventure list: %w", err)
	}

	for _, rawAdventure := range rawAdventures {
		err = srv.Send(&adventuresv1.ListAdventuresResponse{
			Adventure: rawAdventure.Write(),
		})
		if err != nil {
			return fmt.Errorf("list adventure send: %w", err)
		}
	}

	return nil
}

// GetAdventure gets a adventure from the store.
func (b *Backend) GetAdventure(ctx context.Context, req *adventuresv1.GetAdventureRequest) (*adventuresv1.GetAdventureResponse, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	adventure, err := db.Adventure(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("adventure ability: %w", err)
	}

	resp := &adventuresv1.GetAdventureResponse{
		Adventure: adventure.Write(),
	}

	return resp, nil
}

// UpdateAdventure updates a adventure in the store.
func (b *Backend) UpdateAdventure(ctx context.Context, req *adventuresv1.UpdateAdventureRequest) (*adventuresv1.UpdateAdventureResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if req.Adventure == nil {
		return nil, fmt.Errorf("adventure is nil")
	}
	adventure := &raw.Adventure{}
	adventure.Read(req.Adventure)
	adventure.ID = int(req.Id)

	err := db.AdventureUpdate(ctx, adventure)
	if err != nil {
		return nil, fmt.Errorf("adventure ability update: %w", err)
	}
	resp := &adventuresv1.UpdateAdventureResponse{
		Adventure: adventure.Write(),
	}

	return resp, nil
}

// DeleteAdventure deletes a adventure from the store.
func (b *Backend) DeleteAdventure(ctx context.Context, req *adventuresv1.DeleteAdventureRequest) (*adventuresv1.DeleteAdventureResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	err := db.AdventureDelete(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("adventure ability delete: %w", err)
	}

	resp := &adventuresv1.DeleteAdventureResponse{}
	return resp, nil
}
