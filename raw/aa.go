package raw

import aasv1 "github.com/xackery/nf/proto/aas/v1"

// AAAbility is a representation of the aa_ability table.
type AAAbility struct {
	ID              int32  `db:"id"`
	Name            string `db:"name"`
	Category        int32  `db:"category"`
	Classes         int32  `db:"classes"`
	Races           int32  `db:"races"`
	DrakkinHeritage int32  `db:"drakkin_heritage"`
	Deities         int32  `db:"deities"`
	Status          int32  `db:"status"`
	Type            int32  `db:"type"`
	Charges         int32  `db:"charges"`
	GrantOnly       int32  `db:"grant_only"`
	FirstRankID     int32  `db:"first_rank_id"`
	Enabled         int32  `db:"enabled"`
	ResetOnDeath    int32  `db:"reset_on_death"`
}

func (e *AAAbility) Read(pb *aasv1.Aa) {
	if pb == nil {
		return
	}
	e.ID = pb.Id
	e.Name = pb.Name
	e.Category = pb.Category
	e.Classes = pb.Classes
	e.Races = pb.Races
	e.DrakkinHeritage = pb.DrakkinHeritage
	e.Deities = pb.Deities
	e.Status = pb.Status
	e.Type = pb.Type
	e.Charges = pb.Charges
	if pb.GrantOnly {
		e.GrantOnly = 1
	}
	e.FirstRankID = pb.FirstRankId
	if pb.Enabled {
		e.Enabled = 1
	}
	if pb.ResetOnDeath {
		e.ResetOnDeath = 1
	}
}

func (e *AAAbility) Write() *aasv1.Aa {
	return &aasv1.Aa{
		Id:              e.ID,
		Name:            e.Name,
		Category:        e.Category,
		Classes:         e.Classes,
		Races:           e.Races,
		DrakkinHeritage: e.DrakkinHeritage,
		Deities:         e.Deities,
		Status:          e.Status,
		Type:            e.Type,
		Charges:         e.Charges,
		GrantOnly:       e.GrantOnly == 1,
		FirstRankId:     e.FirstRankID,
		Enabled:         e.Enabled == 1,
		ResetOnDeath:    e.ResetOnDeath == 1,
	}
}
