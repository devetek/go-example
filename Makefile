build-docker:
	@docker build -t prakasa1904/go-example:latest .

ci-test:
	@act workflow_dispatch --container-architecture linux/amd64

build:
	@go build -o goexample main.go

build-dpanel:
	@go1.21.4 build -o go-example main.go