IMAGE="quicknotes"
TAG=IMAGE+":latest"

.PHONY: docker.build
docker.build:
		@docker build -t $(TAG) .

.PHONY: docker.run
docker.run:
	@docker run -p 8080:8080 $(IMAGE)

.PHONY: run
run:
	@go run cmd/main.go
