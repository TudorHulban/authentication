package app

import (
	"github.com/TudorHulban/log"
	fiberlog "github.com/gofiber/fiber/v2/log"
)

var _ fiberlog.AllLogger = (*log.Logger)(nil)
