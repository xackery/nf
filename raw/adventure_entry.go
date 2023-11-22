package raw

import (
	adventuresv1 "github.com/xackery/nf/proto/adventures/v1"
)

type AdventureEntry struct {
	ID         int `db:"id"`
	TemplateID int `db:"template_id"`
}

func (e *AdventureEntry) Read(pb *adventuresv1.AdventureEntry) {
	if pb == nil {
		return
	}
	e.ID = int(pb.Id)
	e.TemplateID = int(pb.TemplateId)
}

func (e *AdventureEntry) Write() *adventuresv1.AdventureEntry {
	pb := &adventuresv1.AdventureEntry{
		Id:         int32(e.ID),
		TemplateId: int32(e.TemplateID),
	}
	return pb
}
