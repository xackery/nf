package raw

import (
	adventuresv1 "github.com/xackery/nf/proto/adventures/v1"
)

type Adventure struct {
	ID              int     `db:"id"`
	Zone            string  `db:"zone"`
	ZoneVersion     int     `db:"zone_version"`
	IsHard          bool    `db:"is_hard"`
	IsRaid          bool    `db:"is_raid"`
	MinLevel        int     `db:"min_level"`
	MaxLevel        int     `db:"max_level"`
	Type            int     `db:"type"`
	TypeData        int     `db:"type_data"`
	TypeCount       int     `db:"type_count"`
	AssaX           float64 `db:"assa_x"`
	AssaY           float64 `db:"assa_y"`
	AssaZ           float64 `db:"assa_z"`
	AssaH           float64 `db:"assa_h"`
	Text            string  `db:"text"`
	Duration        int     `db:"duration"`
	ZoneInTime      int     `db:"zone_in_time"`
	WinPoints       int     `db:"win_points"`
	LosePoints      int     `db:"lose_points"`
	Theme           int     `db:"theme"`
	ZoneInZoneID    int     `db:"zone_in_zone_id"`
	ZoneInX         float64 `db:"zone_in_x"`
	ZoneInY         float64 `db:"zone_in_y"`
	ZoneInObjID     int     `db:"zone_in_object_id"`
	DestX           float64 `db:"dest_x"`
	DestY           float64 `db:"dest_y"`
	DestZ           float64 `db:"dest_z"`
	DestH           float64 `db:"dest_h"`
	GraveyardZoneID int     `db:"graveyard_zone_id"`
	GraveyardX      float64 `db:"graveyard_x"`
	GraveyardY      float64 `db:"graveyard_y"`
	GraveyardZ      float64 `db:"graveyard_z"`
	GraveyardRadius float64 `db:"graveyard_radius"`
	Entries         []*AdventureEntry
}

func (e *Adventure) Read(pb *adventuresv1.Adventure) {
	pb.Id = int32(e.ID)
	pb.Zone = e.Zone
	pb.ZoneVersion = int32(e.ZoneVersion)
	pb.IsHard = e.IsHard
	pb.IsRaid = e.IsRaid
	pb.MinLevel = int32(e.MinLevel)
	pb.MaxLevel = int32(e.MaxLevel)
	pb.Type = int32(e.Type)
	pb.TypeData = int32(e.TypeData)
	pb.TypeCount = int32(e.TypeCount)
	pb.AssaX = float32(e.AssaX)
	pb.AssaY = float32(e.AssaY)
	pb.AssaZ = float32(e.AssaZ)
	pb.AssaH = float32(e.AssaH)
	pb.Text = e.Text
	pb.Duration = int32(e.Duration)
	pb.ZoneInTime = int32(e.ZoneInTime)
	pb.WinPoints = int32(e.WinPoints)
	pb.LosePoints = int32(e.LosePoints)
	pb.Theme = int32(e.Theme)
	pb.ZoneInZoneId = int32(e.ZoneInZoneID)
	pb.ZoneInX = int32(e.ZoneInX)
	pb.ZoneInY = int32(e.ZoneInY)
	pb.ZoneInObjectId = int32(e.ZoneInObjID)
	pb.DestX = int32(e.DestX)
	pb.DestY = int32(e.DestY)
	pb.DestZ = int32(e.DestZ)
	pb.DestH = int32(e.DestH)
	pb.GraveyardZoneId = int32(e.GraveyardZoneID)
	pb.GraveyardX = int32(e.GraveyardX)
	pb.GraveyardY = int32(e.GraveyardY)
	pb.GraveyardZ = int32(e.GraveyardZ)
	pb.GraveyardRadius = int32(e.GraveyardRadius)
}

func (e *Adventure) Write() *adventuresv1.Adventure {
	return &adventuresv1.Adventure{
		Id:              int32(e.ID),
		Zone:            e.Zone,
		ZoneVersion:     int32(e.ZoneVersion),
		IsHard:          e.IsHard,
		IsRaid:          e.IsRaid,
		MinLevel:        int32(e.MinLevel),
		MaxLevel:        int32(e.MaxLevel),
		Type:            int32(e.Type),
		TypeData:        int32(e.TypeData),
		TypeCount:       int32(e.TypeCount),
		AssaX:           float32(e.AssaX),
		AssaY:           float32(e.AssaY),
		AssaZ:           float32(e.AssaZ),
		AssaH:           float32(e.AssaH),
		Text:            e.Text,
		Duration:        int32(e.Duration),
		ZoneInTime:      int32(e.ZoneInTime),
		WinPoints:       int32(e.WinPoints),
		LosePoints:      int32(e.LosePoints),
		Theme:           int32(e.Theme),
		ZoneInZoneId:    int32(e.ZoneInZoneID),
		ZoneInX:         int32(e.ZoneInX),
		ZoneInY:         int32(e.ZoneInY),
		ZoneInObjectId:  int32(e.ZoneInObjID),
		DestX:           int32(e.DestX),
		DestY:           int32(e.DestY),
		DestZ:           int32(e.DestZ),
		DestH:           int32(e.DestH),
		GraveyardZoneId: int32(e.GraveyardZoneID),
		GraveyardX:      int32(e.GraveyardX),
		GraveyardY:      int32(e.GraveyardY),
		GraveyardZ:      int32(e.GraveyardZ),
		GraveyardRadius: int32(e.GraveyardRadius),
	}
}
