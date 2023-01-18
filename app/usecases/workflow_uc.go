package usecases

import (
	"github.com/yafiesetyo/poc-workflow/app/model"
	"github.com/yafiesetyo/poc-workflow/app/repositories"
)

type WorkflowUsecaseImpl struct {
	Repo repositories.WorkflowRepoImpl
}

func NewWorkflowUsecase(repo repositories.WorkflowRepoImpl) WorkflowUsecaseImpl {
	return WorkflowUsecaseImpl{
		Repo: repo,
	}
}

func (uc *WorkflowUsecaseImpl) Create(req model.Workflow) error {
	return uc.Repo.Create(req)
}
