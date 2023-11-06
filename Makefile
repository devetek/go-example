build:
	@docker build -t prakasa1904/go-example:latest .

ci:
	@act workflow_dispatch --container-architecture linux/amd64
