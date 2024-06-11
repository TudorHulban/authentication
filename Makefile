.PHONY: test
test: ## Run test with check race  and coverage
	@go test -failfast -count=1 ./... -json -cover -race | tparse -smallscreen

clean-db:
	@rm -f cmd/app-run/local_*.json