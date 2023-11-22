package raw

import (
	"database/sql"
	"time"

	accountsv1 "github.com/xackery/nf/proto/accounts/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Account represents a user account in the system.
type Account struct {
	ID                   int          `db:"id"`
	Name                 string       `db:"name"`
	CharName             string       `db:"charname"`
	SharedPlat           uint         `db:"sharedplat"`
	Password             string       `db:"password"`
	Status               int          `db:"status"`
	LoginServerID        string       `db:"ls_id"`
	LoginServerAccountID uint         `db:"lsaccount_id"`
	GMSpeed              uint8        `db:"gmspeed"`
	Invulnerable         bool         `db:"invulnerable"`
	FlyMode              bool         `db:"flymode"`
	IgnoreTells          bool         `db:"ignore_tells"`
	Revoked              uint8        `db:"revoked"`
	Karma                uint         `db:"karma"`
	MiniLoginIP          string       `db:"minilogin_ip"`
	HideMe               bool         `db:"hideme"`
	RulesFlag            bool         `db:"rulesflag"`
	SuspendedUntil       sql.NullTime `db:"suspendeduntil"`
	TimeCreation         uint         `db:"time_creation"`
	BanReason            string       `db:"ban_reason"`
	SuspendReason        string       `db:"suspend_reason"`
	CRCEqgame            string       `db:"crc_eqgame"`
	CRCSkillcaps         string       `db:"crc_skillcaps"`
	CRCBaseData          string       `db:"crc_basedata"`
}

func (e *Account) Read(pb *accountsv1.Account) {
	if pb == nil {
		return
	}
	e.ID = int(pb.Id)
	e.Name = pb.Name
	e.CharName = pb.Charname
	e.SharedPlat = uint(pb.Sharedplat)
	e.Password = pb.Password
	e.Status = int(pb.Status)
	e.LoginServerID = pb.LsId
	e.LoginServerAccountID = uint(pb.LsaccountId)
	e.GMSpeed = uint8(pb.Gmspeed)
	e.Invulnerable = pb.Invulnerable
	e.FlyMode = pb.Flymode
	e.IgnoreTells = pb.IgnoreTells
	e.Revoked = uint8(0)
	if pb.Revoked {
		e.Revoked = uint8(1)
	}
	e.Karma = uint(pb.Karma)
	e.MiniLoginIP = pb.MiniloginIp
	e.HideMe = pb.Hideme
	e.RulesFlag = pb.Rulesflag
	e.SuspendedUntil = sql.NullTime{
		Valid: pb.Suspendeduntil != nil,
		Time:  time.Unix(pb.Suspendeduntil.GetSeconds(), int64(pb.Suspendeduntil.GetNanos())).UTC(),
	}
	e.TimeCreation = uint(pb.TimeCreation.GetSeconds())
	e.BanReason = pb.BanReason
	e.SuspendReason = pb.SuspendReason
}

func (e *Account) Write() *accountsv1.Account {
	account := &accountsv1.Account{
		Id:            int32(e.ID),
		Name:          e.Name,
		Charname:      e.CharName,
		Sharedplat:    int32(e.SharedPlat),
		Password:      e.Password,
		Status:        int32(e.Status),
		LsId:          e.LoginServerID,
		LsaccountId:   int32(e.LoginServerAccountID),
		Gmspeed:       int32(e.GMSpeed),
		Invulnerable:  e.Invulnerable,
		Flymode:       e.FlyMode,
		IgnoreTells:   e.IgnoreTells,
		Revoked:       e.Revoked == 1,
		Karma:         int32(e.Karma),
		MiniloginIp:   e.MiniLoginIP,
		Hideme:        e.HideMe,
		Rulesflag:     e.RulesFlag,
		BanReason:     e.BanReason,
		SuspendReason: e.SuspendReason,
	}
	if e.SuspendedUntil.Valid {
		account.Suspendeduntil = &timestamppb.Timestamp{Seconds: int64(e.SuspendedUntil.Time.Unix()), Nanos: int32(e.SuspendedUntil.Time.Nanosecond())}
	}
	if e.TimeCreation != 0 {
		account.TimeCreation = &timestamppb.Timestamp{Seconds: int64(e.TimeCreation)}
	}
	return account
}
