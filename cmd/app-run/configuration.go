package main

import "github.com/TudorHulban/authentication/app"

var configuration = app.ParamsNewApp{
	Port:                   "9000",
	TemplateFolder:         "../../views/",
	TemplateFilesExtension: ".jet",

	// AuthenticationDisabled: true,
}
