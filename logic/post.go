package logic

import (
	"database/sql"
	"errors"
	"posthub/dao/mysql"
	"posthub/dao/redis"
	"posthub/model"
	"posthub/pkg/snowflake"
	"posthub/util/page"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var ErrInvalidPostId = errors.New("无效的帖子id")

func CreatePost(p *model.Post) (err error) {
	p.ID = snowflake.GenID()
	if err = mysql.InsertPost(p); err != nil {
		return
	}
	err = redis.CreatePost(p.ID)
	return
}

func PostDetailById(id int64) (pd *model.PostDetail, err error) {
	post, err := mysql.PostById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrInvalidPostId
		}
		zap.L().Error("mysql.PostById failed", zap.Int64("id", id), zap.Error(err))
		return
	}
	username, err := mysql.UsernameById(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.UsernameById", zap.Int64("id", post.AuthorId), zap.Error(err))
		return
	}
	comm, err := mysql.CommunityById(post.CommunityId)
	if err != nil {
		zap.L().Error("mysql.CommunityById failed", zap.Int64("id", post.CommunityId), zap.Error(err))
		return
	}
	pd = &model.PostDetail{
		AuthorName: username,
		Post:       post,
		Community:  comm,
	}
	return
}

func PostList(pageNum int) (pds []*model.PostDetail, err error) {
	ps, err := mysql.PostList(page.Offset(pageNum), viper.GetInt("app.page_size"))
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		zap.L().Error("mysql.PostList failed", zap.Error(err))
		return
	}
	pds = make([]*model.PostDetail, 0, len(ps))
	for _, post := range ps {
		var username string
		username, err = mysql.UsernameById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.UsernameById", zap.Int64("id", post.AuthorId), zap.Error(err))
			return
		}
		var comm *model.Community
		comm, err = mysql.CommunityById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.CommunityById failed", zap.Int64("id", post.CommunityId), zap.Error(err))
			return
		}
		pd := &model.PostDetail{
			AuthorName: username,
			Post:       post,
			Community:  comm,
		}
		pds = append(pds, pd)
	}
	return
}

func Posts(p *model.ParamPostList) (pds []*model.PostDetail, err error) {
	ids, err := redis.PostIDInOrder(p)
	if err != nil {
		zap.L().Error("redis.PostIDInOrder failed", zap.Error(err))
		return
	}
	if len(ids) == 0 {
		return
	}
	votes, err := redis.PostVotes(ids)
	if err != nil {
		zap.L().Error("redis.PostVotes failed", zap.Error(err))
		return
	}
	ps, err := mysql.Posts(ids)
	if err != nil {
		zap.L().Error("mysql.Posts failed", zap.Error(err))
		return
	}
	pds = make([]*model.PostDetail, 0, len(ps))
	for i, post := range ps {
		var username string
		username, err = mysql.UsernameById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.UsernameById", zap.Int64("id", post.AuthorId), zap.Error(err))
			return
		}
		var comm *model.Community
		comm, err = mysql.CommunityById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.CommunityById failed", zap.Int64("id", post.CommunityId), zap.Error(err))
			return
		}
		pd := &model.PostDetail{
			AuthorName: username,
			Vote:       votes[i],
			Post:       post,
			Community:  comm,
		}
		pds = append(pds, pd)
	}
	return
}
