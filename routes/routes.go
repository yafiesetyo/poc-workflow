package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yafiesetyo/poc-workflow/controller"
)

type HandlerList struct {
	KtpHandler      controller.KtpHandlerImpl
	WorkflowHandler controller.WorkflowHandlerImpl
}

func InitRouter(ktp controller.KtpHandlerImpl, workflow controller.WorkflowHandlerImpl) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	ktp.Mount(v1)
	workflow.Mount(v1)

	return r
}
