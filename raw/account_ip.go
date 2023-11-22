package raw

import (
	"time"

	accountsv1 "github.com/xackery/nf/proto/accounts/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountIP struct {
	AccountID int32     `db:"accid"`
	IP        string    `db:"ip"`
	Count     int32     `db:"count"`
	LastUsed  time.Time `db:"lastused"`
}

func (e *AccountIP) Read(pb *accountsv1.AccountIP) {
	if pb == nil {
		return
	}
	e.AccountID = pb.AccountId
	e.IP = pb.Ip
	e.Count = pb.Count
	e.LastUsed = pb.LastUsed.AsTime()
}

func (e *AccountIP) Write() *accountsv1.AccountIP {
	pb := &accountsv1.AccountIP{
		AccountId: e.AccountID,
		Ip:        e.IP,
		Count:     e.Count,
		LastUsed:  timestamppb.New(e.LastUsed),
	}
	return pb
}
