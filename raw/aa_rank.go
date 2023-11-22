package raw

import (
	ranksv1 "github.com/xackery/nf/proto/aas/ranks/v1"
)

// AARank is a representation of the aa_ability table.
type AARank struct {
	ID             int32 `db:"id"`
	UpperHotkeySID int32 `db:"upper_hotkey_sid"`
	LowerHotkeySID int32 `db:"lower_hotkey_sid"`
	TitleSID       int32 `db:"title_sid"`
	DescSID        int32 `db:"desc_sid"`
	Cost           int32 `db:"cost"`
	LevelReq       int32 `db:"level_req"`
	Spell          int32 `db:"spell"`
	SpellType      int32 `db:"spell_type"`
	RecastTime     int32 `db:"recast_time"`
	Expansion      int32 `db:"expansion"`
	PrevID         int32 `db:"prev_id"`
	NextID         int32 `db:"next_id"`
	Effects        []*AARankEffect
	Prereqs        []*AARankPrereq
}

func (e *AARank) Read(pb *ranksv1.Rank) {
	if pb == nil {
		return
	}
	e.ID = pb.Id
	e.UpperHotkeySID = pb.UpperHotkeySid
	e.LowerHotkeySID = pb.LowerHotkeySid
	e.TitleSID = pb.TitleSid
	e.DescSID = pb.DescSid
	e.Cost = pb.Cost
	e.LevelReq = pb.LevelReq
	e.Spell = pb.Spell
	e.SpellType = pb.SpellType
	e.RecastTime = pb.RecastTime
	e.Expansion = pb.Expansion
	e.PrevID = pb.PrevId
	e.NextID = pb.NextId
	for _, effect := range pb.Effects {
		rawEffect := &AARankEffect{}
		rawEffect.Read(effect)
		e.Effects = append(e.Effects, rawEffect)
	}
	for _, prereq := range pb.Prereqs {
		rawPrereq := &AARankPrereq{}
		rawPrereq.Read(prereq)
		e.Prereqs = append(e.Prereqs, rawPrereq)
	}
}

func (e *AARank) Write() *ranksv1.Rank {
	pb := &ranksv1.Rank{
		Id:             e.ID,
		UpperHotkeySid: e.UpperHotkeySID,
		LowerHotkeySid: e.LowerHotkeySID,
		TitleSid:       e.TitleSID,
		DescSid:        e.DescSID,
		Cost:           e.Cost,
		LevelReq:       e.LevelReq,
		Spell:          e.Spell,
		SpellType:      e.SpellType,
		RecastTime:     e.RecastTime,
		Expansion:      e.Expansion,
		PrevId:         e.PrevID,
		NextId:         e.NextID,
	}
	for _, effect := range e.Effects {
		pb.Effects = append(pb.Effects, effect.Write())
	}
	for _, prereq := range e.Prereqs {
		pb.Prereqs = append(pb.Prereqs, prereq.Write())
	}
	return pb
}
