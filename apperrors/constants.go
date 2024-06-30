package apperrors

const _space = " "

var (
	OSExitForConfigurationIssues        = 1
	OSExitForDatabaseIssues             = 2
	OSExitForRepositoryIssues           = 3
	OSExitForRepositoryMigrationsIssues = 4
	OSExitForServiceIssues              = 5
	OSExitForApplicationIssues          = 6
	OSExitForFileOperationsIssues       = 7
	OSExitForGRPCIssues                 = 8
	OSExitForFixtureIssues              = 9
	OSExitForWebServerIssues            = 10
	OSExitForControllerIssues           = 11
	OSExitForLoggingIssues              = 12
)
