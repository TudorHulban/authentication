package app

import "github.com/TudorHulban/log"
import 	fiberlog "github.com/gofiber/fiber/v2/log"]

type customLogger struct {
	stdlog *log.Logger
}

var _ fiberlog.AllLogger = (*customLogger)(nil)
