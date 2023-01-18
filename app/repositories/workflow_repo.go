package repositories

import (
	"github.com/yafiesetyo/poc-workflow/app/model"
	"gorm.io/gorm"
)

type WorkflowRepoImpl struct {
	DB *gorm.DB
}

func NewWorkflowRepo(db *gorm.DB) WorkflowRepoImpl {
	return WorkflowRepoImpl{
		DB: db,
	}
}

func (r *WorkflowRepoImpl) Create(req model.Workflow) error {
	return r.DB.Table(`workflow_rule`).Create(&req).Error
}

func (r *WorkflowRepoImpl) FindByName(name string) (res model.Workflow, err error) {
	return res, r.DB.Table(`workflow_rule`).
		Where(`"deletedAt" isnull and "name" = ?`, name).
		Scan(&res).Error
}
