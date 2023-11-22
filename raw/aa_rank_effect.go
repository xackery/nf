package raw

import (
	ranksv1 "github.com/xackery/nf/proto/aas/ranks/v1"
)

type AARankEffect struct {
	RankID   int32 `json:"rank_id" db:"rank_id"`
	Slot     int32 `json:"slot" db:"slot"`
	EffectID int32 `json:"effect_id" db:"effect_id"`
	Base1    int32 `json:"base1" db:"base1"`
	Base2    int32 `json:"base2" db:"base2"`
}

func (e *AARankEffect) Read(pb *ranksv1.RankEffect) {
	if pb == nil {
		return
	}
	e.RankID = pb.RankId
	e.Slot = pb.Slot
	e.EffectID = pb.EffectId
	e.Base1 = pb.Base1
	e.Base2 = pb.Base2
}

func (e *AARankEffect) Write() *ranksv1.RankEffect {
	return &ranksv1.RankEffect{
		RankId:   e.RankID,
		Slot:     e.Slot,
		EffectId: e.EffectID,
		Base1:    e.Base1,
		Base2:    e.Base2,
	}
}
