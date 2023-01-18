package boot

import (
	"log"

	"github.com/yafiesetyo/poc-workflow/app/repositories"
	"github.com/yafiesetyo/poc-workflow/app/usecases"
	"github.com/yafiesetyo/poc-workflow/controller"
	"github.com/yafiesetyo/poc-workflow/infra"
)

type BootSetup struct {
	KTPHandler      controller.KtpHandlerImpl
	WorkflowHandler controller.WorkflowHandlerImpl
}

func Boot() BootSetup {
	// init db
	db, err := infra.InitDB()
	if err != nil {
		log.Fatalf("failed to init db, err: %v \n", db)
	}

	// init pubsub
	pubsub := infra.NewPubSubClient()

	// init repo
	ktpRepo := repositories.NewKtpRepo(db)
	workflowRepo := repositories.NewWorkflowRepo(db)

	// init usecase
	ktpUc := usecases.NewKtpUsecase(ktpRepo, pubsub, workflowRepo)
	workflowUc := usecases.NewWorkflowUsecase(workflowRepo)

	// init handler
	ktpHandler := controller.NewKtpHandler(ktpUc, "ktp")
	workflowHandler := controller.NewWorkflowHandler(workflowUc, "workflow")

	return BootSetup{
		KTPHandler:      ktpHandler,
		WorkflowHandler: workflowHandler,
	}
}
