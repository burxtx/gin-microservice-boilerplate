package models

import "time"

type App struct {
	Id           int64     `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	RepositoryId int64     `json:"repository_id"`
	CreateTime   time.Time `json:"create_time" time_format:"2006-01-02 15:04:05"`
	UpdateTime   time.Time `xorm:"updated" time_format:"2006-01-02 15:04:05"`
}
