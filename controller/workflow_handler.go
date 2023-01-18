package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yafiesetyo/poc-workflow/app/model"
	"github.com/yafiesetyo/poc-workflow/app/usecases"
)

type WorkflowHandlerImpl struct {
	Prefix  string
	Usecase usecases.WorkflowUsecaseImpl
}

func NewWorkflowHandler(uc usecases.WorkflowUsecaseImpl, prefix string) WorkflowHandlerImpl {
	return WorkflowHandlerImpl{
		Usecase: uc,
		Prefix:  prefix,
	}
}

func (h *WorkflowHandlerImpl) Mount(r *gin.RouterGroup) {
	g := r.Group("/" + h.Prefix)
	g.POST("/", h.Create)
}

func (h *WorkflowHandlerImpl) Create(c *gin.Context) {
	var req model.CreateWorkflow
	now := time.Now()
	if err := c.ShouldBind(&req); err != nil {
		log.Default().Println("getting error when bind json, err:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.Usecase.Create(model.Workflow{
		Name:      req.Name,
		First:     req.First,
		Second:    req.Second,
		Third:     req.Third,
		CreatedAt: &now,
	}); err != nil {
		log.Default().Println("getting error when bind json, err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})
}
