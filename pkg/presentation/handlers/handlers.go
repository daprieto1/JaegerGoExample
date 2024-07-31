package handlers

import (
	"errors"
	"net/http"

	"github.com/daprieto1/tracing/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type HandlersImplementation struct {
	usecase usecase.UseCaseImplementation
}

func NewHandlersImplementation(
	usecase usecase.UseCaseImplementation,
) HandlersImplementation {
	return HandlersImplementation{
		usecase: usecase,
	}
}

func jsonErrorResponse(c *gin.Context, statusCode int, err error) {
	c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
}

func (h HandlersImplementation) CreateProduct(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.Start(ctx, "CreateProduct Service")
	defer span.End()
	var input usecase.Product

	if err := c.BindJSON(&input); err != nil {
		jsonErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if input.Name == "" {
		jsonErrorResponse(c, http.StatusBadRequest, errors.New("name is required"))
		return
	}

	if input.Price == 0.0 {
		jsonErrorResponse(c, http.StatusBadRequest, errors.New("price is required"))
		return
	}

	response, err := h.usecase.CreateProduct(ctx, input)
	if err != nil {
		jsonErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
