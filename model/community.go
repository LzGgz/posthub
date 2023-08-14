package model

import "time"

type Community struct {
	ID           int        `json:"id" db:"id"`
	Name         string     `json:"name" db:"name"`
	Introduction string     `json:"introduction,omitempty" db:"introduction"`
	CreatedTime  *time.Time `json:"created_time,omitempty" db:"created_time"`
	UpdatedTime  *time.Time `json:"updated_time,omitempty" db:"updated_time"`
}
