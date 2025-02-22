package api

import (
	"net/http"

	"github.com/hasanbakirci/api-observability-demo/internal/models"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(e *echo.Echo, service *Service) {
	handler := &Handler{
		service: service,
	}

	e.Use(echoprometheus.NewMiddleware("demo_api"))
	e.GET("/metrics", echoprometheus.NewHandler())
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "API is running")
	})

	g := e.Group("/api")
	g.POST("/events", handler.handleEvents)
}

func (h *Handler) handleEvents(c echo.Context) error {
	var event *models.Event
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	publishEventError := h.service.PublishEvent(*event)
	if publishEventError != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": publishEventError.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Event published"})
}
