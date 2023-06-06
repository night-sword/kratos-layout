package data

import (
	"github.com/night-sword/kratos-layout/internal/dao"
)

var _ dao.Querier = &Database{}

type Database struct {
	data  *Data
	query *dao.Queries
}
