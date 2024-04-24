run:
	@echo "Initializing..."
	@node ./nodejs/src/index.js & go run ./go/cmd/main.go

.PHONY: run
