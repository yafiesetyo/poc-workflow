package model

import "time"

type Workflow struct {
	ID        uint64     `gorm:"column:id"`
	Name      string     `gorm:"column:name"`
	First     bool       `gorm:"column:first"`
	Second    bool       `gorm:"column:second"`
	Third     bool       `gorm:"column:third"`
	CreatedAt *time.Time `gorm:"column:createdAt"`
	UpdatedAt *time.Time `gorm:"column:updatedAt"`
}

type CreateWorkflow struct {
	Name   string `json:"name"`
	First  bool   `json:"first"`
	Second bool   `json:"second"`
	Third  bool   `json:"third"`
}
