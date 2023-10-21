package authservices

import (
	abstract "github.com/Dparty/dao/abstract"
	"github.com/Dparty/dao/auth"
)

type AccountRole string

const (
	ROOT  AccountRole = "ROOT"
	ADMIN AccountRole = "ADMIN"
	USER  AccountRole = "USER"
)

type Account struct {
	entity auth.Account
}

func (a Account) ID() uint {
	return a.entity.ID()
}

func (a Account) Own(asset abstract.Asset) bool {
	return a.ID() == asset.Owner().ID()
}

func (a Account) Owner() abstract.Owner {
	return nil
}

func (a Account) Role() AccountRole {
	return AccountRole(a.entity.Role)
}
