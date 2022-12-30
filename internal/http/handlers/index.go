package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/AliRasoulinejad/hetab/internal/app"
)

func Index(ctx echo.Context) error {
	log.Info("index")
	return ctx.String(http.StatusOK, app.Banner())
}

func Health(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNoContent)
}
