package entity

import "time"

type Job struct {
	ID        string    `gorm:"primaryKey; char(16)" json:"id" form:"id"`
	JobName   string    `gorm:"varchar(255)" json:"job_name" form:"job_name"`
	Status    int       `gorm:"tinyint(1)" json:"status" form:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Jobs []*Job
