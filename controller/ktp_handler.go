package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yafiesetyo/poc-workflow/app/model"
	"github.com/yafiesetyo/poc-workflow/app/usecases"
)

type KtpHandlerImpl struct {
	Prefix  string
	Usecase usecases.KTPUsecaseImpl
}

func NewKtpHandler(uc usecases.KTPUsecaseImpl, prefix string) KtpHandlerImpl {
	return KtpHandlerImpl{
		Usecase: uc,
		Prefix:  prefix,
	}
}

func (h *KtpHandlerImpl) Mount(r *gin.RouterGroup) {
	g := r.Group("/" + h.Prefix)
	g.POST("/", h.Register)
	// call later (either using pubsub, cloudtask, rabbitMQ)
	g.POST("/first/:id", h.First)
	g.POST("/second/:id", h.Second)
	g.POST("/third/:id", h.Third)
}

func (h *KtpHandlerImpl) Register(c *gin.Context) {
	var req model.CommonCreateReq

	if err := c.ShouldBind(&req); err != nil {
		log.Default().Println("getting error when binding json, err", err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.Usecase.Register(req); err != nil {
		log.Default().Println("getting error when register ktp, err", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})
}

func (h *KtpHandlerImpl) First(c *gin.Context) {
	var header model.ValidateKTPHeader

	if err := c.BindHeader(&header); err != nil {
		log.Default().Println("getting error when binding headers, err", err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Default().Println("getting error when convert id, err", err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.Usecase.FirstValidation(c, id, header.XWorkflowName); err != nil {
		log.Default().Println("getting error when do first validation, error", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}
}

func (h *KtpHandlerImpl) Second(c *gin.Context) {
	var header model.ValidateKTPHeader

	if err := c.BindHeader(&header); err != nil {
		log.Default().Println("getting error when binding headers, err", err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Default().Println("getting error when convert id, err", err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.Usecase.SecondValidation(c, id, header.XWorkflowName); err != nil {
		log.Default().Println("getting error when do second validation, error", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}
}

func (h *KtpHandlerImpl) Third(c *gin.Context) {
	var header model.ValidateKTPHeader

	if err := c.BindHeader(&header); err != nil {
		log.Default().Println("getting error when binding headers, err", err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Default().Println("getting error when convert id, err", err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.Usecase.ThirdValidation(c, id, header.XWorkflowName); err != nil {
		log.Default().Println("getting error when do third validation, error", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}
}
