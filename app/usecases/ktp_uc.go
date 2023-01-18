package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/yafiesetyo/poc-workflow/app/model"
	"github.com/yafiesetyo/poc-workflow/app/repositories"
	"github.com/yafiesetyo/poc-workflow/infra"
)

var (
	pubsubReady = true
)

type KTPUsecaseImpl struct {
	Repo         repositories.KtpRepoImpl
	WorkflowRepo repositories.WorkflowRepoImpl
	PubSub       infra.PubSubClient
}

func NewKtpUsecase(repo repositories.KtpRepoImpl, pubsub infra.PubSubClient, workflowRepo repositories.WorkflowRepoImpl) KTPUsecaseImpl {
	return KTPUsecaseImpl{
		Repo:         repo,
		PubSub:       pubsub,
		WorkflowRepo: workflowRepo,
	}
}

func (uc *KTPUsecaseImpl) Register(req model.CommonCreateReq) error {
	dbPayload := model.KTP{
		Name:  req.Name,
		Stage: model.StageSubmitted,
	}

	return uc.Repo.Register(dbPayload)
}

func (uc *KTPUsecaseImpl) FirstValidation(ctx context.Context, id uint64, workflowName string) error {
	pubsubPayload := model.CommonPublishReq{
		ID:       id,
		RuleName: workflowName,
		Type:     model.TypeSecondEndpoint,
	}

	isProceed, err := uc.extractRule(workflowName, "first")
	if err != nil {
		log.Default().Println("getting error when extract rule, error:", err)
		return err
	}
	if !isProceed {
		return uc.send(ctx, pubsubPayload)
	}

	participant, err := uc.Repo.FindByID(id)
	if err != nil {
		log.Default().Println("getting error when get participant, error:", err)
		return err
	}

	if strings.ToLower(participant.Name) == "hitler" || strings.ToLower(participant.Name) == "mussolini" {
		err = errors.New("DONT BE FACIST OK!!!")

		if err := uc.Repo.Updates(id, map[string]interface{}{
			"stage": model.StageFirstRejected,
		}); err != nil {
			log.Default().Println("getting error when update participant, error:", err)
			return err
		}

		return err
	}

	if err := uc.Repo.Updates(id, map[string]interface{}{
		"first": true,
	}); err != nil {
		log.Default().Println("getting error when update participant, error:", err)
		return err
	}

	return uc.send(ctx, pubsubPayload)
}

func (uc *KTPUsecaseImpl) SecondValidation(ctx context.Context, id uint64, workflowName string) error {
	pubsubPayload := model.CommonPublishReq{
		ID:       id,
		RuleName: workflowName,
		Type:     model.TypeThirdEndpoint,
	}

	now := time.Now()
	isProceed, err := uc.extractRule(workflowName, "second")
	if err != nil {
		log.Default().Println("getting error when extract rule, error:", err)
		return err
	}
	if !isProceed {
		return uc.send(ctx, pubsubPayload)
	}

	participant, err := uc.Repo.FindByID(id)
	if err != nil {
		log.Default().Println("getting error when get participant, error:", err)
		return err
	}

	if strings.ToLower(participant.Name) == "son" || strings.ToLower(participant.Name) == "kane" {
		err = errors.New("SP*RS ARE SH**TTT!")

		if err := uc.Repo.Updates(id, map[string]interface{}{
			"stage":     model.StageSecondRejected,
			"updatedAt": now,
		}); err != nil {
			log.Default().Println("getting error when update participant, error:", err)
			return err
		}

		return err
	}

	if err := uc.Repo.Updates(id, map[string]interface{}{
		"second":    true,
		"updatedAt": now,
	}); err != nil {
		log.Default().Println("getting error when update participant, error:", err)
		return err
	}

	return uc.send(ctx, pubsubPayload)
}

func (uc *KTPUsecaseImpl) ThirdValidation(ctx context.Context, id uint64, workflowName string) error {
	now := time.Now()
	isProceed, err := uc.extractRule(workflowName, "third")
	if err != nil {
		log.Default().Println("getting error when extract rule, error:", err)
		return err
	}

	if !isProceed {
		if err := uc.Repo.Updates(id, map[string]interface{}{
			"updatedAt": now,
			"isDone":    true,
		}); err != nil {
			log.Default().Println("getting error when update participant, error:", err)
			return err
		}
		return nil
	}

	participant, err := uc.Repo.FindByID(id)
	if err != nil {
		log.Default().Println("getting error when get participant, error:", err)
		return err
	}

	if strings.ToLower(participant.Name) == "stalin" || strings.ToLower(participant.Name) == "lenin" {
		err = errors.New("Tercium bau2 PKI, tangkap pak!")

		if err := uc.Repo.Updates(id, map[string]interface{}{
			"stage":     model.StageThirdRejected,
			"updatedAt": now,
		}); err != nil {
			log.Default().Println("getting error when update participant, error:", err)
			return err
		}

		return err
	}

	if err := uc.Repo.Updates(id, map[string]interface{}{
		"third":     true,
		"updatedAt": now,
		"isDone":    true,
	}); err != nil {
		log.Default().Println("getting error when update participant, error:", err)
		return err
	}

	return nil
}

func (uc *KTPUsecaseImpl) extractRule(name, current string) (isProceed bool, err error) {
	rule, err := uc.WorkflowRepo.FindByName(name)
	if err != nil {
		log.Default().Println("getting error when get rule, error:", err.Error())
		return false, err
	}

	switch current {
	case "first":
		if !rule.First {
			return false, nil
		}
	case "second":
		if !rule.Second {
			return false, nil
		}
	case "third":
		if !rule.Third {
			return false, nil
		}
	default:
		return false, nil
	}

	return true, nil
}

func (uc *KTPUsecaseImpl) send(ctx context.Context, req model.CommonPublishReq) error {
	if pubsubReady {
		bt, err := json.Marshal(req)
		if err != nil {
			log.Default().Println("getting error when marshal pubsub payload, error:", err)
			return err
		}

		err = uc.PubSub.Publish(ctx, bt)
		if err != nil {
			log.Default().Println("getting error when publish, error:", err)
			return err
		}
	}
	return nil
}
