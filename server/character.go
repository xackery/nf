package server

import (
	"context"
	"fmt"

	"github.com/xackery/nf/db"
	charactersv1 "github.com/xackery/nf/proto/characters/v1"
	"github.com/xackery/nf/raw"
)

// CreateCharacter adds a character to the in-memory store.
func (b *Backend) CreateCharacter(ctx context.Context, req *charactersv1.CreateCharacterRequest) (*charactersv1.CreateCharacterResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if req.Character == nil {
		return nil, fmt.Errorf("character is nil")
	}
	character := &raw.Character{}
	character.Read(req.Character)

	character, err := db.CharacterCreate(ctx, character)
	if err != nil {
		return nil, fmt.Errorf("character ability create: %w", err)
	}

	resp := &charactersv1.CreateCharacterResponse{
		Character: character.Write(),
	}

	return resp, nil
}

// ListCharacters lists all characters in the store.
func (b *Backend) ListCharacters(req *charactersv1.ListCharactersRequest, srv charactersv1.CharacterService_ListCharactersServer) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rawCharacters, err := db.CharacterList(ctx)
	if err != nil {
		return fmt.Errorf("character list: %w", err)
	}

	for _, rawCharacter := range rawCharacters {
		err = srv.Send(&charactersv1.ListCharactersResponse{
			Character: rawCharacter.Write(),
		})
		if err != nil {
			return fmt.Errorf("list character send: %w", err)
		}
	}

	return nil
}

// GetCharacter gets a character from the store.
func (b *Backend) GetCharacter(ctx context.Context, req *charactersv1.GetCharacterRequest) (*charactersv1.GetCharacterResponse, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	character, err := db.Character(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("character ability: %w", err)
	}

	resp := &charactersv1.GetCharacterResponse{
		Character: character.Write(),
	}

	return resp, nil
}

// UpdateCharacter updates a character in the store.
func (b *Backend) UpdateCharacter(ctx context.Context, req *charactersv1.UpdateCharacterRequest) (*charactersv1.UpdateCharacterResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if req.Character == nil {
		return nil, fmt.Errorf("character is nil")
	}
	character := &raw.Character{}
	character.Read(req.Character)
	character.ID = int(req.Id)

	err := db.CharacterUpdate(ctx, character)
	if err != nil {
		return nil, fmt.Errorf("character ability update: %w", err)
	}
	resp := &charactersv1.UpdateCharacterResponse{
		Character: character.Write(),
	}

	return resp, nil
}

// DeleteCharacter deletes a character from the store.
func (b *Backend) DeleteCharacter(ctx context.Context, req *charactersv1.DeleteCharacterRequest) (*charactersv1.DeleteCharacterResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	err := db.CharacterDelete(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("character ability delete: %w", err)
	}

	resp := &charactersv1.DeleteCharacterResponse{}
	return resp, nil
}
