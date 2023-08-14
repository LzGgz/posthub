package mysql

import (
	"posthub/model"
	"strings"

	"github.com/jmoiron/sqlx"
)

func InsertPost(p *model.Post) (err error) {
	sqlstr := "insert into post(id,author_id,community_id,title,content) values(?,?,?,?,?)"
	_, err = db.Exec(sqlstr, p.ID, p.AuthorId, p.CommunityId, p.Title, p.Content)
	return
}

func PostById(id int64) (p *model.Post, err error) {
	p = new(model.Post)
	p.ID = id
	sqlstr := "select content,title,author_id,community_id,status,created_time from post where id = ?"
	err = db.Get(p, sqlstr, id)
	return
}

func PostList(offset int, limit int) (ps []*model.Post, err error) {
	sqlstr := "select id,content,title,author_id,community_id,status,created_time from post limit ? offset ?"
	ps = make([]*model.Post, 0, limit)
	err = db.Select(&ps, sqlstr, limit, offset)
	return
}

func Posts(ids []string) (ps []*model.Post, err error) {
	sqlstr := `select id,content,title,author_id,community_id,status,created_time,updated_time 
	from post
	where id in (?) 
	order by FIND_IN_SET(id,?)`
	query, args, err := sqlx.In(sqlstr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&ps, query, args...)
	return
}
