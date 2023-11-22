package raw

import (
	"database/sql"
	"time"

	charactersv1 "github.com/xackery/nf/proto/characters/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Character represents a user character in the system.
type Character struct {
	ID                    uint32       `db:"id"`
	AccountID             uint32       `db:"account_id"`
	Name                  string       `db:"name"`
	LastName              string       `db:"last_name"`
	Title                 string       `db:"title"`
	Suffix                string       `db:"suffix"`
	ZoneID                uint32       `db:"zone_id"`
	ZoneInstance          uint32       `db:"zone_instance"`
	Y                     float32      `db:"y"`
	X                     float32      `db:"x"`
	Z                     float32      `db:"z"`
	Heading               float32      `db:"heading"`
	Gender                uint8        `db:"gender"`
	Race                  uint16       `db:"race"`
	Class                 uint8        `db:"class"`
	Level                 uint32       `db:"level"`
	Deity                 uint32       `db:"deity"`
	Birthday              uint32       `db:"birthday"`
	LastLogin             uint32       `db:"last_login"`
	TimePlayed            uint32       `db:"time_played"`
	Level2                uint8        `db:"level2"`
	Anon                  uint8        `db:"anon"`
	GM                    uint8        `db:"gm"`
	Face                  uint32       `db:"face"`
	HairColor             uint8        `db:"hair_color"`
	HairStyle             uint8        `db:"hair_style"`
	Beard                 uint8        `db:"beard"`
	BeardColor            uint8        `db:"beard_color"`
	EyeColor1             uint8        `db:"eye_color_1"`
	EyeColor2             uint8        `db:"eye_color_2"`
	DrakkinHeritage       uint32       `db:"drakkin_heritage"`
	DrakkinTattoo         uint32       `db:"drakkin_tattoo"`
	DrakkinDetails        uint32       `db:"drakkin_details"`
	AbilityTimeSeconds    uint8        `db:"ability_time_seconds"`
	AbilityNumber         uint8        `db:"ability_number"`
	AbilityTimeMinutes    uint8        `db:"ability_time_minutes"`
	AbilityTimeHours      uint8        `db:"ability_time_hours"`
	Exp                   uint32       `db:"exp"`
	ExpEnabled            uint8        `db:"exp_enabled"`
	AAPointsSpent         uint32       `db:"aa_points_spent"`
	AAExp                 uint32       `db:"aa_exp"`
	AAPoints              uint32       `db:"aa_points"`
	GroupLeadershipExp    uint32       `db:"group_leadership_exp"`
	RaidLeadershipExp     uint32       `db:"raid_leadership_exp"`
	GroupLeadershipPoints uint32       `db:"group_leadership_points"`
	RaidLeadershipPoints  uint32       `db:"raid_leadership_points"`
	Points                uint32       `db:"points"`
	CurHP                 uint32       `db:"cur_hp"`
	Mana                  uint32       `db:"mana"`
	Endurance             uint32       `db:"endurance"`
	Intoxication          uint32       `db:"intoxication"`
	STR                   uint32       `db:"str"`
	STA                   uint32       `db:"sta"`
	CHA                   uint32       `db:"cha"`
	DEX                   uint32       `db:"dex"`
	INT                   uint32       `db:"int"`
	AGI                   uint32       `db:"agi"`
	WIS                   uint32       `db:"wis"`
	ZoneChangeCount       uint32       `db:"zone_change_count"`
	Toxicity              uint32       `db:"toxicity"`
	HungerLevel           uint32       `db:"hunger_level"`
	ThirstLevel           uint32       `db:"thirst_level"`
	AbilityUp             uint32       `db:"ability_up"`
	LdonPointsGuk         uint32       `db:"ldon_points_guk"`
	LdonPointsMir         uint32       `db:"ldon_points_mir"`
	LdonPointsMmc         uint32       `db:"ldon_points_mmc"`
	LdonPointsRuj         uint32       `db:"ldon_points_ruj"`
	LdonPointsTak         uint32       `db:"ldon_points_tak"`
	LdonPointsAvailable   uint32       `db:"ldon_points_available"`
	TributeTimeRemaining  uint32       `db:"tribute_time_remaining"`
	CareerTributePoints   uint32       `db:"career_tribute_points"`
	TributePoints         uint32       `db:"tribute_points"`
	TributeActive         uint32       `db:"tribute_active"`
	PvPStatus             uint8        `db:"pvp_status"`
	PvPKills              uint32       `db:"pvp_kills"`
	PvPDeaths             uint32       `db:"pvp_deaths"`
	PvPCurrentPoints      uint32       `db:"pvp_current_points"`
	PvPCareerPoints       uint32       `db:"pvp_career_points"`
	PvPBestKillStreak     uint32       `db:"pvp_best_kill_streak"`
	PvPWorstDeathStreak   uint32       `db:"pvp_worst_death_streak"`
	PvPCurrentKillStreak  uint32       `db:"pvp_current_kill_streak"`
	PvP2                  uint32       `db:"pvp2"`
	PvPType               uint32       `db:"pvp_type"`
	ShowHelm              uint32       `db:"show_helm"`
	GroupAutoConsent      uint8        `db:"group_auto_consent"`
	RaidAutoConsent       uint8        `db:"raid_auto_consent"`
	GuildAutoConsent      uint8        `db:"guild_auto_consent"`
	LeadershipExpOn       uint8        `db:"leadership_exp_on"`
	RestTimer             uint32       `db:"RestTimer"`
	AirRemaining          uint32       `db:"air_remaining"`
	AutosplitEnabled      uint8        `db:"autosplit_enabled"`
	Lfp                   uint8        `db:"lfp"`
	Lfg                   uint8        `db:"lfg"`
	Mailkey               string       `db:"mailkey"`
	Xtargets              uint32       `db:"xtargets"`
	Firstlogon            uint32       `db:"firstlogon"`
	EAaEffects            uint32       `db:"e_aa_effects"`
	EPercentToAA          uint32       `db:"e_percent_to_aa"`
	EExpendedAaSpent      uint32       `db:"e_expended_aa_spent"`
	AAPointsSpentOld      uint32       `db:"aa_points_spent_old"`
	AAPointsOld           uint32       `db:"aa_points_old"`
	ELastInvsnapshot      uint32       `db:"e_last_invsnapshot"`
	DeletedAt             sql.NullTime `db:"deleted_at"`
}

func (e *Character) Read(pb *charactersv1.Character) {
	if pb == nil {
		return
	}
	e.ID = uint32(pb.Id)
	e.Name = pb.Name
	e.LastName = pb.LastName
	e.Title = pb.Title
	e.Suffix = pb.Suffix
	e.ZoneID = uint32(pb.ZoneId)
	e.ZoneInstance = uint32(pb.ZoneInstance)
	e.Y = pb.Y
	e.X = pb.X
	e.Z = pb.Z
	e.Heading = pb.Heading
	e.Gender = uint8(pb.Gender)
	e.Race = uint16(pb.Race)
	e.Class = uint8(pb.Class)
	e.Level = uint32(pb.Level)
	e.Deity = uint32(pb.Deity)
	e.Birthday = uint32(pb.Birthday)
	e.LastLogin = uint32(pb.LastLogin)
	e.TimePlayed = uint32(pb.TimePlayed)
	e.Level2 = uint8(pb.Level2)
	e.Anon = uint8(pb.Anon)
	e.GM = uint8(pb.Gm)
	e.Face = uint32(pb.Face)
	e.HairColor = uint8(pb.HairColor)
	e.HairStyle = uint8(pb.HairStyle)
	e.Beard = uint8(pb.Beard)
	e.BeardColor = uint8(pb.BeardColor)
	e.EyeColor1 = uint8(pb.EyeColor_1)
	e.EyeColor2 = uint8(pb.EyeColor_2)
	e.DrakkinHeritage = uint32(pb.DrakkinHeritage)
	e.DrakkinTattoo = uint32(pb.DrakkinTattoo)
	e.DrakkinDetails = uint32(pb.DrakkinDetails)
	e.AbilityTimeSeconds = uint8(pb.AbilityTimeSeconds)
	e.AbilityNumber = uint8(pb.AbilityNumber)
	e.AbilityTimeMinutes = uint8(pb.AbilityTimeMinutes)
	e.AbilityTimeHours = uint8(pb.AbilityTimeHours)
	e.Exp = uint32(pb.Exp)
	e.ExpEnabled = uint8(pb.ExpEnabled)
	e.AAPointsSpent = uint32(pb.AaPointsSpent)
	e.AAExp = uint32(pb.AaExp)
	e.AAPoints = uint32(pb.AaPoints)
	e.GroupLeadershipExp = uint32(pb.GroupLeadershipExp)
	e.RaidLeadershipExp = uint32(pb.RaidLeadershipExp)
	e.GroupLeadershipPoints = uint32(pb.GroupLeadershipPoints)
	e.RaidLeadershipPoints = uint32(pb.RaidLeadershipPoints)
	e.Points = uint32(pb.Points)
	e.CurHP = uint32(pb.CurHp)
	e.Mana = uint32(pb.Mana)
	e.Endurance = uint32(pb.Endurance)
	e.Intoxication = uint32(pb.Intoxication)
	e.STR = uint32(pb.Str)
	e.STA = uint32(pb.Sta)
	e.CHA = uint32(pb.Cha)
	e.DEX = uint32(pb.Dex)
	e.INT = uint32(pb.Int)
	e.AGI = uint32(pb.Agi)
	e.WIS = uint32(pb.Wis)
	e.ZoneChangeCount = uint32(pb.ZoneChangeCount)
	e.Toxicity = uint32(pb.Toxicity)
	e.HungerLevel = uint32(pb.HungerLevel)
	e.ThirstLevel = uint32(pb.ThirstLevel)
	e.AbilityUp = uint32(pb.AbilityUp)
	e.LdonPointsGuk = uint32(pb.LdonPointsGuk)
	e.LdonPointsMir = uint32(pb.LdonPointsMir)
	e.LdonPointsMmc = uint32(pb.LdonPointsMmc)
	e.LdonPointsRuj = uint32(pb.LdonPointsRuj)
	e.LdonPointsTak = uint32(pb.LdonPointsTak)
	e.LdonPointsAvailable = uint32(pb.LdonPointsAvailable)
	e.TributeTimeRemaining = uint32(pb.TributeTimeRemaining)
	e.CareerTributePoints = uint32(pb.CareerTributePoints)
	e.TributePoints = uint32(pb.TributePoints)
	e.TributeActive = uint32(pb.TributeActive)
	e.PvPStatus = uint8(pb.PvpStatus)
	e.PvPKills = uint32(pb.PvpKills)
	e.PvPDeaths = uint32(pb.PvpDeaths)
	e.PvPCurrentPoints = uint32(pb.PvpCurrentPoints)
	e.PvPCareerPoints = uint32(pb.PvpCareerPoints)
	e.PvPBestKillStreak = uint32(pb.PvpBestKillStreak)
	e.PvPWorstDeathStreak = uint32(pb.PvpWorstDeathStreak)
	e.PvPCurrentKillStreak = uint32(pb.PvpCurrentKillStreak)
	e.PvP2 = uint32(pb.Pvp2)
	e.PvPType = uint32(pb.PvpType)
	e.ShowHelm = uint32(pb.ShowHelm)
	e.GroupAutoConsent = uint8(pb.GroupAutoConsent)
	e.RaidAutoConsent = uint8(pb.RaidAutoConsent)
	e.GuildAutoConsent = uint8(pb.GuildAutoConsent)
	e.LeadershipExpOn = uint8(pb.LeadershipExpOn)
	e.RestTimer = uint32(pb.RestTimer)
	e.AirRemaining = uint32(pb.AirRemaining)
	e.AutosplitEnabled = uint8(pb.AutosplitEnabled)
	e.Lfp = uint8(pb.Lfp)
	e.Lfg = uint8(pb.Lfg)
	e.Mailkey = pb.Mailkey
	e.Xtargets = uint32(pb.Xtargets)
	e.Firstlogon = uint32(pb.Firstlogon)
	e.EAaEffects = uint32(pb.EAaEffects)
	e.EPercentToAA = uint32(pb.EPercentToAa)
	e.EExpendedAaSpent = uint32(pb.EExpendedAaSpent)
	e.AAPointsSpentOld = uint32(pb.AaPointsSpentOld)
	e.AAPointsOld = uint32(pb.AaPointsOld)
	e.ELastInvsnapshot = uint32(pb.ELastInvsnapshot)
	e.DeletedAt = sql.NullTime{
		Valid: pb.DeletedAt != nil,
		Time:  time.Unix(pb.DeletedAt.GetSeconds(), int64(pb.DeletedAt.GetNanos())).UTC(),
	}
}

func (e *Character) Write() *charactersv1.Character {
	character := &charactersv1.Character{
		Id:                    int32(e.ID),
		Name:                  e.Name,
		LastName:              e.LastName,
		Title:                 e.Title,
		Suffix:                e.Suffix,
		ZoneId:                int32(e.ZoneID),
		ZoneInstance:          int32(e.ZoneInstance),
		Y:                     e.Y,
		X:                     e.X,
		Z:                     e.Z,
		Heading:               e.Heading,
		Gender:                int32(e.Gender),
		Race:                  int32(e.Race),
		Class:                 int32(e.Class),
		Level:                 int32(e.Level),
		Deity:                 int32(e.Deity),
		Birthday:              int32(e.Birthday),
		LastLogin:             int32(e.LastLogin),
		TimePlayed:            int32(e.TimePlayed),
		Level2:                int32(e.Level2),
		Anon:                  int32(e.Anon),
		Gm:                    int32(e.GM),
		Face:                  int32(e.Face),
		HairColor:             int32(e.HairColor),
		HairStyle:             int32(e.HairStyle),
		Beard:                 int32(e.Beard),
		BeardColor:            int32(e.BeardColor),
		EyeColor_1:            int32(e.EyeColor1),
		EyeColor_2:            int32(e.EyeColor2),
		DrakkinHeritage:       int32(e.DrakkinHeritage),
		DrakkinTattoo:         int32(e.DrakkinTattoo),
		DrakkinDetails:        int32(e.DrakkinDetails),
		AbilityTimeSeconds:    int32(e.AbilityTimeSeconds),
		AbilityNumber:         int32(e.AbilityNumber),
		AbilityTimeMinutes:    int32(e.AbilityTimeMinutes),
		AbilityTimeHours:      int32(e.AbilityTimeHours),
		Exp:                   int32(e.Exp),
		ExpEnabled:            int32(e.ExpEnabled),
		AaPointsSpent:         int32(e.AAPointsSpent),
		AaExp:                 int32(e.AAExp),
		AaPoints:              int32(e.AAPoints),
		GroupLeadershipExp:    int32(e.GroupLeadershipExp),
		RaidLeadershipExp:     int32(e.RaidLeadershipExp),
		GroupLeadershipPoints: int32(e.GroupLeadershipPoints),
		RaidLeadershipPoints:  int32(e.RaidLeadershipPoints),
		Points:                int32(e.Points),
		CurHp:                 int32(e.CurHP),
		Mana:                  int32(e.Mana),
		Endurance:             int32(e.Endurance),
		Intoxication:          int32(e.Intoxication),
		Str:                   int32(e.STR),
		Sta:                   int32(e.STA),
		Cha:                   int32(e.CHA),
		Dex:                   int32(e.DEX),
		Int:                   int32(e.INT),
		Agi:                   int32(e.AGI),
		Wis:                   int32(e.WIS),
		ZoneChangeCount:       int32(e.ZoneChangeCount),
		Toxicity:              int32(e.Toxicity),
		HungerLevel:           int32(e.HungerLevel),
		ThirstLevel:           int32(e.ThirstLevel),
		AbilityUp:             int32(e.AbilityUp),
		LdonPointsGuk:         int32(e.LdonPointsGuk),
		LdonPointsMir:         int32(e.LdonPointsMir),
		LdonPointsMmc:         int32(e.LdonPointsMmc),
		LdonPointsRuj:         int32(e.LdonPointsRuj),
		LdonPointsTak:         int32(e.LdonPointsTak),
		LdonPointsAvailable:   int32(e.LdonPointsAvailable),
		TributeTimeRemaining:  int32(e.TributeTimeRemaining),
		CareerTributePoints:   int32(e.CareerTributePoints),
		TributePoints:         int32(e.TributePoints),
		TributeActive:         int32(e.TributeActive),
		PvpStatus:             int32(e.PvPStatus),
		PvpKills:              int32(e.PvPKills),
		PvpDeaths:             int32(e.PvPDeaths),
		PvpCurrentPoints:      int32(e.PvPCurrentPoints),
		PvpCareerPoints:       int32(e.PvPCareerPoints),
		PvpBestKillStreak:     int32(e.PvPBestKillStreak),
		PvpWorstDeathStreak:   int32(e.PvPWorstDeathStreak),
		PvpCurrentKillStreak:  int32(e.PvPCurrentKillStreak),
		Pvp2:                  int32(e.PvP2),
		PvpType:               int32(e.PvPType),
		ShowHelm:              int32(e.ShowHelm),
		GroupAutoConsent:      int32(e.GroupAutoConsent),
		RaidAutoConsent:       int32(e.RaidAutoConsent),
		GuildAutoConsent:      int32(e.GuildAutoConsent),
		LeadershipExpOn:       int32(e.LeadershipExpOn),
		RestTimer:             int32(e.RestTimer),
		AirRemaining:          int32(e.AirRemaining),
		AutosplitEnabled:      int32(e.AutosplitEnabled),
		Lfp:                   int32(e.Lfp),
		Lfg:                   int32(e.Lfg),
		Mailkey:               e.Mailkey,
		Xtargets:              int32(e.Xtargets),
		Firstlogon:            int32(e.Firstlogon),
		EAaEffects:            int32(e.EAaEffects),
		EPercentToAa:          int32(e.EPercentToAA),
		EExpendedAaSpent:      int32(e.EExpendedAaSpent),
		AaPointsSpentOld:      int32(e.AAPointsSpentOld),
		AaPointsOld:           int32(e.AAPointsOld),
		ELastInvsnapshot:      int32(e.ELastInvsnapshot),
	}

	if e.DeletedAt.Valid {
		character.DeletedAt = &timestamppb.Timestamp{Seconds: int64(e.DeletedAt.Time.Unix()), Nanos: int32(e.DeletedAt.Time.Nanosecond())}
	}
	return character
}
