package logic

import (
	"posthub/dao/mysql"
	"posthub/model"
)

func CommunityList() (data []*model.Community, err error) {
	return mysql.CommunityList()
}

func CommunityById(id int64) (comm *model.Community, err error) {
	return mysql.CommunityById(id)
}
