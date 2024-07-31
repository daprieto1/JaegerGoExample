package handlers

import (
	"errors"
	"net/http"

	"github.com/daprieto1/tracing/pkg/usecase"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("github.com/daprieto1/tracing/pkg/presentation/handlers/handlers")

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
	ctx, span := tracer.Start(ctx, "CreateProduct Handler")
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

func (h HandlersImplementation) GetProductByName(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.Start(ctx, "GetProductByName Handler")
	defer span.End()

	name := c.Query("name")
	response, err := h.usecase.GetProductByName(ctx, name)
	if err != nil {
		jsonErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h HandlersImplementation) GetProductByDescription(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.Start(ctx, "GetProductByDescription Handler")
	defer span.End()

	description := c.Query("description")
	response, err := h.usecase.GetProductByDescription(ctx, description)
	if err != nil {
		jsonErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
