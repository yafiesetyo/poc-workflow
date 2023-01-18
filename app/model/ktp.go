package model

import "time"

var (
	StageSubmitted      = "submitted"
	StageFirstRejected  = "first_reject"
	StageSecondRejected = "second_reject"
	StageThirdRejected  = "third_reject"
)

type KTP struct {
	ID        uint64     `gorm:"column:id"`
	Name      string     `gorm:"column:name"`
	Stage     string     `gorm:"column:stage"`
	First     bool       `gorm:"column:first"`
	Second    bool       `gorm:"column:second"`
	Third     bool       `gorm:"column:third"`
	IsDone    bool       `gorm:"column:isDone"`
	CreatedAt *time.Time `gorm:"column:createdAt"`
	UpdatedAt *time.Time `gorm:"column:updatedAt"`
}

type ValidateKTPHeader struct {
	XWorkflowName string `header:"x-workflow-name"`
}
