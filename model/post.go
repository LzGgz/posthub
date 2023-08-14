package model

import "time"

type Post struct {
	ID          int64      `json:"id,string" db:"id"`
	AuthorId    int64      `json:"author_id,string" db:"author_id"`
	CommunityId int64      `json:"community_id,string" binding:"required" db:"community_id"`
	Status      int16      `json:"status" db:"status"`
	Title       string     `json:"title" binding:"required" db:"title"`
	Content     string     `json:"content" binding:"required" db:"content"`
	CreatedTime *time.Time `json:"created_time" db:"created_time"`
	UpdatedTime *time.Time `json:"updated_time" db:"updated_time"`
}

type PostDetail struct {
	AuthorName string `json:"author_name"`
	Vote       int64  `json:"votes"`
	*Post
	*Community `json:"community"`
}
