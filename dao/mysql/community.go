package mysql

import (
	"database/sql"
	"posthub/model"

	"go.uber.org/zap"
)

func CommunityList() (data []*model.Community, err error) {
	sqlstr := "select id,name from community"
	if err = db.Select(&data, sqlstr); err == sql.ErrNoRows {
		zap.L().Warn("no records in table community")
		err = nil

	}
	return
}

func CommunityById(id int64) (comm *model.Community, err error) {
	comm = new(model.Community)
	sqlstr := "select id,name,introduction,created_time,updated_time from community where id = ?"
	err = db.Get(comm, sqlstr, id)
	return
}
