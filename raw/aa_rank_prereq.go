package raw

import (
	ranksv1 "github.com/xackery/nf/proto/aas/ranks/v1"
)

type AARankPrereq struct {
	RankID int32 `json:"rank_id" db:"rank_id"`
	AaID   int32 `json:"aa_id" db:"aa_id"`
	Points int32 `json:"points" db:"points"`
}

func (e *AARankPrereq) Read(pb *ranksv1.RankPrereq) {
	if pb == nil {
		return
	}
	e.RankID = pb.RankId
	e.AaID = pb.AaId
	e.Points = pb.Points
}

func (e *AARankPrereq) Write() *ranksv1.RankPrereq {
	return &ranksv1.RankPrereq{
		RankId: e.RankID,
		AaId:   e.AaID,
		Points: e.Points,
	}
}
