GOCMD=go
## Docker up
up:
	docker compose up
## Run tests
test: 
	$(GOCMD) test ./... -cover